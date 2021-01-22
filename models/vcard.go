package models

// VCard model
type VCard struct {
	CardNo   string `gorm:"column:CRDNO"`
	CardType string `gorm:"column:CRTYPE"`
	// CardTypeData   *VCardType `gorm:"foreignKey:CRTYPE"`
	NameOnCard string `gorm:"column:CRDNAM"`
	CifName    string `gorm:"column:CRACFN"`
	Cif        string `gorm:"column:CRACIF"`
	CardBranch string `gorm:"column:CRBRCR"`
	// CardBranchCode *Branch    `gorm:"foreignKey:CRBRCR"`
	Status     string `gorm:"column:CRSTS"`
	InstantNon string `gorm:"column:CRJNSC"`
	AccFlag    string `gorm:"column:ACCFLAG"`
}
