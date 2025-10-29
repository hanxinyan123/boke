package main

import "fmt"

func main() {
	var day int
	fmt.Scanln(&day)
	switch day {
	//case中1.条件 2.关键字，完成一个不会继续进行，fallthrough指令继续向下一步
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	case 6:
		fmt.Println(6)
		//使用 fallthrough 可以强制继续执行下一个 case
		fallthrough
	case 7:
		fmt.Println(7)
	default:
		fmt.Println("错误")
	}

}
