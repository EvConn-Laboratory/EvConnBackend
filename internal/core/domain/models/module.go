package models

type Module struct {
	Base
	Name        string       `gorm:"type:varchar(100);not null"`
	Description string       `gorm:"type:text"`
	CourseID    uint         `gorm:"not null"`
	Course      Course       `gorm:"foreignKey:CourseID"`
	Assignments []Assignment `gorm:"foreignKey:ModuleID"`
}
