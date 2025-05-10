package model

type Data struct {
	Schools        []School       `json:"schools"`
	Teachers       []Teacher      `json:"teachers"`
	Classes        []Class        `json:"classes"`
	Students       []Student      `json:"students"`
	StudentClasses []StudentClass `json:"student_classes"`
}

type IData interface {
	getClassesByStudentID(studentID int)
	FilterTeachersByHomeroom() []Teacher
	FilterTeacherByMinStudentCount(min int) []Teacher
	FilterTeacherByMinHomeroomClasses(min int) []Teacher
}

func (d Data) GetClassesOfStudent(studentID int) []Class {
	var classes []Class
	for _, sc := range d.StudentClasses {
		if sc.StudentID == studentID {
			for _, class := range d.Classes {
				if class.ID == sc.ClassID {
					classes = append(classes, class)
				}
			}
		}
	}
	return classes
}

func (d Data) GetAllTeacher() []Teacher {
	return d.Teachers
}

func (d Data) FilterTeachersByHomeroom() []Teacher {
	teacherMap := make(map[int]bool)
	for _, class := range d.Classes {
		teacherMap[class.HomeroomTeacherID] = true
	}

	var result []Teacher
	for _, teacher := range d.Teachers {
		if teacherMap[teacher.ID] {
			result = append(result, teacher)
		}
	}
	return result
}

func (d Data) FilterTeacherByMinStudentCount(min int) []Teacher {
	// Map class_id -> []student_id
	classStudentCount := make(map[int]int)
	for _, sc := range d.StudentClasses {
		classStudentCount[sc.ClassID]++
	}

	// Map teacher_id -> tổng số học sinh dạy
	teacherStudentCount := make(map[int]int)
	for _, class := range d.Classes {
		studentCount := classStudentCount[class.ID]
		teacherStudentCount[class.HomeroomTeacherID] += studentCount
	}

	var result []Teacher
	for _, teacher := range d.Teachers {
		if teacherStudentCount[teacher.ID] > min {
			result = append(result, teacher)
		}
	}
	return result
}

func (d Data) FilterTeacherByMinHomeroomClasses(min int) []Teacher {
	teacherClassCount := make(map[int]int)
	for _, class := range d.Classes {
		teacherClassCount[class.HomeroomTeacherID]++
	}

	var result []Teacher
	for _, teacher := range d.Teachers {
		if teacherClassCount[teacher.ID] > min {
			result = append(result, teacher)
		}
	}
	return result
}
