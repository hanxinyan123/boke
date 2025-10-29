package main

import (
	"fmt"
	"math"
)

func main() {
	//同类型相加，float64,float32
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxFloat32)
	a := 10
	b := -1.2
	c := float64(a) + b
	fmt.Println(c)

	//正无穷,比所有正值大
	a1 := math.Inf(1)
	//负无穷，比所有负值小
	a2 := math.Inf(-1)
	fmt.Println(a1 > 99999999999999) // 比较正无穷大和一个很大的数
	fmt.Println(a2 < -99999999999)
	fmt.Println("----------------")

	//NaN（Not a Number，非数字）表示未定义或不可表示的值。
	fmt.Println(math.NaN())
	fmt.Println(math.IsNaN(1))
	// 1. 检查一个显式的 NaN 值
	fmt.Println(math.IsNaN(math.NaN())) // 输出: true

	//fmt.Println(0.0 / 0.0)             // 直接输出: NaN
	//fmt.Println(math.IsNaN(0.0 / 0.0)) // 检测NaN: true
}
