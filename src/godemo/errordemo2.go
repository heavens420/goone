package main

import "fmt"

type myError struct {
	msg    string
	width  float64
	height float64
}

//错误信息
func (e *myError) Error() string {
	return e.msg
}

//判断宽
func (e *myError) widthError() bool {
	return e.width > 0
}

func (e *myError) heightError() bool {
	return e.height > 0
}

func area(width, height float64) (float64, error) {
	msg := ""
	if width < 0 {
		msg = "width lt 0"
	}
	if height < 0 {
		if msg == "" {
			msg = "height is lt 0"
		} else {
			msg += " and height is lt 0,too"
		}
	}

	if msg != "" {
		return 0, &myError{
			msg:    msg,
			width:  width,
			height: height,
		}
	}
	return width * height, nil
}

func main() {
	ss, err := area(2, 13)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("the area is ", ss)
}
