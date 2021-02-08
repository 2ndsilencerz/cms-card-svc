package models

// VPCard ...
type VPCard struct {
	CardType    string `gorm:"column:VFTYPE"`
	Status      int    `gorm:"column:VPSTSU"`
	ServiceCode string `gorm:"column:VPSVCD"`
}
