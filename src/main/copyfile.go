package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//srcfile := "ll/a.txt"
	//desfile := "ll/b.txt"
	pic1 := "ll/pic.png"
	pic2 := "ll/pic2.png"
	//n, err := copytext(srcfile, desfile)
	//n, err := copytext(pic1, pic2)
	//n, err := copytext2(pic1, pic2)
	n, err := copytext3(pic1, pic2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n)
}

//最原始方法
func copytext(srcfile, desfile string) (int, error) {
	file1, err := os.Open(srcfile)
	if err != nil {
		fmt.Println(err)
	}
	file2, err := os.OpenFile(desfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, os.ModePerm)
	bt := make([]byte, 1024, 1024)

	var total = 0
	for {
		n, err := file1.Read(bt)
		if err == io.EOF || n == 0 {
			fmt.Println("copy over")
			break
		} else if err != nil {
			fmt.Println(err)
			return total, err
		}
		total += n
		file2.Write(bt[:n])
	}
	defer file1.Close()
	defer file2.Close()
	return total, err
}

//常用方法
func copytext2(file1, file2 string) (int64, error) {
	f1, err := os.Open(file1)
	if nil != err {
		fmt.Println(err)
		return 0, err
	}
	f2, err := os.OpenFile(file2, os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if nil != err {
		fmt.Println(err)
		return 0, err
	}
	defer f1.Close()
	defer f2.Close()
	return io.Copy(f2, f1)

}

//此方法不适用于拷贝较大文件
func copytext3(file1, file2 string) (int, error) {
	bytes, err := ioutil.ReadFile(file1)
	if nil != err {
		fmt.Println(err)
		return 0, err
	}

	err = ioutil.WriteFile(file2, bytes, 0777)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	return len(bytes), nil
}
