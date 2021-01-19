package models

// VCard model
type VCard struct {
	CardNo     string `gorm:"column:CRDNO"`
	CardType   string `gorm:"column:CRTYPE"`
	NameOnCard string `gorm:"column:CRDNAM"`
	CifName    string `gorm:"column:CRACFN"`
	Cif        string `gorm:"column:CRACIF"`
	CardBranch string `gorm:"column:CRBRCR"`
	Status     string `gorm:"column:CRSTS"`
	InstantNon string `gorm:"column:CRJNSC"`
}
