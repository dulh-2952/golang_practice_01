package main

import (
	"fmt"
	"os"
	"sort"

	"practice_01/model"

	"practice_01/utils"
)

func printClassASC(data model.Data) {
	sort.Slice(data.Classes, func(i, j int) bool {
		return data.Classes[i].Name < data.Classes[j].Name
	})
	for _, c := range data.Classes {
		tName := ""
		for _, t := range data.Teachers {
			if t.ID == c.HomeroomTeacherID {
				tName = t.Name
				break
			}
		}
		fmt.Printf("ID: %d | Tên: %s | GVCN: %s\n", c.ID, c.Name, tName)
	}
}

func printStudentDESC(data model.Data) {
	sort.Slice(data.Students, func(i, j int) bool {
		return data.Students[i].Name > data.Students[j].Name
	})
	for _, s := range data.Students {
		fmt.Printf("ID: %d | Tên: %s | Địa chỉ: %s\n", s.ID, s.Name, s.Address)
	}
}

func printStudentClasses(data model.Data, studentID int) {
	for _, sc := range data.StudentClasses {
		if sc.StudentID == studentID {
			for _, class := range data.Classes {
				if class.ID == sc.ClassID {
					fmt.Printf("ID: %d | Tên: %s\n", class.ID, class.Name)
				}
			}
		}
	}
}

func printTeacherFilter(data model.Data) {
	for {

		filter := utils.PrintTeacherFilterMenu()

		switch filter {
		case 1:
			for _, t := range data.Teachers {
				fmt.Printf("ID: %d | Tên: %s | Địa chỉ: %s\n", t.ID, t.Name, t.Address)
			}

		case 2:
			fmt.Println("Giáo viên là chủ nhiệm:")
			for _, t := range data.Teachers {
				for _, c := range data.Classes {
					if c.HomeroomTeacherID == t.ID {
						fmt.Printf("- %s\n", t.Name)
						break
					}
				}
			}

		case 3:
			x := utils.PromptInt("Nhập số học sinh tối thiểu (X): ")

			teacherStudentCount := make(map[int]int)

			for _, class := range data.Classes {
				for _, sc := range data.StudentClasses {
					if sc.ClassID == class.ID {
						teacherStudentCount[class.HomeroomTeacherID]++
					}
				}
			}

			fmt.Printf("Giáo viên chủ nhiệm có trên %d học sinh:\n", x)
			for _, t := range data.Teachers {
				if count := teacherStudentCount[t.ID]; count > x {
					fmt.Printf("- %s (%d học sinh)\n", t.Name, count)
				}
			}

		case 4:
			x := utils.PromptInt("Nhập số lớp tối thiểu (X): ")

			teacherClassCount := make(map[int]int)
			for _, class := range data.Classes {
				teacherClassCount[class.HomeroomTeacherID]++
			}

			fmt.Printf("Giáo viên chủ nhiệm trên %d lớp:\n", x)
			for _, t := range data.Teachers {
				if count := teacherClassCount[t.ID]; count > x {
					fmt.Printf("- %s (%d lớp)\n", t.Name, count)
				}
			}

		case 5:
			fmt.Println("✅ Thoát filter.")
			return

		default:
			fmt.Println("❌ Lựa chọn không hợp lệ. Vui lòng chọn từ 1 đến 5.")
		}
	}
}

func main() {
	data := utils.LoadData("./seed/data.json")
	for {
		choice := utils.PrintMenu()

		switch choice {
		case 1:
			printClassASC(data)
		case 2:
			printStudentDESC(data)
		case 3:
			id := utils.PromptInt("Nhập student_id: ")
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
