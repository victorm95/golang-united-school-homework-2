package square

import "math"

type Sides int

const (
		Circle Sides = 0
		Triangle Sides = 3
		Square Sides = 4
)

func CalcSquare(sideLen float64, sidesNum Sides) float64 {
		switch sidesNum {
		case Circle:
				return math.Pi * sideLen

		case Triangle:
				h := math.Sqrt(math.Pow(sideLen/2, 2) - math.Pow(sideLen, 2))
				return h * sideLen / 2

		case Square:
				return math.Pow(sideLen, 2)

		default:
				return 0
		}
}
