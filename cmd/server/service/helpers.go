package service

import (
	"encoding/json"
	"io"
)

func JsonDecode(data io.Reader, model any) error {
    if err := json.NewDecoder(data).Decode(model); err != nil {
        return err
    }

    return nil
}