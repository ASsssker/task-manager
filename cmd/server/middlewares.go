package server

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
	"time"
)

type (
    responseInfo struct {
        size int
        status int
    }

    loggingResponseWriter struct {
        responseInfo
        http.ResponseWriter
    }
)

func (lrw *loggingResponseWriter) WriteHeader(status int) {
	lrw.status = status
	lrw.ResponseWriter.WriteHeader(status)
}

func (lrw *loggingResponseWriter) Write(b []byte) (int, error) {
	size, err := lrw.ResponseWriter.Write(b)
	lrw.size += size
	return size, err
}

type GzipWriter struct {
	http.ResponseWriter
	Writer io.Writer // Writer будет использоваться для записи в ответ данных в сжатом виде
}

func (gw GzipWriter) Write(b []byte) (int, error) {
	return gw.Writer.Write(b)
}

func (app *Applicaton) RequestsInfo(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        lrw := &loggingResponseWriter{
            ResponseWriter: w,
            responseInfo: responseInfo{},
        }
        start := time.Now()
        next.ServeHTTP(lrw, r)
        duration :=time.Since(start)
        app.InfoLog.Printf("URI: %s\tMethod: %s\tTime: %v\tSize: %d\tStatus: %d", r.RequestURI, r.Method, duration, lrw.size, lrw.status)
    })
}

func (app *Applicaton) CompressResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
		if err != nil {
			io.WriteString(w, err.Error())
			return
		}
		defer gz.Close()

		w.Header().Set("Content-Encoding", "gzip")
		next.ServeHTTP(GzipWriter{Writer: gz, ResponseWriter: w}, r)
	})
}