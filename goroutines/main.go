package main

import (
	"fmt"
	"runtime"
	"time"
)

//	Goroutine bản chất là một lightweight execution thread
//	Goroutines có kích thước nhỏ hơn rất đáng kể so với Thread. Goroutines sử dụng 2KB memory stack, trong khi đó các OS Thread lên đến 2MB.
//	Goroutines có thể linh động tăng giảm bộ nhớ sử dụng trong khi đó OS Thread là cố định.
//	Một chương trìng Golang có thể có hàng trăm nghìn Goroutine trong khi Thread chỉ được vài trăm đến hàng nghìn.
//	Goroutines có thời gian khởi động nhanh hơn Thread.
//	Goroutines có thể giao tiếp an toàn với nhau thông qua các kênh trong Golang (Channel).
//	Các channel hỗ trợ mutex lock vì thế tránh được các lỗi liên quan tới cùng ghi và đọc lên vùng dữ liệu chia sẻ (data race).
//	Goroutines có thể được ánh xạ và hoạt động trên nhiều OS threads thay vì là ánh xạ 1:1.

//Cách khai báo một Goroutine - ( go nameFunction)

func name() {
	// statements
}
func main() {
	go name()

	// capture variable trong Goroutine Golang
	//vòng lặp for khi print ra giá trị của i sẽ bị trùng lặp
	// vì i trong func đang là 1 tham chiếu , không phải giá trị, các goroutines không được thục hiện ngay dẫn đến sự thay đổi của i
	for i := 1; i <= 100; i++ {
		go func() {
			fmt.Println(i) // biến i ở đây là một pointer
		}()
	}
	// để tránh điều này t có thể copy từng giá trị của i
	for i := 1; i <= 100; i++ {
		go func(value int) {
			fmt.Println(value) // value ở đây độc lập với i ở ngoài
		}(i) // value i được copy ở đây
	}

	//Sử dụng Gosched()
	// khi thực hiện vòng lặp với nhiều goroutine có thể xẩy ra việc chiếm xử lý
	// để tránh việc này và đảm bảo tất cả các vòng lặp đều được thực hiện như nhau
	// Sử dụng hàm Gosched() của package runtime để schedule goroutine ngay lập tức sau mỗi vòng lặp.
	go func() {
		for i := 1; i <= 50; i++ {
			fmt.Println("I am Goroutine 1")
			runtime.Gosched()
		}
	}()

	go func() {
		for i := 1; i <= 50; i++ {
			fmt.Println("I am Goroutine 2")
			runtime.Gosched()
		}
	}()
	time.Sleep(time.Second)

}
