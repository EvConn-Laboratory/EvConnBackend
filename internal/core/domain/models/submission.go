package models

type Submission struct {
	Base
	AssignmentID uint       `gorm:"not null"`
	StudentID    uint       `gorm:"column:user_id;not null"` // Changed column name to match query
	Content      string     `gorm:"type:text"`
	Score        *float64   `gorm:"type:decimal(5,2)"`
	Status       string     `gorm:"type:enum('pending','graded');default:'pending'"`
	SubmittedAt  string     `gorm:"type:datetime;not null"`
	Assignment   Assignment `gorm:"foreignKey:AssignmentID"`
	Student      User       `gorm:"foreignKey:StudentID;references:ID"`
}
