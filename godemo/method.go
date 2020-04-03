package main

import "fmt"

func sum(n int  ){
	var sum = 0
	for i:= 0;i<= n ;i++  {
		sum += i
		fmt.Println(sum)
	}
}

func sum2 (a,b int)(int)  {
	var sum = 0
	sum = a + b
	return sum
}

func ractangle(a,b,d int)(int,int,int)  {
	var s = a*b
	var c = a+b
	return s,c,d
}

func main()  {
	sum(10)
	var q = sum2(3,4)
	fmt.Println(q)
	a,b,d := ractangle(1,2,3)
	fmt.Println(a,b,d)

	//使用 _ 表示舍弃某个值
	i,_,_ := ractangle(2,3,4)
	fmt.Println(i)
}