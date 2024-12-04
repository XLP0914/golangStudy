package gobasic

import (
	"fmt"
	"time"
)

func main57() {

	now := time.Now()
	fmt.Println(now)
	// 获取当前时间的 Unix 时间戳（秒）
	fmt.Println(now.Unix())
	fmt.Println(now.UnixMilli())
	fmt.Println(now.UnixNano())

	// 使用 Unix 时间戳恢复时间
	fmt.Println(time.Unix(now.Unix(), 0)) // 只传递秒，不带纳秒
	fmt.Println(time.Unix(0, now.UnixNano()))
}
