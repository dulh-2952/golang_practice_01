package model

import (
	"fmt"
	"practice_01/utils"
)

type Data struct {
	Schools        []School       `json:"schools"`
	Teachers       []Teacher      `json:"teachers"`
	Classes        []Class        `json:"classes"`
	Students       []Student      `json:"students"`
	StudentClasses []StudentClass `json:"student_classes"`
}

type IData interface {
	PrintClass(order utils.SortOrder, field string)
	PrintStudents(order utils.SortOrder, field string)
	PrintClassesOfStudent(studentID int)
}

func (d Data) PrintClass(order utils.SortOrder, field string) {
	classes := utils.SortData(d.Classes, order, field)
	fmt.Println("======= DANH SÁCH LỚP HỌC =======")
	for _, c := range classes {
		c.PrintClassRelation(d.Teachers)
	}
}

func (d Data) PrintStudent(order utils.SortOrder, field string) {
	students := utils.SortData(d.Students, order, field)
	fmt.Println("======= DANH SÁCH HỌC SINH =======")
	for _, s := range students {
		s.PrintStudent()
	}
}

func (d Data) PrintClassesOfStudent(studentID int) {
	fmt.Printf("======= Danh Sách Lớp Học Của Học Sinh - ID = %d =======\n", studentID)
	for _, sc := range d.StudentClasses {
		if sc.StudentID == studentID {
			for _, class := range d.Classes {
				if class.ID == sc.ClassID {
					fmt.Printf("ID: %d | Tên: %s\n", class.ID, class.Name)
				}
			}
		}
	}
}
