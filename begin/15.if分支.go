package main

import "fmt"

func main() {
	//1.判断
	//if
	//if else
	//if else if else if else(无条件)
	var score = 89
	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else {
		fmt.Println("D")
	}
	//2.枚举
	age := 20
	if age >= 90 {
		fmt.Println("A")
	}
	if age >= 80 && age < 90 {
		fmt.Println("B")
	}
	if age >= 60 && age < 80 {
		fmt.Println("C")
	}
	if age < 60 {
		fmt.Println("D")
	}
	//3.嵌套
	high := 161
	if high > 170 {
		if high < 180 {
			fmt.Println("高")
		} else {
			fmt.Println("很高")
		}
	} else {
		if high >= 160 {
			fmt.Println("正常")
		} else {
			fmt.Println("矮")
		}
	}
	//4.卫语句
	age1 := 71
	if age1 >= 90 {
		fmt.Println("A")
		return
	}
	if age1 >= 80 {
		fmt.Println("B")
		return
	}
	if age1 >= 60 {
		fmt.Println("C")
		return
	}
	if age1 < 60 {
		fmt.Println("D")
		return
	}
}
