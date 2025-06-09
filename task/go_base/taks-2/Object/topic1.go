package Object

import (
	"fmt"
	"math"
)

type Shape interface {
	Area()
	Perimeter() //周长
}

// 长方体
type Rectangle struct {
	width, height float64
}

// 圆形
type Circle struct {
	radius float32
}

func (r Rectangle) Area() {
	fmt.Println("长方形 进行面积计算", r.width*r.height)
}
func (r Rectangle) Perimeter() {
	fmt.Println("长方形 进行周长计算", (r.width+r.height)*2)
}

func (c Circle) Area() {
	fmt.Println("圆形 进行面积计算", math.Pi*c.radius*c.radius)
}
func (c Circle) Perimeter() {
	fmt.Println("圆形 进行周长计算", 2*math.Pi*c.radius)
}
func maindx() {
	circle := Circle{2}
	circle.Perimeter()
	circle.Area()

	pect := Rectangle{2, 3}
	pect.Perimeter()
	pect.Area()
}
