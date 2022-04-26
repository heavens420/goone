package main

import "fmt"

func main() {
	myTest1()
}

func myTest1() {
	aa := 1
	// 不满5位填充至5位 默认用空格填充 这里指定用0填充
	key := fmt.Sprintf("zifuchuan:%05d", aa)
	fmt.Println(key)
}
