package square

import (
	"math"
)

type intCustomType int

const (
	SidesTriangle = 3
	SidesSquare   = 4
	SidesCircle   = 0
)

func CalcSquare(sideLen float64, sidesNum intCustomType) float64 {
	var res float64
	switch sidesNum {
	case SidesCircle:
		res = math.Pi * math.Pow(sideLen, 2)
		return res
	case SidesTriangle:
		res = math.Sqrt(3) * math.Pow(sideLen, 2) / 4
		return res
	case SidesSquare:
		res = math.Pow(sideLen, 2)
		return res
	default:
		return 0
	}
}
