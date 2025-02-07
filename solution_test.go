package square

import "testing"

var caseTriangle = 43.30127018922193
var caseSquare = 100.0
var caseCircle = 314.1593

func TestCalcSquareTriangle(t *testing.T) {
	result := CalcSquare(10.0, SidesTriangle)

	if result != caseTriangle {
		t.Error("Wrong triangle", result)
	}
}

func TestCalcSquareSquare(t *testing.T) {
	result := CalcSquare(10.0, SidesSquare)

	if result != caseSquare {
		t.Error("Wrong squere", result)
	}
}

func TestCalcSquareCircle(t *testing.T) {
	result := CalcSquare(10.0, SidesCircle)

	if result != caseCircle {
		t.Error("Wrong circle", result)
	}
}
