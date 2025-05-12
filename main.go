package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"practice_01/model"
	printService "practice_01/service"
	"practice_01/utils"
)

func LoadData(path string) model.Data {
	bytes, err := os.ReadFile(path)

	if err != nil {
		log.Fatalf("Error reading data file: %v", err)
	}

	var data model.Data
	if err := json.Unmarshal(bytes, &data); err != nil {
		log.Fatalf("Error Unmarshal data: %v", err)
	}

	return data
}

func PrintTeacherFilter(d model.Data) {
	for {

		filter := printService.PrintTeacherFilterMenu()

		switch filter {
		case 1:
			printService.PrintTeachers("Tất cả giáo viên", d.Teachers)
		case 2:
			printService.PrintTeachers("Giáo viên là chủ nhiệm:", d.FilterTeachersByHomeroom())
		case 3:
			x := utils.PromptInt("Nhập số học sinh tối thiểu (X): ")
			title := fmt.Sprintf("Giáo viên chủ nhiệm có trên %d học sinh:", x)
			printService.PrintTeachers(title, d.FilterTeacherByMinStudentCount(x))
		case 4:
			x := utils.PromptInt("Nhập số lớp tối thiểu (X): ")
			title := fmt.Sprintf("Giáo viên chủ nhiệm trên %d lớp:", x)
			printService.PrintTeachers(title, d.FilterTeacherByMinHomeroomClasses(x))
		case 5:
			fmt.Println("✅ Thoát filter.")
			return
		default:
			fmt.Println("❌ Lựa chọn không hợp lệ. Vui lòng chọn từ 1 đến 5")
		}
	}
}

func main() {
	d := LoadData("./seed/data.json")
	for {
		choice := printService.PrintMenu()

		switch choice {
		case 1:
			classes := utils.SortData(d.Classes, utils.Asc, "Name")
			printService.PrintClassesRelation(classes, d.Teachers)
		case 2:
			students := utils.SortData(d.Students, utils.Desc, "Name")
			printService.PrintStudents(students)
		case 3:
			id := utils.PromptInt("Nhập student_id: ")
			classes := d.GetClassesOfStudent(id)
			printService.PrintClassesOfStudent(id, classes)
		case 4:
			PrintTeacherFilter(d)
		case 0:
			os.Exit(0)
		default:
			fmt.Println("Lựa chọn không hợp lệ.")
		}
	}
}
