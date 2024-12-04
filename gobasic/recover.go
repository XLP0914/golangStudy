package gobasic

import "fmt"

func mayPanic() {
	panic("a problem")
}

func main49() {

	//recover 的主要用途是捕获程序中的异常，让程序从 panic 中恢复，继续正常执行
	defer func() {
		if r := recover(); r != nil {

			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	mayPanic()

	//不会执行,因为前面 panic 了,而 defer后面 recover了
	fmt.Println("After mayPanic()")
}
