package gobasic

import (
	"errors"
	"fmt"
)

// 自定义错误
type argError struct {
	arg     int
	message string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.message)
}

func f(arg int) (int, error) {
	if arg == 42 {
		//两个返回值，-1及 结构体 argError
		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}

func main28() {

	_, err := f(42)
	var ae *argError
	if errors.As(err, &ae) {
		fmt.Println(ae.arg)
		fmt.Println(ae.message)
	} else {
		fmt.Println("err doesn't match argError")
	}
}
