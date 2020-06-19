package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	myerror()
	myerror2()
	checkAge(-20)

	openFile("C:\\Workplaces\\go\\w1\\goone\\godemo\\array.go")
}

func myerror() {
	f, err := os.Open("C:\\Workplaces\\go\\w1\\goone\\godemo\\array.go")
	if err != nil {
		log.Fatal(err)
		fmt.Println(err, "ooo")
	}
	fmt.Printf(f.Name(), "文件打开成功\n")
}

func myerror2() {
	err1 := errors.New("this is my error")
	fmt.Println(err1)

	err2 := fmt.Errorf("another error\n ")
	fmt.Println(err2)
}

func checkAge(age int) error {
	if age < 0 {
		er := errors.New("输入有误")
		er2 := fmt.Errorf("输入错误")
		fmt.Println(er2)
		fmt.Println(er)
		return er
	}
	fmt.Println(age)
	return nil
}

func openFile(path string) {
	files, error3 := filepath.Glob(path)
	if error3 != nil && error3 == filepath.ErrBadPattern {
		fmt.Println(error3)
		return
	}
	fmt.Println(files)
}
