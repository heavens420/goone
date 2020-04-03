package main

import "fmt"
/**
		定义变量
 */

func main()  {
	fmt.Println("hello go")

	var num int = 1
	fmt.Println(num)

	var name = "nihaoa "
	fmt.Println(name)

	addr := "beijing"
	fmt.Println(addr)

	var a,b,c = 1,2,"nihaoa zheshijie"
	fmt.Println(a,b,c)

	var (
		age = 18
		phone = 23141
		ss = "nnnn"
	)
	fmt.Println(age,phone,ss)

	//定义常量  常量可以未使用 但不能修改  赋值方式与变量一致
	const PATH  = "kkkkkk"

	//枚举类型 枚举类型 可看作一组常量
	const(
		Q = 11
		S = 99
		d  //未赋值 默认为上面常量的值
	)

	//特殊常量 iota 可修改
	const (
		r = iota //iota = 0 defaut
		t        //iota = 1  default inc
		y = "iii" //iota = 2 default inc ...
		u		 //iota = 3
	)
	const o = iota //iota = 0 default

	fmt.Println(r,t,y,u,o)

}
