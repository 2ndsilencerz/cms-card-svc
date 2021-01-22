package models

type Branch struct {
	BranchCode string `gorm:"column:BRCCODE"`
	BranchName string `gorm:"column:BRCNAME"`
}
