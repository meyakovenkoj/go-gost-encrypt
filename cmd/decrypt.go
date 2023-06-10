package cmd

import (
	"io"

	"github.com/meyakovenkoj/go-gost-encrypt/encrypt"
)

type Decrypter struct {
	key []byte
	r   io.Reader
}

func NewReader(r io.Reader, key []byte) (*Decrypter, error) {

	var d Decrypter
	d.key = key
	if r != nil {
		d.Reset(r)
	}
	return &d, nil
}

func (d *Decrypter) Reset(reader io.Reader) {
	d.r = reader
}

func (e *Decrypter) Read(p []byte) (n int, err error) {
	var data []byte
	n, err = e.r.Read(data)
	p = encrypt.EncryptData(data, e.key)
	return
}
