package main

import "fmt"

// The Computer Memory (RAM): có thể được coi như một chuỗi các ô, được đặt nối tiếp nhau trên một dòng.
// Mỗi ô được đánh dấu bằng một số duy nhất, số này là địa chỉ của ô hoặc vị trí bộ nhớ của nó

// var number int = 10
// Khi này ta tạo ra một nhãn là number đại diện cho địa chỉ ô nhớ nơi mà giá trị 10 thuộc kiểu int được lưu trữ
// Pointer là một biến lưu trữ địa chỉ ô nhớ của một biến khác
// Pointer sẽ bằng nil nếu biến chưa được khởi tạo

// Với biến number trên giá trị 10 được lưu tại ô nhớ có địa chỉ 0x1040a123 chẳng hạn, thì biến Pointer là biến trữ
// 0x1040a123

func main() {
	var number int = 10
	// Để sử dụng Pointer ta sử dụng ký tự &
	fmt.Println("Pointer of number:", &number)
	numberPointer := &number
	fmt.Println("Pointer of number:", numberPointer)
	fmt.Printf("Type of numberPointer: %T\n", numberPointer)

	var secondPointer *float64
	// Giá trị default của Pointer là nil
	fmt.Println("The default value of Pointer:", secondPointer)

	// Để thay đổi giá trị của biến thông qua Pointer, sử dụng *(Dereferencing Operators)
	secondPointer = new(float64)
	*secondPointer = 31.5
	fmt.Println("The value of secondPointer:", secondPointer, "The value stored at", secondPointer, "is", *secondPointer)
	// &value => pointer
	// *pointer => value

	// Sử dụng pointer của một pointer
	var secondNumber int = 20
	pointerOfNumber := &secondNumber
	pointerOfPointerNumber := &pointerOfNumber
	fmt.Printf("The value of secondNumber: %d, pointer of secondNumber: %v, pointer of pointer of secondNumber: %v\n", secondNumber, pointerOfNumber, pointerOfPointerNumber)
	// Thay đổi giá trị của biến qua 2 lần pointer
	**pointerOfPointerNumber = 40
	fmt.Printf("The value of secondNumber: %d, pointer of secondNumber: %v, pointer of pointer of secondNumber: %v\n", secondNumber, pointerOfNumber, pointerOfPointerNumber)

	// So sánh hai pointer
	// Hai pointer được coi là bằng nhau khi chúng đều được trỏ đến địa chỉ của cùng một biến
	firstPointer := &secondNumber
	thirdPointer := &secondNumber
	var fourthPointer *int
	fmt.Println("firstPointer == thirdPointer:", firstPointer == thirdPointer)
	// Ta có thể so sánh pointer với nil
	fmt.Println("firstPointer == nil:", firstPointer == nil)
	fmt.Println("fourthPointer == nil:", fourthPointer == nil)

	// Ở bài function khi ta truyền một giá trị qua thâm số vào một function thì giá trị đó sẽ được copy ra một bản rồi
	// được sử dụng trong function đó, khi thay đổi giá trị của tham số đầu vào đó sẽ không làm ảnh hưởng đến giá trị của
	// biến bên ngoài, để có thể thực hiện thay đổi giá trị của biến bên ngoài ta thực hiện việc truyền pointer của biến
	// đó vào trong function
	fifthNumber := 100
	fmt.Println("The value of fifthNumber:", fifthNumber)
	increase(&fifthNumber)
	fmt.Println("The value of fifthNumber:", fifthNumber)

	// Ta cũng có thể return về một pointer
	// Như đã biết các kiểu giá trị string có default value là "", chúng không thể so sánh với nil, ta sẽ sử dụng pointer của chúng
	var input int
	if result := convertIntToString(input); result == nil {
		fmt.Println("The input is 0")
	} else {
		fmt.Println("The string value of input:", result)
	}
}

func increase(number *int) {
	*number++
}

func convertIntToString(input int) *string {
	if input == 0 {
		return nil
	}
	result := fmt.Sprintf("%d", input)

	return &result
}
