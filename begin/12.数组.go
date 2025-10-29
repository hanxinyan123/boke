package main

import "fmt"

func main() {
	//1.必须声明数组长度 2.数组元素类型一致
	var name1 = "hanhan"
	//完整赋值格式 注意赋值类型和赋值个数
	//var 数组变量名 [数组长度]元素类型 = [数组长度]元素类型{元素1,元素2,...}
	//声明方式
	//var arr1 [3]int           // 声明一个长度为3的int数组，初始化为零值 [0 0 0]
	//arr2 := [3]int{1, 2, 3}   // 声明并初始化 [1 2 3]
	//arr3 := [...]int{4, 5, 6} // 编译器推导长度为3 [4 5 6]
	var nameList [3]string
	nameList = [3]string{name1, "jingjing", "lan"}
	fmt.Println(nameList[0]) //索引从0开始
	fmt.Println(nameList[2])
	nameList[2] = "xiaowang"
	for i, s := range nameList {
		fmt.Println(i, s)
	}

	//	go不支持负向索引，求倒数元素
	fmt.Println(len(nameList[1])) //第二个元素长度
	fmt.Println(len(nameList))    //数组长度
	fmt.Printf("%s", nameList[len(nameList)-1])

	var array2 = [...]int{1, 2, 3}
	fmt.Println(array2)

	a := [...]int{2, 3, 6}
	fmt.Println(a)
}
