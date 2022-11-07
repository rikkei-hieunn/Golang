package main

// thuật toán sắp xếp bubble sort sử dụng 2 vòng for để duyệt phần tử trong array với các phần từ còn lại để tìm ra giá trị bé nhất sau mỗi vòng lặp
// đô phức tạo 0(n2)

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	bubbleSort(slice)
	fmt.Println("\n--- Sorted ---\n\n", slice, "\n")
}

// Tạo ra 1 slice ngẫu nhiên
func generateSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func bubbleSort(items []int) {
	var n = len(items)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if items[i] > items[j] {
				items[i], items[j] = items[j], items[i]
			}
		}
	}
}
