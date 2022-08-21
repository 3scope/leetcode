package main

import (
	"math"
	"math/rand"
	"time"
)

type CircleSolution struct {
	Radius  float64
	XCenter float64
	YCenter float64
}

// “radius”是圆的半径，“(x_center, y_center)”是圆心位置。
func CircleConstructor(radius float64, x_center float64, y_center float64) CircleSolution {
	return CircleSolution{
		Radius:  radius,
		XCenter: x_center,
		YCenter: y_center,
	}
}

// 已知圆的半径“radius”，可以得到这个圆的面积，然后通过面积随机，根据面积可以反推它到圆心的距离为多少。
func (this *CircleSolution) RandPoint() []float64 {
	rand.Seed(time.Now().UnixNano())
	// 得到一个随机的面积。
	area := math.Pi * this.Radius * this.Radius * rand.Float64()
	r := math.Sqrt(area / math.Pi)
	// 360度等于2 * Pi。
	theta := 2 * math.Pi * rand.Float64()
	x := this.XCenter + r*math.Cos(theta)
	y := this.YCenter + r*math.Sin(theta)
	return []float64{x, y}
}
