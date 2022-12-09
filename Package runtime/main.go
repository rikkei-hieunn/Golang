package main

import "runtime"

// GOGC đặt giá trị phần trăm thu gom rác ban đầu là GOGC=100
//GOGC = off sẽ tắt hoàn toàn quá trình thu gom rác
// func SetGCPercent(percent int) int cho phép thay đổi tỷ lệ % này trong thời gián chạy

func SetGCPercent(percent int) int {
	//nếu percent == âm thì vô hiệu hóa việc thu gom rác
}

//GOMAXPROCS giới hạn số luồng hệ điều hành có thể thực thi
func GOMAXPROCS(n int) int {
	// nếu n<1 thì hàm này sẽ k được thực thi
}

//NumCPU trả về số lượng CPU có thể sử dụng
func NumCPU() int {
}

func MaxParallelism() int {
	maxProcs := runtime.GOMAXPROCS(0)
	numCPU := runtime.NumCPU()
	if maxProcs < numCPU {
		return maxProcs
	}
	return numCPU
}

func main() {
	MaxParallelism()
}
