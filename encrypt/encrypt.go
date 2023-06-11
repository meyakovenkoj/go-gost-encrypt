package encrypt

/*
#cgo LDFLAGS: -lgostcrypt
#cgo CFLAGS: -I/usr/local/include/gostcrypt/
#include "kuznechik.h"
*/
import "C"
import (
	"unsafe"
)

func EncryptData(in_data []byte, key []byte) []byte {
	BLOCK_SIZE := 16
	padding := 0
	in_size := len(in_data)
	res := in_size % BLOCK_SIZE
	if res > 0 {
		padding = BLOCK_SIZE - in_size%BLOCK_SIZE
	}
	in_size += padding
	b2 := make([]byte, in_size)
	cb2 := (*C.uint8_t)(unsafe.Pointer(&b2[0]))
	cin_data := (*C.uint8_t)(unsafe.Pointer(&in_data[0]))

	ckey := (*C.uint8_t)(unsafe.Pointer(&key[0]))
	C.ctr_encrypt(cin_data, cb2, ckey, (C.uint)(in_size/BLOCK_SIZE))
	bytes_data := C.GoBytes(unsafe.Pointer(cb2), (C.int)(in_size))
	return bytes_data
}
