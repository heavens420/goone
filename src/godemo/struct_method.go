package main

import "fmt"

//方法 与 函数

//方法: x.x
//函数：x(x)

func main() {

	w := Work{"kk", 22, "nan"}
	w.work()
	w.rest()
	w.speak()
	eat(w)
}

type Work struct {
	name string
	age  int
	sex  string
}

func (w Work) work() {
	fmt.Println(w.name + " is working")
}

func (w Work) rest() {
	fmt.Println(w.name + " is resting")
}

func (w *Work) speak() {
	fmt.Println(w.name + " is speaking")
}
func eat(w Work) {
	fmt.Println(w.name + " is eatting")
}
