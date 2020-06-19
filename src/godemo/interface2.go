package main

import "fmt"

func main() {
	var a1 = A("xiaomao")
	var a2 = A("xiaogou")
	var a3 = "nihaoa "
	var a4 = 888
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
	fmt.Println(a4)

	var a = A("jjj")
	test1(a)

	var b = A("kkk")
	test2(b)

	//可以接收任意类型的map
	var maps = make(map[string]interface{})
	maps["w"] = "nini"
	maps["j"] = "hffkhk"
	maps["l"] = Cat{"little"}
	fmt.Println(maps)

	//切片
	var s1 = make([]interface{}, 1, 10)
	s1[0] = "fdsffsf"
	s1 = append(s1, a1, a2, a3, a4, b)
	fmt.Println(s1)

	//函数
	test3(s1)
}

//定义一个空接口 可以接收任意类型变量
type A interface {
}

type Cat struct {
	name string
}

type Dog struct {
	name string
}

func test1(a A) {
	fmt.Println(a)
}

func test2(a interface{}) {
	fmt.Println(a)
}

func test3(s1 []interface{}) {
	for i := 1; i < len(s1); i++ {
		fmt.Printf("第%d个数据是%v\n", i, s1[i-1])
	}
}
