package repository

import (
	"github.com/2ndsilencerz/cms-card-svc/configs/database"
	"github.com/2ndsilencerz/cms-card-svc/models"
)

// VCardTypeRepository ...
type VCardTypeRepository struct {
	Setting
	CardTypeList []models.VCardType
}

// GetInstantCardType ...
func (p *VCardTypeRepository) GetInstantCardType() error {

	db := database.InitDB()
	defer database.CloseDB(db)
	var err error

	db = db.WithContext(p.Ctx).Where("ISREGULER = ? OR ISREGULER = ?", "*", "0")
	err = db.Find(&p.CardTypeList).Error
	return err
}
