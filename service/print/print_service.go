package printService

import (
	"fmt"
	"practice_01/model"
	"practice_01/utils"
)

func PrintMenu() int {
	fmt.Println("\n=== MENU ===")
	fmt.Println("1. Danh sách lớp sort theo name ASC")
	fmt.Println("2. Danh sách học sinh sort theo name DESC")
	fmt.Println("3. Danh sách lớp học sinh X đã tham gia")
	fmt.Println("4. Danh sách giáo viên với filter")
	fmt.Println("0. Thoát")
	return utils.PromptInt("Chọn: ")
}

func PrintTeacherFilterMenu() int {
	fmt.Println("\nChọn filter:")
	fmt.Println("1. Tất cả giáo viên")
	fmt.Println("2. Giáo viên là chủ nhiệm")
	fmt.Println("3. Giáo viên chủ nhiệm có trên X học sinh")
	fmt.Println("4. Giáo viên chủ nhiệm trên X lớp")
	fmt.Println("5. Thoát")
	return utils.PromptInt("Chọn: ")
}

func PrintClassesRelation(classes []model.Class, teachers []model.Teacher) {
	fmt.Println("======= DANH SÁCH LỚP HỌC =======")
	for _, c := range classes {
		c.PrintClassRelation(teachers)
	}
}

func PrintStudents(students []model.Student) {
	fmt.Println("======= DANH SÁCH HỌC SINH =======")
	for _, s := range students {
		s.PrintStudent()
	}
}

func PrintClassesOfStudent(studentID int, classes []model.Class) {
	fmt.Printf("======= Danh Sách Lớp Học Của Học Sinh - ID = %d =======\n", studentID)
	for _, c := range classes {
		c.PrintClass()
	}
}

func PrintTeachers(title string, teachers []model.Teacher) {
	fmt.Println("===", title, "===")
	for _, t := range teachers {
		fmt.Printf("ID: %d | Name: %s | Address: %s\n", t.ID, t.Name, t.Address)
	}
	fmt.Println()
}
