package main

import "fmt"

func main()  {
	var arr[4] int
	arr[0] = 1
	arr[3] = 5

	var arr2 = [3] int{1,2,3}
	arr2[2] = 99

	//下标为1 的数值为1 下表为3的数值为4
	var arr3 = [7] int{1:1,3:4}
	fmt.Println(arr3)

	//println(len(arr3))
	//获取数组的最大容量
	//println(cap(arr3))

	var str = [5]string{"3rw","r2r","54"}
	var str2 = [...]string{"424","35","535"}
	fmt.Printf("str:%s\n",str)
	fmt.Printf("Str2:%s\n",str2)
	fmt.Println(len(str))
	fmt.Println(len(str2))
	fmt.Println(cap(str))
	fmt.Println(cap(str2))

	for i := 0;i < len(str) ;i++  {
		fmt.Println(str[i])
	}

	for index,value := range arr{
		fmt.Printf("%d,%d\t",index,value)
	}
	fmt.Println()
	//遍历string
	for index,value := range str{
		fmt.Printf("%d,%s\t",index,value)
	}

	fmt.Println()
	//冒泡排序
	var sort = [...]int{2,425,6,34,5433,453}
	for i := 1; i<len(sort);i++{
		for j := 0;j<len(sort) - i;j++{
			if sort[j] > sort[j+1] {
				sort[j+1],sort[j] = sort[j],sort[j+1]
			}
		}
	}
	fmt.Println(sort)
}
