package main

import "fmt"

func main() {
	var name = "ddddd "
	fmt.Println(name)

	fmt.Printf("\"hanhan\"")
	fmt.Printf("%q %T", name, name)
	fmt.Printf("\n")
	//先定义格式
	var num = "化"
	fmt.Printf("%s\n", num)
	var num1 = '化'
	fmt.Printf("%d\n", num1)
	//格式化赋值
	var age = 12
	s := fmt.Sprintf("我的年龄是%d\n", age)
	fmt.Println(s)

	a1 := fmt.Sprintf("12")
	a2 := fmt.Sprintf("13")
	fmt.Printf("%s+%s=%s", a1, a2, a1+a2)

}
