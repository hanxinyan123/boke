package main

import "fmt"

//var age = 20
//var (
//	x1 = "q"
//	x2 = "p"
//)

// 函数体外定义赋值，可不使用
func main() {
	//函数体内定义变量必使用
	var name string
	name = "韩韩"
	fmt.Println(name)

	var name1 string = "韩韩"
	fmt.Println(name1)

	var name2 = "hanhan"
	fmt.Println(name2)

	name3 := "hanhan"
	fmt.Println(name3)

	var a1, a2, a3 = "1", "2", "3"
	fmt.Println(a1, a2, a3)

	//println(x1, x2)

	//	常量,不能修改
	const vr string = "v2"
	fmt.Println(vr)
	fmt.Println("----------------")
	fmt.Printf("%#v", vr)
}
