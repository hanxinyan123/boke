package main

import "fmt"

func main() {
	//	缓存无序，必初始化
	//var pageMap map[string]string
	var pageMap = map[string]string{}
	//var pageMap = make(map[string]string)
	pageMap["/page1"] = "{\"data\":\"page1\"}"
	fmt.Println(pageMap)
	fmt.Println(pageMap["/page1"])

	pageMap["/page2"] = "{\"data\":\"page2\"}"
	//pageMap["/page3"] = "{data:page3}"
	//fmt.Println(pageMap)
	//fmt.Println(pageMap["/page3"])
	//	循环
	for key, value := range pageMap {
		fmt.Println(key, value)
	}

	value, ok := pageMap["/page3"]
	fmt.Println(value, ok)

	//	删除
	delete(pageMap, "/page2")
	fmt.Println(pageMap)

	// 声明并赋值
	var m1 = map[string]int{
		"age1": 21,
	}
	fmt.Println(m1)
	//age1 := m1["age1"] // 取一个不存在的
	//fmt.Println(age1)
	age2, o := m1["age1"] //age2参数,o布尔型
	fmt.Println(age2, o)

	var m map[string]string
	m = make(map[string]string, 2)
	m["name"] = "hanhan"
	fmt.Println(m)

	fmt.Println("-------------简记---------------------")
	fmt.Println("map[键(key)]=值,有则修改，无则添加")
	fmt.Println("delete(map名,键(key))")
	map1 := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4, "e": 5}
	map1["a"] = 100
	fmt.Println(map1)
	map1["cv"] = 188
	fmt.Println(map1)
	delete(map1, "a")
	fmt.Println(map1)
	fmt.Println("map迭代节点随机")
	for key, value := range map1 {
		fmt.Println(key, value)
	}
	fmt.Println("拓展：字符串迭代")
	for key, value := range "hello" {
		//fmt.Println(key, value)
		fmt.Printf("key:%v,value:%c\n", key, value)
	}
	b := map1["b"]
	fmt.Println("b:", b)
}
