package cmd

import (
	"io"

	"github.com/meyakovenkoj/go-gost-encrypt/encrypt"
)

type Encrypter struct {
	key []byte
	w   io.Writer
}

func NewWriter(w io.Writer, key []byte) (*Encrypter, error) {

	var e Encrypter
	e.key = key
	if w != nil {
		e.Reset(w)
	}
	return &e, nil
}

func (w *Encrypter) Reset(writer io.Writer) {
	w.w = writer
}

func (e *Encrypter) Write(p []byte) (n int, err error) {
	data := encrypt.EncryptData(p, e.key)
	n, err = e.w.Write(data)
	return
}
