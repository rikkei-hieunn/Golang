package main

import "fmt"

// syntax khai báo Channel

var channelName chan int
//or
channelName2 := make(chan int)

// để gửi và nhận dữ liệu qua channel chúng ta sử dụng toán tử <-
//Gửi
//channelName <- value

//Nhận
//myVar := <- channelName

//Dưới đây là cách sử dụng Channel chỉ gửi
func recieveOnly(c <-chan int) {
	fmt.Printf("Received: %d\n", <-c)
//	c <- 2 // error
}

//Tương tự đây là cách sử dụng Channel chỉ nhận
func sendOnly(c chan<- int) {
	c <- 2 // OK
//	fmt.Printf("Received: %d\n", <-c) // error
}

//Case 1: Channel c có thể tiếp nhận dữ liệu x (không bị block), thực hiện phép toán x = y và y = x + y.
//Case 2: Channel quit có dữ liệu, return thoát hàm luôn.
func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}


func main() {
	myChan := make(chan int)

	go func() {
		// gửi giá trị 1 vào channel
		myChan <- 1
	}()

	// nhận giá trị channel và in ra màn hình
	// Chú ý khi sử dụng channel việc gủi và nhận phải được diễn ra nếu chỉ có mình gửi giữ liệu thì chương trình sẽ bị dealook
	// Trường hợp này xảy ra khi k còn goroutine nào lấy dữ liệu từ channel
	fmt.Println(<-myChan)

	go recieveOnly(myChan)

	go sendOnly(myChan)

	// Để đóng một Channel chúng ta sẽ dùng hàm close().
	//Bình thường nếu chỉ lấy dữ liệu từ channel đơn thuần thì chúng ta không biết được channel đã đóng hay vẫn còn hoạt động
	value, isAlive := <-myChan
	if !isAlive {
		fmt.Printf("Value: %d. Channel has been closed.\n", value)
	}

	//cách sử dụng Select để chọn Channel
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)


}