package main

import (
	"fmt"
)

type Point struct {
	X, Y int
}

func main() {
	// 構造体は値型なのでpointerを利用することが多い
	p := &Point{X: 1, Y: 2}

	// 組み込み関数 New
	p2 := new(Point)
	p2.X = 1
	p2.Y = 2

	p3 := NewPoint(1, 2)

	p.Render()
	p2.Render()
	p3.Render()
}

func (p *Point) Render() {
	fmt.Printf("<%d, %d>\n", p.X, p.Y)
}

// 型のコンストラクタという慣例の書き方
func NewPoint(x, y int) *Point {
	p := new(Point)
	p.X = x
	p.Y = y
	return p
}
