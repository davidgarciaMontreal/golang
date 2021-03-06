package main

import (
	"fmt"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 500                 // number of grid cells
	xyrange       = 4 * 30.0            // axis ranges (-x...+x) (-y ...+y)
	xyscale       = width / 2 / xyrange // pixels per x|y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) //

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
	NextCellJ:
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			polygonPoints := []float64{ax, ay, bx, by, cx, cy, dx, dy}
			for i, v := range polygonPoints {
				if math.IsNaN(v) || math.IsInf(v, 0) {
					fmt.Fprintf(os.Stderr, "Index[%d] in polygonPoints is invalid.\n", i)
					continue NextCellJ
				}
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := width/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)  // distance from (0,0)
	return math.Sin(r) / r // y = sin(||r||)
}

func w(x, y float64) float64 {
	return 0.099 * (math.Cos(x) + math.Cos(y))
}

func v(x, y float64) float64 {
	return -y*y/(20*20) + x*x/(99*99)
}
