package main

import "fmt"

/**
panic函数
recover函数
*/

func main() {
	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg, "恐慌结束了")
		}
	}()
	AA()
	defer B()
	D()
	panic("恐慌出现\n") //panic前的函数可以执行 之后的函数不执行
	defer E()()     //正常执行 但return的内容还是defer执行
	C()
}

func AA() {
	fmt.Println("A函数")
}

func B() {
	fmt.Println("B函数")
}

func C() {
	fmt.Println("C函数")
}

func D() {
	defer B()
	fmt.Println("D 函数")
	defer C()
	AA()
}

func E() func() {
	fmt.Println("EEEE")
	return func() {
		fmt.Println("rrr")
	}
}
