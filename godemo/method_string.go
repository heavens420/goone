package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main(){
	var name  = "nihaoazheshijie"
	//字符串中是否包含指定字符串
	fmt.Println(strings.Contains(name,"ee"))

	//包含其中任意一个字符 返回true
	fmt.Println(strings.ContainsAny(name,"ifaknvae"))

	//以xxx开头，xxx结尾
	fmt.Println(strings.HasPrefix(name,"ni"))

	fmt.Println(strings.HasSuffix(name,"kk"))

	//查找第一个出现的位置  找不到返回-1
	fmt.Println(strings.Index(name,"i"))

	//查找指定字符串中任意一个字符的位置
	fmt.Println(strings.IndexAny(name,"fsf"))

	//将字符串转换为 int
	i,err := strconv.Atoi("342")
	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Println(i)
}
