package model

import "fmt"

type Student struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type IStudent interface {
	ShowInfo()
}

func (s Student) ShowInfo() {
	fmt.Printf("ID: %d | Tên: %s | Địa chỉ: %s\n", s.ID, s.Name, s.Address)
}
