// ecdsa signature based on Stark curve

package starkcurve

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// [1, max)
func GetRandomNum(max *big.Int) (*big.Int, error) {
	rndPrivKey, err := rand.Int(rand.Reader, max)
	if err != nil {
		return nil, err
	}
	return rndPrivKey, nil
}

func Sign(privateKey, hash *big.Int) (*big.Int, *big.Int, error) {
	if privateKey.Cmp(big.NewInt(1)) != 1 || privateKey.Cmp(Stark().N) != -1 {
		return nil, nil, errors.New("error, private key must in scope [1, orderCurve)")
	}

	if hash.BitLen() > Stark().N.BitLen() {
		return nil, nil, errors.New("error, the length of hash can not be longer than order of curve)")
	}

	counter := 0

	for true {
		counter++
		if counter > 1000 {
			break
		}

		k, _ := GetRandomNum(Stark().N)
		xKG, _ := Stark().ScalarBaseMult(k.Bytes())
		r := new(big.Int).Mod(xKG, Stark().N)

		if r.Cmp(big.NewInt(1)) != 1 || r.Cmp(Stark().Max) != -1 {
			continue
		}

		mRPriv := new(big.Int).Mul(r, privateKey)
		mRPriv = new(big.Int).Add(mRPriv, hash)
		mRPriv = new(big.Int).Mod(mRPriv, Stark().N)

		kInv := new(big.Int).ModInverse(k, Stark().N)

		s := new(big.Int).Mul(kInv, mRPriv)
		s = new(big.Int).Mod(s, Stark().N)

		if s.Cmp(big.NewInt(1)) != 1 || s.Cmp(Stark().N) != -1 {
			continue
		}

		sInv := new(big.Int).ModInverse(s, Stark().N)

		if sInv.Cmp(big.NewInt(1)) != 1 || sInv.Cmp(Stark().Max) != -1 {
			continue
		}

		return r, s, nil
	}
	return nil, nil, errors.New("error, no obtain valid signature after try 1000 times")
}

func Verify(xPublicKey, yPublicKey, r, s, hash *big.Int) bool {
	if hash.BitLen() > Stark().N.BitLen() {
		return false
	}

	if r.Cmp(big.NewInt(1)) != 1 || r.Cmp(Stark().Max) != -1 {
		return false
	}
	if s.Cmp(big.NewInt(1)) != 1 || s.Cmp(Stark().N) != -1 {
		return false
	}
	sInv := new(big.Int).ModInverse(s, Stark().N)
	if sInv.Cmp(big.NewInt(1)) != 1 || sInv.Cmp(Stark().Max) != -1 {
		return false
	}

	u1 := new(big.Int).Mul(hash, sInv)
	u1 = new(big.Int).Mod(u1, Stark().N)
	u2 := new(big.Int).Mul(r, sInv)
	u2 = new(big.Int).Mod(u2, Stark().N)

	xU1G, yU1G := Stark().ScalarBaseMult(u1.Bytes())
	xU2Pk, yU2Pk := Stark().ScalarMult(xPublicKey, yPublicKey, u2.Bytes())

	xCalKG, _ := Stark().Add(xU1G, yU1G, xU2Pk, yU2Pk)
	rCal := new(big.Int).Mod(xCalKG, Stark().N)

	if rCal.Cmp(r) == 0 {
		return true
	} else {
		return false
	}
}



