package gobasic

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	//tine也是进程执行后会影响执行顺序
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main40() {

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

	wg.Wait()

}
