package main

import (
	"flag"
	"fmt"
	"math"

	"github.com/cpmech/gosl/rnd"
	"github.com/vbauerster/mpb/v7"
	"github.com/vbauerster/mpb/v7/decor"
)

// normal returns a number according to the normal distribution in range [0, 1].
// Given in the task
func normal() float64 {
	return rnd.Normal(0.5, 0.12)
}

type Circle struct {
	X, Y, R float64
}

func (c Circle) Contains(x, y float64) bool {
	return math.Hypot(x-c.X, y-c.Y) <= c.R
}

// Square defines a square with center at X and Y coordinates and a helf-length R.
type Square struct {
	X, Y, R float64
}

func (s Square) Contains(x, y float64) bool {
	return math.Abs(x-s.X) < s.R && math.Abs(y-s.Y) < s.R
}

func (s Square) Area() float64 {
	return 4 * s.R * s.R
}

type Progress struct {
	p   *mpb.Progress
	bar *mpb.Bar
}

func (p *Progress) Init(n int) {
	p.p = mpb.New(mpb.WithWidth(60))
	p.bar = p.p.New(int64(n),
		mpb.BarStyle(),
		mpb.PrependDecorators(decor.Name("monte carlo")),
		mpb.AppendDecorators(decor.Percentage()),
	)
}

func (p *Progress) Increment() {
	if p.p == nil {
		return
	}

	p.bar.Increment()
}

func (p *Progress) Wait() {
	if p.p == nil {
		return
	}

	p.p.Wait()
}

var p Progress

func main() {
	nFlag := flag.Int("n", 1_000_000, "number of samples to take")
	rFlag := flag.Float64("radius", 0.1, "the radius of the circle to test samples against")
	barFlag := flag.Bool("bar", true, "Show/hide progress bar")
	flag.Parse()

	n := *nFlag
	radius := *rFlag

	if *barFlag {
		p.Init(n)
	}

	// We'll use the Monte-Carlo method, but it works only with uniform distribution.
	// Idea is to take rather small part of the normal distribution, so that it
	// approximates to the uniform distribution.

	// This in turn means that both the bounding box and the circle should be quite
	// small.
	box := Square{X: 0.5, Y: 0.5, R: radius}
	circle := Circle{X: 0.5, Y: 0.5, R: radius}

	in, out, wasted := 0, 0, 0
	for i := 0; i < n; i++ {
		x, y := normal(), normal()
		p.Increment()

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

	p.Wait()

	fmt.Printf("IN %v  OUT %v  WASTED %v\n", in, out, wasted)

	circleArea := box.Area() * float64(in) / float64(in+out)

	// Area = pi * r^2
	// pi = Area / r^2
	pi := circleArea / circle.R / circle.R

	fmt.Printf("π = %.8f\n", pi)
}
