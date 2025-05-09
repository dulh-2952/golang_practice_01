package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"practice_01/model"
	"time"
)

func randomAlphaAddress() string {
	return string('A' + rune(rand.Intn(26)))
}

func studentName(i int) string {
	return fmt.Sprintf("student_%03d", i)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	school := model.School{ID: 1, Name: "Golang School"}

	teachers := []model.Teacher{
		{ID: 1, Name: "Teacher A", Address: randomAlphaAddress()},
		{ID: 2, Name: "Teacher B", Address: randomAlphaAddress()},
	}

	classes := []model.Class{
		{ID: 1, Name: "Golang Basic", SchoolID: 1, HomeroomTeacherID: 1},
		{ID: 2, Name: "Golang Advance", SchoolID: 1, HomeroomTeacherID: 2},
		{ID: 3, Name: "Golang Master", SchoolID: 1, HomeroomTeacherID: 1},
	}

	// Tạo danh sách học sinh (ví dụ 100 học sinh)
	var students []model.Student
	for i := 1; i <= 100; i++ {
		students = append(students, model.Student{
			ID:      i,
			Name:    studentName(i),
			Address: randomAlphaAddress(),
		})
	}

	// Gán học sinh vào nhiều lớp ngẫu nhiên
	var studentClasses []model.StudentClass
	for _, student := range students {
		numClasses := rand.Intn(3) + 1 // Mỗi học sinh học từ 1 đến 3 lớp
		classIDs := rand.Perm(3)[:numClasses]
		for _, idx := range classIDs {
			studentClasses = append(studentClasses, model.StudentClass{
				StudentID: student.ID,
				ClassID:   classes[idx].ID,
			})
		}
	}

	data := model.Data{
		Schools:        []model.School{school},
		Teachers:       teachers,
		Classes:        classes,
		Students:       students,
		StudentClasses: studentClasses,
	}

	file, err := os.Create("data.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(data); err != nil {
		panic(err)
	}

	fmt.Println("✅ File data.json đã được tạo")
}
