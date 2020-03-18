package main

import "fmt"

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

	p.Render()

	ps := make([]Point, 5)
	for _, p := range ps {
		fmt.Println(p.X, p.Y)
	}
}

// メソッド (任意の型に特化した関数)
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
