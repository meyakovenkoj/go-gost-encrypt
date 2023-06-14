package vko

import "github.com/pedroalbanese/gogost/gost3410"

func VKO2012256(prvRaw1 []byte, pubRaw2 []byte, ukmRaw []byte) ([]byte, error) {
	c := gost3410.CurveIdtc26gost341012256paramSetA()
	prv1, err := gost3410.NewPrivateKey(c, prvRaw1[:])
	if err != nil {
		panic(err)
	}

	pub2, err := gost3410.NewPublicKey(c, pubRaw2[:])
	if err != nil {
		panic(err)
	}
	ukm := gost3410.NewUKM(ukmRaw[:])
	kek1, err := prv1.KEK2012256(pub2, ukm)
	if err != nil {
		panic(err)
	}
	return kek1, nil
}
