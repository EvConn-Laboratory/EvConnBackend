package models

// filepath: /backend/internal/core/domain/models/lab.go
type Lab struct {
	Base
	Name    string   `gorm:"type:varchar(100);not null"`
	Code    string   `gorm:"type:varchar(50);uniqueIndex;not null"`
	Courses []Course `gorm:"foreignKey:LabID"`
}
