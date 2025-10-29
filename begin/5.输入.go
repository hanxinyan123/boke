package main

import "fmt"

func main() {
	var age int
	var sex string
	//scan，enter
	//scanln,空格
	//scantf,空格或自定义分隔符

	//a分隔符位置，或，、...
	num, err := fmt.Scanf("%d a %s", &age, &sex)
	fmt.Println(num, err, age, sex)
}
