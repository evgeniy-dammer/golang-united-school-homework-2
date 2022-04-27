package square

import (
	"math"
)

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

const (
	SidesCircle   = 0
	SidesTriangle = 3
	SidesSquare   = 4
	Pi            = 3.141593
)

type myCustomInt int

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

func CalcSquare(sideLen float64, sidesNum myCustomInt) float64 {
	var result float64

	if sidesNum == 4 {
		for i := 0; i < int(sidesNum); i++ {
			result = sideLen * sideLen
		}
		return result
	} else if sidesNum == 3 {
		result = (math.Sqrt(3) * (sideLen * sideLen)) / 4
		return result
	} else if sidesNum == 0 {
		result = Pi * (sideLen * sideLen)
		return result
	}

	return 0
}
