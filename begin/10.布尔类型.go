package main

import "fmt"

func main() {
	//go不同：
	//不可强制int转bool
	//bool无法参与数值运算，也无法与其他类型在进行转换
	a := true
	b := false
	fmt.Println(a, b)
	//默认值false
	var c bool
	fmt.Println(c)

}
