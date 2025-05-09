package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type School struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Teacher struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Class struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	SchoolID          int    `json:"school_id"`
	HomeroomTeacherID int    `json:"homeroom_teacher_id"`
}

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type StudentClass struct {
	StudentID int `json:"student_id"`
	ClassID   int `json:"class_id"`
}

type Data struct {
	Schools        []School       `json:"schools"`
	Teachers       []Teacher      `json:"teachers"`
	Classes        []Class        `json:"classes"`
	Students       []Student      `json:"students"`
	StudentClasses []StudentClass `json:"student_classes"`
}

func randomAlphaAddress() string {
	return string('A' + rune(rand.Intn(26)))
}

func studentName(i int) string {
	return fmt.Sprintf("student_%03d", i)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	school := School{ID: 1, Name: "Golang School"}

	teachers := []Teacher{
		{ID: 1, Name: "Teacher A", Address: randomAlphaAddress()},
		{ID: 2, Name: "Teacher B", Address: randomAlphaAddress()},
	}

	classes := []Class{
		{ID: 1, Name: "Golang Basic", SchoolID: 1, HomeroomTeacherID: 1},
		{ID: 2, Name: "Golang Advance", SchoolID: 1, HomeroomTeacherID: 2},
		{ID: 3, Name: "Golang Master", SchoolID: 1, HomeroomTeacherID: 1},
	}

	var students []Student
	var studentClasses []StudentClass

	studentID := 1
	for classID := 1; classID <= 3; classID++ {
		for i := 0; i < 50; i++ {
			students = append(students, Student{
				ID:      studentID,
				Name:    studentName(studentID),
				Address: randomAlphaAddress(),
			})
			studentClasses = append(studentClasses, StudentClass{
				StudentID: studentID,
				ClassID:   classID,
			})
			studentID++
		}
	}

	data := Data{
		Schools:        []School{school},
		Teachers:       teachers,
		Classes:        classes,
		Students:       students,
		StudentClasses: studentClasses,
	}

	file, err := os.Create("./data.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		panic(err)
	}

	fmt.Println("✅ File data.json đã được tạo!")
}
