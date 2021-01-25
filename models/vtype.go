package models

// VCardType ...
type VCardType struct {
	TypeCode    string `gorm:"column:CTTYPE"`
	Description string `gorm:"column:CTDESC"`
	MainType    string `gorm:"column:MAINTYPE"`
}
