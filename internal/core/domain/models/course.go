package models

type Course struct {
	Base
	Name        string   `gorm:"type:varchar(100);not null"`
	Description string   `gorm:"type:text"`
	MentorID    uint     `gorm:"not null"`
	LabID       uint     `gorm:"not null"`
	Lab         Lab      `gorm:"foreignKey:LabID"`
	Mentor      User     `gorm:"foreignKey:MentorID"`
	Modules     []Module `gorm:"foreignKey:CourseID"`
}
