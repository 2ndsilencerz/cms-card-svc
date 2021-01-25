package models

// Branch ...
type Branch struct {
	BranchCode string `gorm:"column:BRCCODE"`
	BranchName string `gorm:"column:BRCNAME"`
}
