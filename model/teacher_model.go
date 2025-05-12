package model

import "fmt"

type Teacher struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type ITeacher interface {
	ShowInfo()
}

func (t Teacher) ShowInfo() {
	fmt.Printf("ID: %d | Tên: %s | Địa chỉ: %s\n", t.ID, t.Name, t.Address)
}
