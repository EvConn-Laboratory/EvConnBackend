package models

type Assignment struct {
	Base
	Title       string       `gorm:"type:varchar(100);not null"`
	Description string       `gorm:"type:text"`
	ModuleID    uint         `gorm:"not null"`
	Module      Module       `gorm:"foreignKey:ModuleID"`
	MaxScore    float64      `gorm:"type:decimal(5,2);default:100"`
	Deadline    string       `gorm:"type:varchar(20)"` // Format: "2006-01-02 15:04:05"
	Status      string       `gorm:"type:enum('active','archived');default:'active'"`
	Submissions []Submission `gorm:"foreignKey:AssignmentID"`
}
