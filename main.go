package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/cpmech/gosl/rnd"
)

// 11:10 start
// 11:21 end
func normal() float64 {
	return rnd.Normal(0.5, 0.12)
}

type Circle struct {
	X, Y, R float64
}

func (c Circle) Contains(x, y float64) bool {
	return math.Hypot(x-c.X, y-c.Y) <= c.R
}

type Square struct {
	X, Y, R float64
}

func (s Square) Contains(x, y float64) bool {
	return math.Abs(x-s.X) < s.R && math.Abs(y-s.Y) < s.R
}

func (s Square) Area() float64 {
	return 4 * s.R * s.R
}

func main() {
	nFlag := flag.Int("n", 1_000_000, "number of samples to take")
	rFlag := flag.Float64("radius", 0.1, "the radius of the circle to test samples against")
	flag.Parse()

	n := *nFlag
	radius := *rFlag

	box := Square{X: 0.5, Y: 0.5, R: radius}
	circle := Circle{X: 0.5, Y: 0.5, R: radius}

	in, out, wasted := 0, 0, 0
	for i := 0; i < n; i++ {
		x, y := normal(), normal()
		if !box.Contains(x, y) {
			wasted++
			continue
		}

		if circle.Contains(x, y) {
			in++
		} else {
			out++
		}
	}

	fmt.Printf("IN %v  OUT %v  WASTED %v\n", in, out, wasted)

	circleArea := box.Area() * float64(in) / float64(in+out)

	// A = pi * r^2
	// pi = A / r^2
	pi := circleArea / circle.R / circle.R

	fmt.Printf("π = %.8f\n", pi)
}
