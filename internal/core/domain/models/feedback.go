package models

type Feedback struct {
	Base
	Type     string `gorm:"type:enum('assistant','laboratorium','session');not null"`
	Rating   int    `gorm:"type:tinyint;not null"`
	Comment  string `gorm:"type:text"`
	ModuleID uint   `gorm:"not null"`
	UserID   uint   `gorm:"not null"`
	Module   Module `gorm:"foreignKey:ModuleID"`
	User     User   `gorm:"foreignKey:UserID"`
}
