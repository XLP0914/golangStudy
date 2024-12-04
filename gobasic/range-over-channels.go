package gobasic

import "fmt"

func main37() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	//如果不close，使用range会死锁
	for elem := range queue {
		fmt.Println(elem)
	}
}
