package models

type User struct {
	Base
	NIM         string   `gorm:"type:varchar(20);uniqueIndex;not null"`
	Name        string   `gorm:"type:varchar(100);not null"`
	Email       string   `gorm:"type:varchar(100);uniqueIndex"`
	Password    string   `gorm:"type:varchar(255);not null"`
	Role        string   `gorm:"type:enum('admin','mentor','student');not null;default:'student'"`
	Lab         string   `gorm:"type:varchar(50);not null"`
	Shift       *string  `gorm:"type:varchar(20)"`
	Hari        *string  `gorm:"type:varchar(20)"`
	KodeAsisten *string  `gorm:"type:varchar(20)"`
	Courses     []Course `gorm:"foreignKey:MentorID"`
}

// GetRole returns the user's role
func (u *User) GetRole() string {
	return u.Role
}
