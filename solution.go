package square

import "math"

type Sides int

const (
	SidesCircle   Sides = 0
	SidesTriangle Sides = 3
	SidesSquare   Sides = 4
)

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	switch sidesNum {
	case SidesCircle:
		return math.Pi * math.Pow(sideLen, 2)

	case SidesTriangle:
		h := math.Sqrt(math.Pow(sideLen, 2) - math.Pow(sideLen/2, 2))
		return h * sideLen / 2

	case SidesSquare:
		return math.Pow(sideLen, 2)

	default:
		return 0
	}
}
