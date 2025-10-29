package main

import "fmt"

func main() {
	//int8 一位占位符，范围-128-127
	//uint8 无占位符
	//int 取决于系统（x64）2的64-1
	var v1 uint8 = 128
	fmt.Println(v1)
	fmt.Printf("%b", v1)
	var v2 = 9223372036854775807
	fmt.Println(v2)

	var a = 94786
	fmt.Printf("%c\n", a)
	var a1 byte = 'a'
	var a2 byte = 'b'
	fmt.Println(a1 + a2)

}
