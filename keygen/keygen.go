package keygen

import (
	"crypto/rand"
	"fmt"

	"github.com/pedroalbanese/gogost/gost3410"
)

func GeneratePrvKey() ([]byte, error) {
	c := gost3410.CurveIdtc26gost341012512paramSetA()
	prvKey, err := gost3410.GenPrivateKey(c, rand.Reader)
	return prvKey.Raw(), err
}

func GetPublicKey(prvRaw []byte) ([]byte, error) {
	c := gost3410.CurveIdtc26gost341012512paramSetA()
	prv, err := gost3410.NewPrivateKey(c, prvRaw[:])
	if err != nil {
		fmt.Println("Failed to get key")
	}
	pub, err := prv.PublicKey()
	return pub.Raw(), err
}
