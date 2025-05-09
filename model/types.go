package model

type School struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Class struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	SchoolID          string `json:"school_id"`
	HomeroomTeacherID string `json:"homeroom_teacher_id"`
}

type Teacher struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type Student struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type StudentClass struct {
	StudentID string `json:"student_id"`
	ClassID   string `json:"class_id"`
}

type Data struct {
	School       []School       `json:"schools"`
	Class        []Class        `json:"classes"`
	Teacher      []Teacher      `json:"teachers"`
	Student      []Student      `json:"students"`
	StudentClass []StudentClass `json:"student_classes"`
}
