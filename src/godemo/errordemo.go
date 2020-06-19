package main

import (
	"fmt"
	"math"
)

//定义错误对象结构体
type areaError struct {
	msg    string
	radius float64
}

//实现Error接口 并自定义异常信息
func (r *areaError) Error() string {
	return fmt.Sprintf("error:半径 = %.2f,%s", r.radius, r.msg)
}

func circleArea(r float64) (float64, error) {
	if r < 0 {
		return 0, &areaError{msg: "数据非法", radius: r}
	}
	return math.Pow(r, 2) * math.Pi, nil
}

func main() {
	area, err := circleArea(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(area)
}
