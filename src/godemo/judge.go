package main

import (
	"fmt"
)

func main() {
	var i = 100
	if i > 20 {
		fmt.Println("jjjjj")
	} else {
		fmt.Println("<<<<<")
	}
	fmt.Println("over")

	//go 独有写法 定义局部变量
	if i := 2; i > 3 {
		fmt.Println(">>>>2222")
	} else {
		fmt.Println("<<<<<<<<<<<<2222222")
	}

	var ch = 2
	switch ch {
	case 1:
		fmt.Println("1111111")
	case 2:
		fmt.Println(22222)
	case 3:
		fmt.Println(3333)
	default:
		fmt.Println("others")
	}

	switch { //switch 省略变量 相当于 变量为true
	case true:
		fmt.Println("true")
	case false:
		fmt.Println("false")
	}

	var score int = 88
	switch {
	case score <= 60:
		fmt.Println("bujige")
	case score <= 80:
		fmt.Println("bu优秀")
	case score <= 90:
		fmt.Println("jiang youxiu")
	case score <= 100:
		fmt.Println("太优秀")
	case score == 100:
		fmt.Println("陈独秀")
	}

	var mon int
	var year int
	var sum int
	var day int

	//输入年月日
	fmt.Println("输入年月日，空格隔开")
	fmt.Scanln(&year, &mon, &day)
	fmt.Println(year, mon, day)
	//嵌套 死循环 即可

	switch mon {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31天")
	//case 2 :
	//	fmt.Println("special month")
	case 2, 4, 6, 11:
		fmt.Println("30天")
		//穿透  直接执行后面的case 无需判断条件
		fallthrough
	case 0:
		fmt.Println("circle over")
	default:
		fmt.Println("error month")
	}

	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		//闰年减 一天
		sum -= sum
		sum += day
		fmt.Println("yeap")
	} else {
		//平年减 两天
		sum -= 2
		sum += day
	}
}
