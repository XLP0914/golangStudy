package gobasic

import "fmt"

func main36() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			//more是状态
			if more {
				fmt.Println("received job", j)
				fmt.Println("received more", more)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
	fmt.Println("received more done:", done)
	_, ok := <-jobs
	fmt.Println("received more jobs:", ok)
}
