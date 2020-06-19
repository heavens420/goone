package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	read()
}

func read() {
	path := "src/main/file.go"
	//1 打开文件
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	bs := make([]byte, 1000, 1004)
	//3 关闭文件
	defer file.Close()

	//2 读取文件
	for {
		n, err := file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("文件读取结束")
			break
		}
		fmt.Println(string(bs[:n]))
	}
}
