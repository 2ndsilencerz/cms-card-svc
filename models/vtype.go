package models

type VCardType struct {
	TypeCode    string `gorm:"column:CTTYPE"`
	Description string `gorm:"column:CTDESC"`
}
