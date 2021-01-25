package models

import "time"

// InstantStock ...
type InstantStock struct {
	ID           string    `gorm:"column:ID"`
	EmbossDate   time.Time `gorm:"column:EMBOSS_DATE"`
	EndNumber    string    `gorm:"column:END_NUMBER"`
	GenerateDate time.Time `gorm:"column:GENERATE_DATE"`
	Print        int       `gorm:"column:PRINT"`
	ServiceCode  int       `gorm:"column:SERVICE_CODE"`
	StartNumber  string    `gorm:"column:START_NUMBER"`
	Status       string    `gorm:"column:STATUS"`
	Total        int       `gorm:"column:TOTAL"`
	Type         string    `gorm:"column:TYPE"`
	UserEmboss   string    `gorm:"column:USER_EMBOSS"`
	UserGenerate string    `gorm:"column:USER_GENERATE"`
}
