package main

import (
	"fmt"
	"math/rand"
	"time"
)

//Thuật toán Selection Sort sắp xếp một mảng bằng cách liên tục tìm phần tử nhỏ nhất (xét theo thứ tự tăng dần)
//từ phần chưa được sắp xếp và đặt nó ở đầu. Thuật toán duy trì hai mảng con trong một mảng nhất định.
//1. Mảng con đã được sắp xếp.
//2. Mảng con còn lại chưa được sắp xếp.

func main() {

	slice := generateSlice(20)
	fmt.Println("\n--- Unsorted --- \n\n", slice)
	selectSort(slice)
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

func selectSort(slice []int) {

	for i, _ := range slice {
		minValue := i
		for j := i; j < len(slice); j++ {
			if slice[minValue] < slice[j] {
				minValue = j
			}
			slice[minValue], slice[j] = slice[j], slice[minValue]
		}
	}
}
