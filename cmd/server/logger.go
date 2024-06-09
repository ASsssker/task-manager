package server

import (
	"io"
	"log"

	c "github.com/fatih/color"
)

func getLogger(out io.Writer, preifx string, color c.Attribute, flag int) *log.Logger {
	colorPrifixe := c.New(color).Sprintf("%s\t", preifx)
	l := log.New(out, colorPrifixe, flag)

	return l
}