package main

import "fmt"

func main()  {

	var arr = make([]int,10)
	for i := 0;i<10 ;i++  {
		arr[i] = i
	}
	var arr2 = make([]int,20)
	//将arr中的元素拷贝到arr2
	copy(arr2,arr)
	fmt.Println(arr2)
	fmt.Println(arr)

	//从第七个元素开始拷贝 默认覆盖前面
	copy(arr2,arr[7:])
	fmt.Println(arr2)
}
