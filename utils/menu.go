package utils

import "fmt"

func PrintMenu() int {
	fmt.Println("\n=== MENU ===")
	fmt.Println("1. Danh sách lớp sort theo name ASC")
	fmt.Println("2. Danh sách học sinh sort theo name DESC")
	fmt.Println("3. Danh sách lớp học sinh X đã tham gia")
	fmt.Println("4. Danh sách giáo viên với filter")
	fmt.Println("0. Thoát")
	return PromptInt("Chọn: ")
}

func PrintTeacherFilterMenu() int {
	fmt.Println("\nChọn filter:")
	fmt.Println("1. Tất cả giáo viên")
	fmt.Println("2. Giáo viên là chủ nhiệm")
	fmt.Println("3. Giáo viên chủ nhiệm có trên X học sinh")
	fmt.Println("4. Giáo viên chủ nhiệm trên X lớp")
	fmt.Println("5. Thoát")
	return PromptInt("Chọn: ")
}
