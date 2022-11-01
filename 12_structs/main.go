package main

import "fmt"

// Struct trong Golang là một kiểu do người dùng định nghĩa cho phép nhóm / kết hợp các mục có thể thuộc nhiều kiểu khác nhau thành một kiểu duy nhất.
// Bất kỳ thực thể trong thế giới thực nào có một số thuộc tính / trường đều có thể được biểu diễn dưới dạng Struct.
// Khái niệm này thường được so sánh với các Class trong lập trình hướng đối tượng. Nó có thể được gọi là một Class không hỗ trợ kế thừa nhưng hỗ trợ thành phần.

// Định nghĩa một Struct
type Student struct {
	id   int
	name string
	age  int
}

func main() {
	// Khởi tạo một đối tượng từ Struct Student
	firstStudent := Student{
		id:   1,
		name: "Trung",
		age:  26,
	}

	secondStudent := Student{
		id:   2,
		name: "Quang",
	}

	fmt.Println("firstStudent", firstStudent)
	fmt.Println("secondStudent", secondStudent)

	// Cập nhật giá trị của các hạng mục trong Struct
	firstStudent.name = "Trung1"
	firstStudent.age = 25
	fmt.Println("firstStudent", firstStudent)

	// So sánh 2 Struct
	fmt.Println("firstStudent == secondStudent:", firstStudent == secondStudent)
	secondStudent.id = 1
	secondStudent.name = "Trung1"
	secondStudent.age = 25
	fmt.Println("firstStudent == secondStudent:", firstStudent == secondStudent)

	// Sử dụng struct nhưng không cần định nghĩa Struct từ trước (Anonymous Struct)
	firstTeacher := struct {
		name string
		age  int
	}{
		"Lalala",
		30,
	}
	fmt.Println("firstTeacher:", firstTeacher)

	// Nhúng Struct vào trong một Struct khác
	type ClassRoom struct {
		name    string
		student Student
	}
	
	firstClass := ClassRoom{
		name:    "KHMT",
		student: Student{10, "Le Quang Trung", 25},
	}

	fmt.Println("firstClass", firstClass)
}
