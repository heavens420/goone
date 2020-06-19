package main

import "fmt"

//type的使用

//定义结构体
//type stu struct {
//	name stirng
//	age int
//}

func main() {
	//自定义变量类型
	type myint int
	var b int
	var d myint
	fmt.Printf("%T,%T", d, b)
}

//定义函数类型
type myfun func(int, int) string

func fun1() myfun {

}

//定义别名 给int起个别名
type myint2 = int
