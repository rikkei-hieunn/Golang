package main

import "fmt"

// Golang không có class. Tuy nhiên, chúng ta có thể xác định methods on types.
// Nghĩa là đối với một kiểu dữ liệu cụ thể nào đó, ta hoàn toàn có thể xây dựng các functions cho kiểu dữ liệu đó
// Kiểu dữ liệu đó được gọi là receiver, còn các function là methods of receiver
// Cú pháp khai báo một method của receiver có cú pháp như sau:
// func (receiver) name(parameters) (returns) {
//	-> TODO logic here
// }

// Định nghĩa một kiểu tên là Names với giá trị là một slice string
type Names []string

// Names khi này là receiver và ta định nghĩa một method có chức năng hiện thị toàn bộ value của Names
func (n Names) print() {
	for _, name := range n {
		fmt.Println(name)
	}
}

// Định nghĩa một kiểu tên là Student với giá trị đại diện cho nó là một struct như sau
type Student struct {
	name string
	age  int
}

// Student khi này là receiver và ta định nghĩa một method có chức năng hiện thị name của Student
func (s Student) getName() string {
	return s.name
}

// Định nghĩa một method thực hiện gán lại giá trị age cho Student
func (s Student) setAge(newAge int) {
	s.age = newAge
}

// Định nghĩa một method thực hiện gán lại giá trị age cho Student với receiver là pointer
func (s *Student) setAgeWithPointer(newAge int) {
	s.age = newAge
}

func main() {
	// Khởi tạo biến firstNames với kiểu dữ liệu là Names
	firstNames := Names{"Trung", "Quang", "Le"}
	firstNames.print()

	// Khởi tạo biến firstStudent với kiểu dữ liệu là Student
	firstStudent := Student{
		name: "Le Quang Trung",
		age:  26,
	}
	fmt.Println("Name of first student:", firstStudent.getName())

	// Thực hiện gán lại age với method setAge()
	firstStudent.setAge(30)
	fmt.Println("Age of first student:", firstStudent.age)

	// Thực hiện gán lại age với method setAgeWithPointer()
	firstStudent.setAgeWithPointer(30)
	fmt.Println("Age of first student:", firstStudent.age)
	// Nhận thấy với method setAge() sẽ không làm thay đổi giá trị age cảu firstStudent do receiver ở trong method là một bản
	// copy của firstStudent nên việc thay đổi giá trị age là thay đổi giá trị của bản copy sẽ không làm ảnh hưởng bản gốc
	// Ngược lại với method setAgeWithPointer() sử dụng pointer receiver nên việc thay đổi giá trị age ở đây sẽ làm thay đổi
	// trực tiếp giá trị gốc
}
