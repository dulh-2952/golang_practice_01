package printService

import (
	"fmt"
	"practice_01/model"
)

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
