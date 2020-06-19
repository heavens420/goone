package main

import (
	"fmt"
)

//go 中只有 for循环

func main() {
	var i = 1
	var sum = 0
	//for 简写
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println(sum)

	//for标准写法
	for i := 1; i < 5; i++ {
		fmt.Println(i)
	}

	//无条件 死循环 相当于while(true)
	for {
		i++
		if i == 100000 {
			break
		}
		fmt.Println(i)
	}

	//乘法表
	for k := 1; k <= 9; k++ {
		for j := 1; j <= k; j++ {
			fmt.Printf("%d * %d = %d\t", k, j, k*j)
		}
		fmt.Println()
	}

	//完全数
	for i := 2; i <= 1000; i++ {
		var sum1 = 1
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				var y = i / j
				sum1 += y
				//fmt.Printf("yyy %d\n",y)
				//fmt.Printf("kkk %d\n",sum1)
			}
		}
		if sum1 == i {
			fmt.Printf("%d 是完全数\n", i)
		}
	}

	//水仙花数
	for i := 100; i < 1000; i++ {
		var sum2 = 0
		var y = i % 10
		var yy = i % 100 / 10
		var yyy = i / 100
		sum2 += y*y*y + yy*yy*yy + yyy*yyy*yyy
		if sum2 == i {
			fmt.Printf("%d 是水仙花数\n", i)
		}
	}
}
