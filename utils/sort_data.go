package utils

import (
	"reflect"
	"sort"
)

type SortOrder int

const (
	Asc  SortOrder = iota // 0
	Desc                  // 1
)

// Hàm sắp xếp chung cho tất cả các kiểu dữ liệu có thể sắp xếp
func SortData[T any](data []T, order SortOrder, field string) []T {
	// Tạo bản sao để tránh thay đổi dữ liệu gốc
	sortedData := append([]T(nil), data...)

	// Sử dụng reflection để lấy giá trị của trường
	sort.Slice(sortedData, func(i, j int) bool {
		valI := reflect.ValueOf(sortedData[i])
		valJ := reflect.ValueOf(sortedData[j])

		// Lấy giá trị của trường cần sắp xếp
		fieldI := valI.FieldByName(field).String()
		fieldJ := valJ.FieldByName(field).String()

		if order == Asc {
			return fieldI < fieldJ
		}
		return fieldI > fieldJ // Sắp xếp giảm dần
	})

	return sortedData
}
