package vko_test

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/meyakovenkoj/go-gost-encrypt/vko"
)

func TestVKO2012256(t *testing.T) {
	ukmRaw, _ := hex.DecodeString("1d80603c8544c727")
	prvRawA, _ := hex.DecodeString("c990ecd972fce84ec4db022778f50fcac726f46708384b8d458304962d7147f8c2db41cef22c90b102f2968404f9b9be6d47c79692d81826b32b8daca43cb667")
	pubRawA, _ := hex.DecodeString("aab0eda4abff21208d18799fb9a8556654ba783070eba10cb9abb253ec56dcf5d3ccba6192e464e6e5bcb6dea137792f2431f6c897eb1b3c0cc14327b1adc0a7914613a3074e363aedb204d38d3563971bd8758e878c9db11403721b48002d38461f92472d40ea92f9958c0ffa4c93756401b97f89fdbe0b5e46e4a4631cdb5a")
	prvRawB, _ := hex.DecodeString("48c859f7b6f11585887cc05ec6ef1390cfea739b1a18c0d4662293ef63b79e3b8014070b44918590b4b996acfea4edfbbbcccc8c06edd8bf5bda92a51392d0db")
	pubRawB, _ := hex.DecodeString("192fe183b9713a077253c72c8735de2ea42a3dbc66ea317838b65fa32523cd5efca974eda7c863f4954d1147f1f2b25c395fce1c129175e876d132e94ed5a65104883b414c9b592ec4dc84826f07d0b6d9006dda176ce48c391e3f97d102e03bb598bf132a228a45f7201aba08fc524a2d77e43a362ab022ad4028f75bde3b79")
	kek, _ := hex.DecodeString("c9a9a77320e2cc559ed72dce6f47e2192ccea95fa648670582c054c0ef36c221")
	kekA, _ := vko.VKO2012256(prvRawA, pubRawB, ukmRaw)
	fmt.Println(kekA)
	kekB, _ := vko.VKO2012256(prvRawB, pubRawA, ukmRaw)
	if bytes.Compare(kekA, kekB) != 0 {
		t.FailNow()
	}
	if bytes.Compare(kekA, kek) != 0 {
		t.FailNow()
	}
}
