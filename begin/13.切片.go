package main

import (
	"fmt"
	"sort"
)

func main() {
	//数组中创建切片
	var arr = [4]string{"a", "b", "c", "d"}
	//					 0    1	   2    3
	//var s = arr[start:end),前开后闭
	var s = arr[1:3]
	fmt.Println(s)
	//索引错一位，便于理解
	var a = [4]string{"a", "b", "c", "d"}
	//				 0    1	   2    3   4
	fmt.Println(a[:1])
	fmt.Println(a[2:])
	fmt.Println(a[1:3])

	//直接声明，切片的大小不能用 [] 表示
	var w []string = []string{"a", "b"}
	//增加
	w = append(w, "c", "d")
	fmt.Println(w)
	//删除(删除第2个数),不能直接删除
	ne := append(w[:1], w[2:]...)
	fmt.Println(ne)
	//修改
	ne[0] = "hanhan"
	fmt.Println(ne)

	//make函数
	var s2 = make([]int, 50)
	fmt.Println(s2)
	var s3 []string //切片不强制赋值
	fmt.Println(s3 == nil)
	var s4 = make([]string, 0)
	fmt.Println(s4 == nil)

	//排序
	s5 := []int{1, 4, 7, 2}
	sort.Ints(s5) //升序
	//降序（万能）
	sort.Slice(s5, func(i, j int) bool {
		return s5[i] > s5[j] //升序<,降序>
	})

	fmt.Println(s5)
}
