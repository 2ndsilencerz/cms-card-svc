package models

import "time"

// VCard model
type VCard struct {
	CardNo   string `gorm:"column:CRDNO"`
	CardType string `gorm:"column:CRTYPE"`
	// CardTypeData   *VCardType `gorm:"foreignKey:CRTYPE"`
	NameOnCard string `gorm:"column:CRDNAM"`
	CifName    string `gorm:"column:CRACFN"`
	Cif        string `gorm:"column:CRACIF"`
	CardBranch string `gorm:"column:CRBRCR"`
	// MainOcBranch string `gorm:"column:MOCBRC"`
	// CardBranchCode *Branch    `gorm:"foreignKey:CRBRCR"`
	Status       string    `gorm:"column:CRSTS"`
	InstantNon   string    `gorm:"column:CRJNSC"`
	AccFlag      string    `gorm:"column:ACCFLAG"`
	BirthDate    time.Time `gorm:"column:TGLLAHIR"`
	BirthPlace   string    `gorm:"column:TMPLAHIR"`
	MotherName   string    `gorm:"column:NAMAIBUK"`
	PhoneNum     string    `gorm:"column:PHONENUM"`
	Address      string    `gorm:"column:ADDRESS"`
	CreateDate   time.Time `gorm:"column:CREATEDATE"`
	LastUpdate   time.Time `gorm:"column:LASTUPDATE"`
	ExpireOnCMS  time.Time `gorm:"column:CREXPR"`
	ExpireOnCard time.Time `gorm:"column:CREXCR"`
	ExpirePin    time.Time `gorm:"column:EXPIREPIN"`
}
