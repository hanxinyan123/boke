package main

import "fmt"

func main() {
	//0-255,byte实质int,单字节字符,0-255
	var c byte = 'a'
	fmt.Println(c)
	fmt.Printf("%c\n", c)

	//rune,int32,多字节字符
	var b rune = '韩'
	fmt.Println(b)
	fmt.Printf("%c\n", b)

}
