package main

import (
	"fmt"
	"github.com/user/go-debug-example/pkg/calculator"
	"github.com/user/go-debug-example/pkg/shapes"
)

func main() {
	// 1. 测试局部变量、函数调用和全局常量
	sum := calculator.Add(10, 20)
	fmt.Printf("Sum: %d, PI: %f\n", sum, calculator.PI)

	// 2. 测试结构体和方法调用
	myCalc := calculator.Calculator{Model: "TI-84"}
	product := myCalc.Multiply(10, 20)
	fmt.Printf("Product: %d\n", product)

	// 3. 测试接口和动态派发
	shapeList := []shapes.Shape{
		shapes.Circle{Radius: 5.0},
		shapes.Rectangle{Width: 4.0, Height: 6.0},
	}
	for _, s := range shapeList {
		fmt.Printf("Area: %f\n", s.Area())
	}
}