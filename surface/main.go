// Surface computes an SVG rendering of a 3D surface function
// This program demonstrates the mapping between 3 different coordinate systems
// #1 - 2D grid of 100x100 cells identified by interger coordinates (i,j)
// #2 - Mesh of 3D floating point coordinates (x,y,z) (cont)
// where x and y are linear functions of i and j, starting at (0,0)
// Plotting from the back to the front so that
// background polygons may be obscured by foreground ones
// #3 - 2D image canvas at (0,0) points in this plane are denoted by (sx,sy)
// an isometric projection to map each 3D point (x,y,z) onto the 2D canvas

package main

import (
	"fmt"
	"math"
)

const (
	width, height = 1200, 1200          // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange .. + xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30deg)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30deg), cos(30deg)

func main() {
	fmt.Printf("<svg xmlns'http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			fmt.Printf("<polygon points='%g,%g %g,%g %g, %g %g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	// find point (x,y) at corner of cell (i, j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// compute surface height of z
	z := f(x, y)

	// project (x,y,z) isometrically onto 2-D SVG Canvas
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0, 0)
	return math.Sin(r) / r
}
