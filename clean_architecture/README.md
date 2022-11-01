#Clean architecture

### Sơ đồ tổng quát
![alt text](clean-arch.png)

### Khái niệm

- Clean Architecture là một kiến trúc tổ chức source code thành các Classes hoặc Files hoặc Components hoặc Modules và chúng sẽ không phải phụ thuộc vào lẫn nhau mà sẽ phụ thuộc vào những thứ mang tính trìu tương như interface,...

### Lợi ích của việc triển khai Clean Architecture

- Thực hiện việc test dễ dàng hơn: khi cách tổ chức code được chia thành các phần riêng biệt và không bị ràng buộc lẫn nhau nên việc viết Unit Test cũng sẽ được áp dụng cho từng phần riêng biệt, giảm thiểu được việc Unit Test phần này phải gọi đến function của phần khác
- Tách biệt được các Framework, Library: trong hệ thống việc sử dụng các Framework và các Library là không tránh khỏi, Clean Architecture giúp các thành phần khác không bị ràng buộc vào Framework cũng như Library nên việc thay đổi Framework hay sử dụng Library khác cũng không làm ảnh hưởng đến hệ thống hiện tại
- Tách biệt được thành phần giao tiếp với UI: cũng giống với việc tách biệt Framework và Library, thành phần tương tác với UI được tách riêng biệt nên việc có thay đổi thành phần này cũng không làm ảnh hưởng đến logic của hệ thống

### Một kiến trúc Clean Architecture bao gồm 4 layer:

- Models Layer: Là nơi chứa tất cả các models/entities sử dụng trong toàn bộ source code

    Ví dụ: Một hệ thống quản lý thông tin sinh viên tiếp nhận các thống tin như id, name, age của sinh viên từ Client và thực hiện xử lý các thông tin đó rồi lưu vào Database
    
    Đối với bài toán quản lý sinh viên, sinh viên là một thực thể ngoài đời thật bao gồm các thông tin id, name, age thì trong các xử lý logic trong source code hoàn toàn có thể truyền cả 3 thuộc tính này qua các functions nó dẫn đến việc riêng rẽ về mặt data nên để gom nhóm data để dễ quản lý và dễ sử dụng thì một Struct Student được tạo ra để gom 3 thuộc tính id, name, age của một sinh viên vào trong Struct Student.
```golang
    type Student struct {
	    id   int,
	    name string,
	    age  string,
    }   
```

- Repository Layer: Repository sẽ lưu trữ bất kỳ trình xử lý Cơ sở dữ liệu nào. Truy vấn hoặc Tạo / Chèn vào bất kỳ cơ sở dữ liệu nào sẽ được lưu trữ tại đây. Lớp này sẽ chỉ hoạt động đối với CRUD đối với cơ sở dữ liệu. Không có quy trình kinh doanh nào xảy ra ở đây. Chỉ có chức năng đơn giản đối với Cơ sở dữ liệu.

