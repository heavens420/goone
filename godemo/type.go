package main

import "fmt"

/**
	基本数据类型
 */

func main()  {

	var flag bool = true
	fmt.Printf("%T,%t\n",flag,flag)

	//int8 int32 int64 值范围不同
	// uint8 uint32 uint64 无符号整数
	var  i32 int32 = 22222
	fmt.Printf("%T,%d\n",i32,i32)

	//默认为int
	var i  = 8888888888888
	fmt.Printf("%T,%d\n",i,i)

	//默认四舍五入
	var f float32 = 3.665554
	fmt.Printf("%T,%.2f\n",f,f)

	//默认float64
	var ft = 4.5436456
	fmt.Printf("%T,%f\n",ft,ft)
}