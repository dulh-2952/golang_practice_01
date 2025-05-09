package main

import (
	"fmt"
	"os"
	"sort"

	"practice_01/model"

	"practice_01/utils"
)

func printClassASC(data model.Data) {
	sort.Slice(data.Class, func(i, j int) bool {
		return data.Class[i].Name < data.Class[j].Name
	})
	for _, c := range data.Class {
		tName := ""
		for _, t := range data.Teacher {
			if t.ID == c.HomeroomTeacherID {
				tName = t.Name
				break
			}
		}
		fmt.Printf("Class: %s | Name: %s | Homeroom Teacher: %s\n", c.ID, c.Name, tName)
	}
}

func printStudentDESC(data model.Data) {
	sort.Slice(data.Student, func(i, j int) bool {
		return data.Student[i].Name > data.Student[j].Name
	})
	for _, s := range data.Student {
		fmt.Printf("ID: %s | Name: %s | Address: %s\n", s.ID, s.Name, s.Address)
	}
}

func printStudentClasses(data model.Data, studentID string) {
	for _, sc := range data.StudentClass {
		if sc.StudentID == studentID {
			for _, class := range data.Class {
				if class.ID == sc.ClassID {
					fmt.Printf("Class ID: %s | Name: %s\n", class.ID, class.Name)
				}
			}
		}
	}
}

func printTeacherFilter(data model.Data) {
	fmt.Println("Chọn filter:")
	fmt.Println("1. all")
	fmt.Println("2. chủ nhiệm")
	fmt.Println("3. có trên X học sinh")
	fmt.Println("4. chủ nhiệm trên X lớp")
	fmt.Print("Chọn: ")
	var filter int
	fmt.Scanln(&filter)

	switch filter {
	case 1:
		for _, t := range data.Teacher {
			fmt.Printf("Teacher ID: %s | Name: %s\n", t.ID, t.Name)
		}
	case 2:
		for _, t := range data.Teacher {
			for _, c := range data.Class {
				if c.HomeroomTeacherID == t.ID {
					fmt.Printf("Chủ nhiệm: %s\n", t.Name)
					break
				}
			}
		}
	default:
		fmt.Println("Chức năng đang phát triển.")
	}
}

func main() {
	data := utils.LoadData("./seed/data.json")
	fmt.Println(data)
	for {
		fmt.Println("\n=== MENU ===")
		fmt.Println("1. Danh sách lớp sort theo name ASC")
		fmt.Println("2. Danh sách học sinh sort theo name DESC")
		fmt.Println("3. Danh sách lớp học sinh X đã tham gia")
		fmt.Println("4. Danh sách giáo viên với filter")
		fmt.Println("0. Thoát")
		fmt.Print("Chọn: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			printClassASC(data)
		case 2:
			printStudentDESC(data)
		case 3:
			var id string
			fmt.Print("Nhập student_id: ")
			fmt.Scanln(&id)
			printStudentClasses(data, id)
		case 4:
			printTeacherFilter(data)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Lựa chọn không hợp lệ.")
		}
	}
}
