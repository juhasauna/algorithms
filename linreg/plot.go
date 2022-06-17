package linreg

import (
	"bufio"
	"flag"
	"fmt"
	"image/color"
	"log"
	"math"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg/draw"
)

// https://www.youtube.com/watch?v=ZPd_fKyrX48
// How to use flag: go run -n 100000
var iterations int

func plotMain() {
	flag.IntVar(&iterations, "n", 1000, "number of iterations")
	flag.Parse()
	xys, err := readData("linreg/data.txt")
	if err != nil {
		log.Fatalf("could not read data: %v", err)
	}

	err = plotData("linreg/out.png", xys)
	if err != nil {
		log.Fatalf("could not plotData: %v", err)
	}

}

func plotData(path string, xys plotter.XYs) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create %s: %v", path, err)
	}
	p := plot.New()

	scatter, err := plotter.NewScatter(plotter.XYs(xys))
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
	}
	scatter.GlyphStyle.Shape = draw.CrossGlyph{}
	scatter.Color = color.RGBA{R: 255, A: 255}
	p.Add(scatter)
	x, c := linearRegression(xys, 0.01)
	line, err := plotter.NewLine(plotter.XYs{
		{X: 0, Y: c}, {X: 20, Y: 20*x + c},
	})
	if err != nil {
		return fmt.Errorf("could not create NewLine: %v", err)
	}
	p.Add(line)
	wt, err := p.WriterTo(512, 512, "png")
	if err != nil {
		return fmt.Errorf("could not create writer: %v", err)
	}
	n, err := wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("could not write to %s: %v", path, err)
	}
	if err := f.Close(); err != nil {
		return fmt.Errorf("could not close %s: %v", path, err)
	}
	fmt.Printf("wrote %d things\n", n)
	return nil
}
func linearRegression(xys plotter.XYs, learningRate float64) (m, c float64) {
	for i := 0; i < iterations; i++ {
		dm, dc := conputeGradient(xys, m, c)
		m += -dm * learningRate
		c += -dc * learningRate
		// fmt.Printf("grad(%.3f, %.3f) = (%.3f, %.3f)\n", m, c, dm, dc)
		fmt.Printf("cost(%.3f, %.3f) = %.3f\n", m, c, computeCost(xys, m, c))
	}

	return m, c
}
func linearRegression_old(xys plotter.XYs) (m, c float64) {
	const (
		min   = -100.0
		max   = 100.0
		delta = 0.1
	)
	minCost := math.MaxFloat64
	for im := min; im < max; im += delta {
		for ic := min; ic < max; ic += delta {
			cost := computeCost(xys, im, ic)
			if cost < minCost {
				minCost = cost
				m, c = im, ic
				dm, dc := conputeGradient(xys, m, c)
				fmt.Printf("grad(%.2f, %.2f) = (%.2f, %.2f)\n", m, c, dm, dc)
			}
		}
	}

	fmt.Printf("cost(%.2f, %.2f) = %.2f\n", m, c, computeCost(xys, m, c))
	return m, c
}

func computeCost(xys plotter.XYs, m, c float64) float64 {
	s := 0.0
	for _, xy := range xys {
		d := xy.Y - (xy.X*m + c)
		s += d * d
	}
	return s / float64(len(xys))
}

func conputeGradient(xys plotter.XYs, m, c float64) (dm, dc float64) {

	for _, xy := range xys {
		d := xy.Y - (xy.X*m + c)
		dm += -xy.X * d
		dc += -d
	}

	n := float64(len(xys))
	dm, dc = 2/n*dm, 2/n*dc
	return dm, dc
}

func readData(path string) (plotter.XYs, error) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	var xys plotter.XYs
	s := bufio.NewScanner(f)
	for s.Scan() {
		var x, y float64
		_, err := fmt.Sscanf(s.Text(), "%f,%f", &x, &y)
		if err != nil {
			log.Printf("discarding bad data point %q: %v", s.Text(), err)
		}
		xys = append(xys, struct{ X, Y float64 }{x, y})
	}
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("could not scan: %v", err)
	}
	return xys, nil
}
