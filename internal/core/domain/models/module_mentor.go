package models

type ModuleMentor struct {
	Base
	ModuleID uint   `json:"module_id"`
	MentorID uint   `json:"mentor_id"`
	Module   Module `gorm:"foreignKey:ModuleID"`
	Mentor   User   `gorm:"foreignKey:MentorID"`
}
