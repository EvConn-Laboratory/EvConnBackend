package models

import "time"

const (
	EnrollmentStatusPending  = "pending"
	EnrollmentStatusApproved = "approved"
	EnrollmentStatusRejected = "rejected"
)

type Enrollment struct {
	Base
	StudentID  uint       `json:"student_id" gorm:"not null"`
	CourseID   uint       `json:"course_id" gorm:"not null"`
	Status     string     `json:"status" gorm:"default:'pending'"`
	Student    User       `json:"student" gorm:"foreignKey:StudentID"`
	Course     Course     `json:"course" gorm:"foreignKey:CourseID"`
	ApprovedAt *time.Time `json:"approved_at"`
	Notes      string     `json:"notes"`
}
