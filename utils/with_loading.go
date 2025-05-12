package utils

import (
	"fmt"
	"time"
)

func WithLoading(task func(), duration time.Duration) {
	done := make(chan struct{})

	// Goroutine hiển thị loading
	go func() {
		dots := ""
		for {
			select {
			case <-done:
				return
			default:
				fmt.Printf("\rVui lòng chờ kết quả%s", dots)
				dots += "."
				if len(dots) > 3 {
					dots = ""
				}
				time.Sleep(300 * time.Millisecond)
			}
		}
	}()

	// Delay trước khi thực hiện task
	time.Sleep(duration)
	close(done)

	// In dòng mới tách biệt loading và kết quả
	fmt.Printf("\rVui lòng chờ kết quả...\n")

	// Gọi hàm xử lý chính
	task()
}
