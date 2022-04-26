package main

import "fmt"

func main() {
	test4()
	test5()
}

func test4() {
	a := make(map[string]int, 2)
	a["haha"] = 100
	fmt.Println(a)

	// map中是否存在某个key值 v为value
	v, ok := a["haha1"]
	fmt.Println(v, ok)
}

// new 类似make  new返回值的指针  make返回对象值本身
func test5() {
	b := new(int)
	*b = 10
	fmt.Println(*b)
}
