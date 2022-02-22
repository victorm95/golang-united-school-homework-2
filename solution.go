package square

import "math"

type Sides int

const (
	SideCircle   Sides = 0
	SideTriangle Sides = 3
	SideSquare   Sides = 4
)

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
	switch sidesNum {
	case SideCircle:
		return math.Pi * math.Pow(sideLen, 2)

	case SideTriangle:
		h := math.Sqrt(math.Pow(sideLen, 2) - math.Pow(sideLen/2, 2))
		return h * sideLen / 2

	case SideSquare:
		return math.Pow(sideLen, 2)

	default:
		return 0
	}
}
