package main

import (
	"fmt"
	"image/color"
	"log"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 640
	screenHeight = 640
)

type game struct {
	p []point
}

type point struct {
	x, y int
}

func (p *point) New() {
	p.x = rand.Intn(screenWidth)
	p.y = rand.Intn(screenHeight)
}

func (g *game) Layout(outWidth, outHeight int) (w, h int) { return screenWidth, screenHeight }
func (g *game) Update() error {

	return nil
}
func (g *game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{230, 230, 250, 255})
	for i := 0; i < len(g.p); i++ {
		drawRandomCircle(screen, g.p[i].x, g.p[i].y)
	}
	g.ParabolicApproximation(2)
}

func drawRandomCircle(screen *ebiten.Image, x, y int) {
	ebitenutil.DrawCircle(screen, float64(x), float64(y), 3, color.RGBA{255, 0, 0, 255})
}

func (g *game) ParabolicApproximation(polpow int) {
	n := polpow + 1
	a := make([][]float64, n)
	for i := 0; i < n; i++ {
		a[i] = make([]float64, n+1)
		for j := 0; j < n; j++ {
			a[i][j] = g.sumofx(i + j)
		}
		a[i][n] = g.sumofxy(i)
	}
	fmt.Println(a)
	ans := GaussElimination(a, n)
	fmt.Println(ans)

}
func (g *game) sumofx(power int) float64 {
	var sum float64
	for i := range g.p {
		sum += math.Pow(float64(g.p[i].x), float64(power))
	}
	return sum
}
func (g *game) sumofxy(power int) float64 {
	var sum float64
	for i := range g.p {
		sum += float64(g.p[i].y) * math.Pow(float64(g.p[i].x), float64(power))
	}
	return sum
}

func GaussElimination(a [][]float64, n int) (ans []float64) {
	ans = make([]float64, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if a[i][i] == 0 { // check if needed to swap rows
				if a[j][i] != 0 { // swap rows
					a[i], a[j] = a[j], a[i]
				}
			}
			// make [j] column for lines [i+1;...] equal to zero (triangular matrix)
			m := a[j][i] / a[i][i] // number to multiply the row [i] to get zero in [j] column
			for k := i; k <= n; k++ {
				a[j][k] += -m * a[i][k] // m is negative because we need to get zero, not 2*number after summation
			}
		}
	}
	for i := n - 1; i >= 0; i-- { // go back, calculating the answer
		ans[i] = a[i][n] / a[i][i]
		for j := i - 1; j >= 0; j-- {
			a[j][n] -= a[j][i] * ans[i]
		}
	}
	return ans
} // whole 5 loops for 1 function... sad

func main() {
	var p []point
	var a point
	for i := 0; i < 10; i++ {
		a.New()
		p = append(p, a)
	}
	ebiten.SetWindowSize(screenWidth, screenHeight)
	if err := ebiten.RunGame(&game{p: p}); err != nil {
		log.Fatal(err)
	}

}
