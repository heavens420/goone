package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var i = rand.Int31n(34242)
	fmt.Println(i)

	//获取当前系统时间
	var t = time.Now()
	fmt.Println(t)

	//将时间转换为秒
	var second = t.Unix()
	fmt.Println(second)

	var ra = rand.Int63n(second)
	fmt.Println(ra)

	//设置时间种子
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		//生成随机数
		rand := rand.Int63()
		fmt.Println(rand)
	}

	//获取指定区间的随机数 [20,120)
	var s = rand.Int63n(100) + 20
	fmt.Println(s)
}
