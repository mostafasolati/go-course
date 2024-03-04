package main

import (
	"errors"
	"fmt"
)

type Something interface{}

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  int
	height int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() float64 {
	return float64(r.width) * float64(r.height)
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * 3.14
}

type Triangle struct{}

func CalculateArea(s Shape) {
	fmt.Println("Area:", s.Area())
}

// func CalculateCircleArea(c Circle) {
// 	// c.Area()
// }

// func CalculateRectangleArea(r Rectangle) {
// 	// r.Area()
// }

// type any = interface{}

func PrintSomething(x any) {
	fmt.Println(x)
}

func Sum(x, y any) error {
	x2, ok1 := x.(int)
	y2, ok2 := y.(int)
	if ok1 && ok2 {
		fmt.Println(x2 + y2)
		return nil
	} else {
		return errors.New("invalid int")
	}
}

func Sum2(x, y any) {
	switch x.(type) {
	case int:
		fmt.Println("x is int", x.(int)+y.(int))
	case float64:
		fmt.Println("x is float64", x.(float64)+y.(float64))
	default:
		fmt.Println("unsupported type")
	}
}

func main() {
	r := Rectangle{width: 10, height: 20}
	c := Circle{radius: 3}
	// t := Triangle{}
	CalculateArea(r)
	CalculateArea(c)
	// CalculateArea(t)

	PrintSomething(5)
	PrintSomething("Hello")
	PrintSomething(true)

	if err := Sum(5.2, 7.3); err != nil {
		fmt.Println("Error on assertion:", err)
	}

	Sum2(10, 3)
	Sum2(7.2, 11.36)

	var x, y float32 = 3.2, 7.5
	Sum2(x, y)
}
