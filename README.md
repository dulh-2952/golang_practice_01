# Quản Lý Học Sinh - Golang Console App

## 🎥 Video Demo

Xem video demo trên Youtube: [Link Demo](https://youtu.be/KBt9bPZeme4)

Lưu ý có nhiều nội dung nhạy cảm từ 4 năm trước =.=

---

## 🚀 Hướng Dẫn Chạy

- **Yêu cầu:** Golang version `go1.22.4`
- Clone dự án về máy và mở terminal tại thư mục project.

### 1. Seed Dữ Liệu

```bash
cd seed
go run main.go
```

> Lệnh trên sẽ tạo file `data.json` chứa dữ liệu mẫu.

### 2. Chạy Chương Trình

```bash
go run main.go
```

---

## 📁 Cấu Trúc Thư Mục

```
GOLANG_PRACTICE_01/
├── model/                 # Chứa các struct dữ liệu (Class, Student, Teacher, ...)
├── seed/                  # Sinh file data.json mẫu
│   ├── data.json
│   └── main.go
├── service/print/         # Các hàm liên quan đến in dữ liệu
│   └── print_service.go
├── utils/                 # Các tiện ích (input, sort, loading...)
│   ├── input.go
│   ├── sort_data.go
│   └── with_loading.go
├── go.mod
└── main.go                # Điểm bắt đầu chương trình
```

---

## 🧩 Chức Năng

Chương trình quản lý học sinh thông qua console terminal, người dùng có thể chọn các tùy chọn sau:

### 1. 📚 Danh Sách Lớp (Sort theo tên ASC)

- Hiển thị: `class_id`, `name`, `Tên giáo viên chủ nhiệm`

### 2. 👨‍🎓 Danh Sách Học Sinh (Sort theo tên DESC)

- Hiển thị: `id`, `name`, `address`, v.v... (tất cả các thông tin học sinh)

### 3. 🏫 Danh Sách Lớp Học Sinh Đã Tham Gia

- Nhập `student_id` từ terminal
- Hiển thị các lớp mà học sinh đã tham gia

### 4. 👩‍🏫 Danh Sách Giáo Viên Theo Filter

- Tùy chọn filter:
  - `all` – Tất cả giáo viên
  - `chủ nhiệm` – Chỉ giáo viên là chủ nhiệm lớp
  - `có trên X học sinh` – Nhập X từ terminal
  - `chủ nhiệm trên X lớp` – Nhập X từ terminal

---

## 🛠 Tác giả

- Tên: Lê Huy Du
- Github: https://github.com/dusainbolt
- Facebook: https://www.facebook.com/dusainbolt
