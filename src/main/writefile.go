package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

}

func write() {
	path := "ll/a.txt"
	//file, err := os.Open(path)
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	handleError(err)
	defer file.Close()

	bytes := []byte{122, 24, 'A', 'g', 's', 'k'}
	n, err1 := file.Write(bytes)
	handleError(err1)
	fmt.Println(n)

	writeString, err := file.WriteString("nihaoa zheshijie ")
	handleError(err)
	fmt.Println(writeString)

	bt, err := file.Write([]byte("wohenhao zheshijie "))
	handleError(err)
	fmt.Println(bt)

}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

}
