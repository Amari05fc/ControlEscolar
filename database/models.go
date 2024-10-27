package database

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Id    int `gorm:"primarykey"`
	Name  string
	Group string
	Email string
}

type Subject struct {
	gorm.Model
	Id   int `gorm:"primaryKey"`
	Name string
}

type Grade struct {
	gorm.Model
	Id        int `gorm:"primaryKey"`
	StudentID int
	Student   Student `gorm: "constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SubjectID int
	Subject   Subject `gorm: "constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Grade     float64
}
