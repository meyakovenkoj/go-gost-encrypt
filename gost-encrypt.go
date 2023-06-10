package main

// /*
// #cgo LDFLAGS: ./lib/kuznechik_cuda/libkuznechik.a
// #cgo CFLAGS: -I./lib/kuznechik_cuda/include/
// #include "kuznechik.h"
// */
// import "C"
import (
	"encoding/hex"
	"flag"
	"fmt"
	"os"

	"github.com/meyakovenkoj/go-gost-encrypt/encrypt"
	"github.com/meyakovenkoj/go-gost-encrypt/keygen"
	"github.com/meyakovenkoj/go-gost-encrypt/vko"
)

func Hello() string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Welcome?")
	return message
}

func main() {
	keyf := flag.String("key", "", "private key file")
	pkeyf := flag.String("pkey", "", "public key file")
	file := flag.String("file", "", "file to encrypt/decrypt")
	outfile := flag.String("outfile", "", "output file")
	gen := flag.Bool("gen", false, "generate and write to file new key")
	vko_send := flag.Bool("vko-send", false, "vko test send (req public key)")
	vko_recv := flag.Bool("vko-recv", false, "vko test recieve (req private key)")
	ephem_key := flag.String("ekey", "", "ephem public key")

	flag.Parse()
	key := [64]byte{}
	if *vko_send && pkeyf != nil {
		prvKey, _ := keygen.GeneratePrvKey()
		pubKey, _ := keygen.GetPublicKey(prvKey)
		userPubKey, err := os.ReadFile(*pkeyf)

		if err != nil {
			panic(err)
		}
		// следует использовать от 1 до 2^n/2-1, длинна до n/2
		ukmRaw, _ := hex.DecodeString("0000000000000001")
		commKey, err := vko.VKO2012256(prvKey, userPubKey, ukmRaw[:])
		if err != nil {
			panic(err)
		}
		fmt.Println("kek: ", hex.EncodeToString(commKey))
		fmt.Println("ephem_pub: ", hex.EncodeToString(pubKey))
		return
	} else if *vko_recv && keyf != nil && ephem_key != nil {
		userPrvKey, err := os.ReadFile(*keyf)
		if err != nil {
			panic(err)
		}
		pubKey, err := hex.DecodeString(*ephem_key)
		if err != nil {
			panic(err)
		}
		// следует использовать от 1 до 2^n/2-1, длинна до n/2
		ukmRaw, _ := hex.DecodeString("0000000000000001")
		commKey, err := vko.VKO2012256(userPrvKey, pubKey, ukmRaw[:])
		if err != nil {
			panic(err)
		}
		fmt.Println("kek: ", hex.EncodeToString(commKey))
		return
	} else if *gen {
		usingKey, _ := keygen.GeneratePrvKey()
		fmt.Println("key: ", hex.EncodeToString(usingKey))
		os.WriteFile("genkey.prv", usingKey, 0644)
		pubKey, _ := keygen.GetPublicKey(usingKey)
		fmt.Println("public key: ", hex.EncodeToString(pubKey))
		os.WriteFile("genkey.pub", pubKey, 0644)
		copy(key[:], usingKey)
	} else if keyf != nil {
		usingKey, err := os.ReadFile(*keyf)
		if err != nil {
			panic(err)
		}
		copy(key[:], usingKey)
	}

	in_data := []byte{}
	if file != nil {
		data, err := os.ReadFile(*file)
		if err != nil {
			panic(err)
		}
		in_data = data
	} else {
		panic("no file")
	}

	outf := ""
	if outfile != nil {
		outf = *outfile
	} else {
		outf = "data"
	}

	fmt.Println("out file: ", outf)
	bytes_data := encrypt.EncryptData(in_data, key[:])
	os.WriteFile(outf, bytes_data, 0644)
}
