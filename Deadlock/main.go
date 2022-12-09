package main

import (
	"fmt"
	"sync"
	"time"
)

// Deadlock là hiện tượng sẽ xảy ra khi sử dụng channel( việc sử dụng channel phải có đủ gửi và nhận dữ liệu )
// ví dụ

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	c := make(chan int)
	go readFromChannel(wg, c, time.After(time.Duration(2)*time.Second))
	go receiveFromChannel(wg, c, time.After(time.Duration(2)*time.Second))
	// Trong trường hợp sử dụng code ntn sẽ bị deadlook
	// Không sử dụng goroutine để nhận dữ liệu vào channel
	/*	time.Sleep(time.Duration(5) * time.Second)
		c <- 10*/
	//Giải thích: vì chúng ta kb chắc chắn khi nào channel sẽ nhận được data nên sẽ xảy ra dealock cụ thể trong trường hợp này là send
	wg.Wait()
}

func readFromChannel(wg *sync.WaitGroup, c chan int, ti <-chan time.Time) {
	select {
	case x := <-c:
		fmt.Println("Read", x)
	case <-ti:
		fmt.Println("TIMED OUT")
	}
	wg.Done()
}

func receiveFromChannel(wg *sync.WaitGroup, c chan int, ti <-chan time.Time) {
	c <- 10
	wg.Done()
}
