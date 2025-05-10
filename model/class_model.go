package model

import "fmt"

type Class struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	SchoolID          int    `json:"school_id"`
	HomeroomTeacherID int    `json:"homeroom_teacher_id"`
}

type IClass interface {
	PrintClassRelation([]Teacher)
	FindTeacherName([]Teacher) string
}

func (c Class) PrintClassRelation(teachers []Teacher) {
	fmt.Printf("ID: %d | Tên: %s | GVCN: %s\n", c.ID, c.Name, c.FindTeacherName(teachers))
}

func (c Class) FindTeacherName(teachers []Teacher) string {
	for _, teacher := range teachers {
		if teacher.ID == c.HomeroomTeacherID {
			return teacher.Name
		}
	}
	return "N/A"
}
