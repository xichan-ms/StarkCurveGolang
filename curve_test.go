package starkcurve

import (
	"math/big"
	"testing"
)

func TestIsOnCurve(t *testing.T) {
	gOnCurve := Stark().IsOnCurve(Stark().Gx, Stark().Gy)
	spOnCurve := Stark().IsOnCurve(Stark().ShiftPointx, Stark().ShiftPointy)
	spNegOnCurve := Stark().IsOnCurve(Stark().MinusShiftPointx, Stark().MinusShiftPointy)

	// not in curve
	xNot, _ := new(big.Int).SetString("1468732614996758835380505372879805860898778283940581072611506469031548390000", 10)
	yNot, _ := new(big.Int).SetString("1468732614996758835380505372879805860898778283940581072611506469031500000000", 10)
	notOnCurve := Stark().IsOnCurve(xNot, yNot)

	if !gOnCurve || !spOnCurve || !spNegOnCurve || notOnCurve {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	// A point 
	xA, _ := new(big.Int).SetString("1468732614996758835380505372879805860898778283940581072611506469031548393285", 10)
	yA, _ := new(big.Int).SetString("1402551897475685522592936265087340527872184619899218186422141407423956771926", 10)
	
	// B point
	xB, _ := new(big.Int).SetString("2089986280348253421170679821480865132823066470938446095505822317253594081284", 10)
	yB, _ := new(big.Int).SetString("1713931329540660377023406109199410414810705867260802078187082345529207694986", 10)

	// expectedSum point
	xExpecSum, _ := new(big.Int).SetString("2573054162739002771275146649287762003525422629677678278801887452213127777391", 10)
	yExpecSum, _ := new(big.Int).SetString("3086444303034188041185211625370405120551769541291810669307042006593736192813", 10)

	// sum point
	xSum, ySum, _ := Stark().Add(xA, yA, xB, yB)
	if xSum.Cmp(xExpecSum) != 0 || ySum.Cmp(yExpecSum) != 0 {
		t.Fail()
	}

	_, _, err := Stark().Add(xA, yA, xA, yA)
	if err == nil {
		t.Fail()
	}
}

func TestDouble(t *testing.T) {
	// A point 
	xA, _ := new(big.Int).SetString("2089986280348253421170679821480865132823066470938446095505822317253594081284", 10)
	yA, _ := new(big.Int).SetString("1713931329540660377023406109199410414810705867260802078187082345529207694986", 10)
	
	// expected2A point
	xExpec2A, _ := new(big.Int).SetString("2286192362596033381044212504641400729870153118659330029788632461585794620008", 10)
	yExpec2A, _ := new(big.Int).SetString("3216729308426181519810461177827015848321409031316838967317660165041827217995", 10)

	x2A, y2A := Stark().Double(xA, yA)

	if x2A.Cmp(xExpec2A) != 0 || y2A.Cmp(yExpec2A) != 0 {
		t.Fail()
	}
}

func TestScalarMult(t *testing.T) {
	s, _ := new(big.Int).SetString("1234567890", 10)
	sBytes := s.Bytes()

	// expectedSG point
	xExpecSG, _ := new(big.Int).SetString("1472053989583636632020927756323198855662302025766883638055997814967334469405", 10)
	yExpecSG, _ := new(big.Int).SetString("2039280606239454840519343618231184777512824729679485983897839747749664735211", 10)

	xSG, ySG := Stark().ScalarMult(Stark().Gx, Stark().Gy, sBytes)

	if xSG.Cmp(xExpecSG) != 0 || ySG.Cmp(yExpecSG) != 0 {
		t.Fail()
	}
}

func TestScalarMultInt(t *testing.T) {
	sInt, _ := new(big.Int).SetString("1234567890", 10)

	// expectedSG point
	xExpecSG, _ := new(big.Int).SetString("1472053989583636632020927756323198855662302025766883638055997814967334469405", 10)
	yExpecSG, _ := new(big.Int).SetString("2039280606239454840519343618231184777512824729679485983897839747749664735211", 10)

	xSG, ySG := Stark().ScalarMultInt(Stark().Gx, Stark().Gy, sInt)

	if xSG.Cmp(xExpecSG) != 0 || ySG.Cmp(yExpecSG) != 0 {
		t.Fail()
	}
}

func TestScalarBaseMult(t *testing.T) {
	s, _ := new(big.Int).SetString("1234567890", 10)
	sBytes := s.Bytes()

	// expectedSG point
	xExpecSG, _ := new(big.Int).SetString("1472053989583636632020927756323198855662302025766883638055997814967334469405", 10)
	yExpecSG, _ := new(big.Int).SetString("2039280606239454840519343618231184777512824729679485983897839747749664735211", 10)

	xSG, ySG := Stark().ScalarBaseMult(sBytes)

	if xSG.Cmp(xExpecSG) != 0 || ySG.Cmp(yExpecSG) != 0 {
		t.Fail()
	}
}

/*
	how to run test:

		cd ./FastMulThreshold-DSA/crypto/secp256k1
		go test -v
*/