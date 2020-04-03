package main

import (
	"fmt"
	"sort"
)

func main(){
	//创建map的方式
	//var map1 = map[int]string
	var map2 = make(map[string]int)
	var  map3  = map[string]int{"eng":99,"chin":89,"math":00}
	//fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	map2["yuwen"] = 99
	map2["shuxue"] = 88
	map2["waiyu"] = 11
	fmt.Println(map2)

	//判断map["pink"]是否存在
	v,ok := map2["pink"]
	if ok {
		fmt.Println(v)
	}else {
		fmt.Println("该值不存在，获取到的是默认值")
	}

	//修改map
	map2["shuxue"] = 100
	fmt.Println(map2)

	//删除map
	delete(map2,"yuwen")
	fmt.Println(map2)

	//map长度
	fmt.Println(len(map2))

	//遍历map map无序
	for k,v := range map2{
		fmt.Println(k,v)
	}

	//有序输出 按照value（key)排序
	var keys = make([]int,0,len(map2))
	fmt.Println(keys)

	for _,v := range map2{
		keys = append(keys,v)
	}
	sort.Ints(keys) //给int类型排序
	fmt.Println(keys)


	var m1 = map[string]string{"age":"11","name":"zhagnsan"}
	var m2 = map[string]string{"age":"22","name":"lisi"}
	var m3  = map[string]string{"age":"33","name":"wangwu"}

	var s = make([]map[string]string,0,3)
	s = append(s,m1)
	s = append(s,m2)
	s = append(s,m3)

	for i,value := range s{
		fmt.Printf("第%d个人的信息是：\n",i+1)
		fmt.Printf("姓名是%s\n年龄是%s\n",value["name"],value["age"],)
	}
}
