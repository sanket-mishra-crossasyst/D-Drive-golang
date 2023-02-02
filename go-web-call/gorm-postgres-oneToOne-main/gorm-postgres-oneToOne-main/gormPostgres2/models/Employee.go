package models

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	EmployeeName string     `json:"employeeName"`
	Salary       string     `json:"salary"`
	Department   Department `json:"department"`
	DepartmentId uint       `gorm:"foreignKey:ID" json:"_"`
}
