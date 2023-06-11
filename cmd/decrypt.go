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
	n, err = e.r.Read(p)
	p = encrypt.EncryptData(p, e.key)
	return
}
