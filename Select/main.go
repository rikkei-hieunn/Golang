package main

import (
	"fmt"
	"sync"
)

//Select được dùng để lựa chọn giữa nhiều hành động nhận/ gửi của channel
//Select sẽ chọn ngẫu nhiên hành động nhận/gửi trong số đó
// việc sử dụng Select có thể xử lý được việc deadlock

// function 1
func portal1(wg *sync.WaitGroup, channel1 chan string) {
	channel1 <- "Welcome to channel 1"
	wg.Done()
}

// function 2
func portal2(wg *sync.WaitGroup, channel2 chan string) {
	channel2 <- "Welcome to channel 2"
	wg.Done()
}

// main function
func main() {
	wg := new(sync.WaitGroup)
	wg.Add(1)
	// Creating channels
	R1 := make(chan string)
	R2 := make(chan string)

	// calling function 1 and
	// function 2 in goroutine
	go portal1(wg, R1)
	go portal2(wg, R2)

	select {

	// case 1 for portal 1
	case op1 := <-R1:
		fmt.Println(op1)

	// case 2 for portal 2
	case op2 := <-R2:
		fmt.Println(op2)
	default:
		fmt.Println("default case executed")
	}
	wg.Wait()
}
