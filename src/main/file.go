package main

import (
	"fmt"
	"os"
)

func main() {
	filestatus()
	mikdir()
	//createfile()
	openfile()
	remoefile()
	removedir()
}

func filestatus() {
	fileInfo, err := os.Stat("C:\\Workplaces\\go\\w1\\goone\\src\\main\\time.go")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fileInfo)
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Sys())
	fmt.Println(fileInfo.Mode()) //文件权限
}

func mikdir() {
	dirs := "C:\\Workplaces\\go\\w1\\goone\\src\\a/b/c/d/e/f/g/h"
	dir := "C:\\Workplaces\\go\\w1\\goone\\ll"
	err := os.Mkdir(dir, os.ModePerm) //只能创建一级文件夹
	err2 := os.MkdirAll(dirs, os.ModePerm)
	if err != nil || err2 != nil {
		fmt.Println(err)
		fmt.Println(err2)
		return
		fmt.Println("create success")
	}
}

func createfile() {
	//绝对路径创建文件
	path := "C:\\Workplaces\\go\\w1\\goone\\src\\a/cc.txt"
	fi, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
		return
	}

	//相对路径创建文件
	file, err := os.Create("./nn.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(fi, "创建成功")
	fmt.Println(file, "创建成功")
}

func openfile() {
	file, err := os.OpenFile("./src/main/file.go", 2, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file, "打开成功")
}

func remoefile() {
	err := os.Remove("./nn.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file删除成功")
}

func removedir() {
	err := os.RemoveAll("./src/a")
	if nil != err {
		fmt.Println(err)
		return
	}
	fmt.Println("dir删除成功")
}
