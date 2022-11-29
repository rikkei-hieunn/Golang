package main

import (
	"fmt"
	"sync"
)

//waitgroup sinh ra để giải quyết vấn đề chờ goroutine thực hiện xong việc thì mới chạy tiếp
// b1 tạo ra 1 biến có kiểu dữ liệu là sync.WaitGroup :wg
// b2 wg.Add( số lượng goroutine cần phải đợi)
// b3  ở mỗi goroutine gọi method wg.Done() trước khi return
// b4 gọi method wg.Wait() - sẽ được gọi ở hàm mà phải chở các goroutine chạy xong

// để gọi được wg.Done() thì phải truyền vào biến wg có kiểu dữ liệu con trỏ *sync.WaitGroup
func runner1(wg *sync.WaitGroup) {
	defer wg.Done() // This decreases counter by 1
	fmt.Print("\nI am first runner")

}

func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Print("\nI am second runner")
}

func execute() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	// We are increasing the counter by 2
	// because we have 2 goroutines
	go runner1(wg)
	go runner2(wg)

	// This Blocks the execution
	// until its counter become 0
	wg.Wait()
}

func main() {
	// Launching both the runners
	execute()
}
