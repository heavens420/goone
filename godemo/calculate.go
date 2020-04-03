package main

import "fmt"

/**
	运算符
 */

func main()  {
	var i = 19
	var j = 2
	var sum int
	sum = i + j
	fmt.Printf("%d+%d = %d\n",i,j,sum)

	var div  = i % j
	fmt.Printf("%d %% %d = %d\n",i,j,div)

	a := 3
	b := 5
	a++
	b--
	//--a 不支持++a --a 写法
	//var  d  = a++ + b++  自增自减不能参与运算

	fmt.Printf("%d,%d\n",a,b)

	//位清空运算 &^
	p := 3
	q := 4
	k := p &^ q
	fmt.Printf("%d &^ %d = %d\n",p,q,k)
	g := q &^ p
	fmt.Printf("%d &^ %d = %d",q,p,g)

	var input int
	var output int
	fmt.Print("请输入两个整数\n")
	fmt.Scanf("%d,%d\n",&input,&output)
	fmt.Println(input,output)
}