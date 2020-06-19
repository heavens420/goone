package main

import "fmt"

func main() {
	//切片 变长数组
	var s1 = []int{1, 3, 4, 5}
	fmt.Println(s1)

	//mkae 创建切片 容量为6 长度为3 超出长度报错
	var s2 = make([]int, 3, 6)
	fmt.Println(cap(s2))
	fmt.Println(len(s2))

	//append 用于在切片后面追加元素
	s1 = append(s1, 100, 200, 300)
	fmt.Println(s1)

	//将切片内容增加到切片
	s2 = append(s1, s2...)
	fmt.Println(s2)

	//切片遍历
	for i := range s2 {
		fmt.Println(i)
	}
}
