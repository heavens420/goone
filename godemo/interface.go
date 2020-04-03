package main

import (
	"fmt"
)

func main()  {
	var m = Mose{"luoji"}
	var f = Flash{"sanxing"}

	interfaceImpl(m)
	interfaceImpl(f)

	var usb USB
	usb = m
	interfaceImpl(usb)
}

type Mose struct {
	name string
}

type Flash struct {
	name string
} 

//定义接口
type USB interface {
	//定义接口中的方法
	start()
	end()
}

//实现方法 不需要指明谁来实现
func (m Mose) start()  {
	fmt.Println("mouse begin ...")
}

func (m Mose) end() {
	fmt.Println("mouse over...")
}

func (f Flash)start() {
	fmt.Println("flash begin...")
}

func (f Flash) end() {
	fmt.Println("flash over ...")
}

//测试实现方法
func interfaceImpl(usb USB)  {
	usb.start()
	usb.end()
}