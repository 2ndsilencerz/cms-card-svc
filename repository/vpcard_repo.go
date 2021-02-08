package repository

import (
	"github.com/2ndsilencerz/cms-card-svc/configs/database"
	"github.com/2ndsilencerz/cms-card-svc/models"
)

// VPCardRepository ...
type VPCardRepository struct {
	Setting
	VPCardList []models.VPCard
}

func (p *VPCardRepository) CountByCardTypeAndStatusAndServiceCode(cardType string, serviceCode string) (int, error) {
	embossed := 4
	db := database.InitDB()
	defer database.CloseDB(db)
	var err error

	var count int64
	db = db.WithContext(p.Ctx).Model(models.VPCard{}).Where("VFTYPE = ? AND VPSTSU = ? AND VPSVCD = ?",
		cardType, embossed, serviceCode)
	err = db.Count(&count).Error
	return int(count), err
}
