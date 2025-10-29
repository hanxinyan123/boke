package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\"韩韩")
	fmt.Println("\\韩韩")
	fmt.Println("\t韩韩")
	fmt.Println("韩韩\n")

	var f = `1
2
3
3`
	fmt.Println(f)

	//	计算次数
	fmt.Println(strings.Count("11111111111", "1"))
	//去除前后方空格
	fmt.Println(strings.TrimSpace("  abcd   "))
	//后缀匹配
	fmt.Println(strings.HasSuffix("aa.jpg", ".jpg"))
	//前缀匹配
	fmt.Println(strings.HasPrefix("outstandinguser", "outstanding"))
	//	字符串包含
	fmt.Println(strings.Contains("aaabbbccc", "ab"))
	//替换
	fmt.Println(strings.Replace("aaabbb4nnn", "a", "c", 2)) //次数
	fmt.Println(strings.ReplaceAll("aaabbb4nnn", "a", "c"))
	//大小写转换
	fmt.Println(strings.ToUpper("abcD"))
	fmt.Println(strings.ToLower("ABCd"))
	fmt.Println(strings.ToTitle("I wish you")) //全大写

}
