package model

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
