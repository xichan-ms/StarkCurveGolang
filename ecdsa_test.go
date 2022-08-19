package starkcurve

import (
	"testing"
	"strconv"
)

// 5000 signature verified by official stark curve sdk, see https://github.com/xichan-ms/starkware-crypto-utils/blob/dev/src/js/signature.js#L539

func TestECDSASignAndVerify(t *testing.T) {
	for i := 1; i < 5000; i++ {
		hash, _ := GetRandomNum(Stark().N)
		privateKey, _ := GetRandomNum(Stark().N)
		xPublicKey, yPublicKey := Stark().ScalarBaseMult(privateKey.Bytes())

		r, s, err := Sign(privateKey, hash)
		if err != nil {
			t.Fail()
		}

		rlt := Verify(xPublicKey, yPublicKey, r, s, hash)
		if !rlt {
			t.Fail()
		}

		t.Log(hash.String())
		t.Log(privateKey.String())
		t.Log(xPublicKey.String())
		t.Log(yPublicKey.String())
		t.Log(r.String())
		t.Log(s.String())

		t.Log(" ----------------------------------- " + strconv.Itoa(i) + " ----------------------------------- ")
	}
}
