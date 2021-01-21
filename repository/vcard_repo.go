package repository

import (
	"github.com/2ndsilencerz/cms-card-svc/configs/database"
	"github.com/2ndsilencerz/cms-card-svc/configs/utils"
	"github.com/2ndsilencerz/cms-card-svc/models"
	"golang.org/x/net/context"
)

const cardFlagAll = 1
const cardFlagKonsya = 2
const cardFlagBranch = 3

// CardActionActivate ...
const CardActionActivate = "activate"

// CardActionBlock ...
const CardActionBlock = "block"

// CardActionClose ...
const CardActionClose = "close"

// CardActionChange ...
const CardActionChange = "change"

const cardStatusBlock = 4
const cardStatusActive = 1
const cardStatusClose = 3
const filterTypeCardNo = "cardNo"
const filterTypeAccFlag = "accFlag"
const syariahBranchPre = "6"

type constCardFlag struct {
	flag   int
	action string
}

func (c *constCardFlag) set(action string) {
	switch action {
	case CardActionActivate, CardActionBlock, CardActionChange:
		c.flag = cardFlagAll
	case CardActionClose:
		c.flag = cardFlagBranch
	}
	c.action = action
}

// VCardRepository used to
type VCardRepository struct {
	Ctx         context.Context
	FilterType  string
	FilterValue string
	Limit       int
	Offsets     int
	VcardList   []models.VCard
}

// Note: Query using Gorm should place Find() at the last place
// or else our condition and parameter will be ignored

// GetCardList ...
func (p *VCardRepository) GetCardList() error {

	db := database.InitDB()
	defer database.CloseDB(db)
	var err error

	res := utils.StrToInt(p.FilterValue)
	db = db.WithContext(p.Ctx).Limit(p.Limit).Offset(p.Offsets)
	if p.FilterType == filterTypeCardNo && res != 0 {
		db = db.Where("CRDNO = ?", p.FilterValue)
	} else if p.FilterType == filterTypeAccFlag && res != 0 {
		db = db.Where("CRACIF = ?", p.FilterValue)
	} else {
		return utils.NewError("parameter cannot be parsed or isn't defined")
	}
	err = db.Find(&p.VcardList).Error

	return err
}

// GetVCardToMaintenance ...
func (p *VCardRepository) GetVCardToMaintenance(action string, branch string) error {

	cardFlag := &constCardFlag{}
	cardFlag.set(action)

	var condition string

	db := database.InitDB()
	defer database.CloseDB(db)

	var err error
	db = db.WithContext(p.Ctx).Limit(p.Limit).Offset(p.Offsets)

	// set filter status
	if cardFlag.action == CardActionActivate {
		condition = "CRSTS = '" + utils.IntToStr(cardStatusBlock) + "'"
	} else if cardFlag.action == CardActionBlock || cardFlag.action == CardActionChange {
		condition = "CRSTS = '" + utils.IntToStr(cardStatusActive) + "'"
	} else if cardFlag.action == CardActionClose {
		condition = "CRSTS = '" + utils.IntToStr(cardStatusClose) + "'"
	}

	res := utils.StrToInt(p.FilterValue)
	// set filter cardNo or accFlag
	if p.FilterType == filterTypeCardNo && res != 0 {
		condition += " AND CRDNO = '" + p.FilterValue + "'"
	} else if p.FilterType == filterTypeAccFlag && res != 0 {
		condition += " AND ACCFLAG = '" + p.FilterValue + "'"
	} else {
		return utils.NewError("parameter cannot be parsed or isn't defined")
	}

	// set if
	if branch != "" && branch[:1] == syariahBranchPre {
		condition += " AND CRBRCR LIKE ?"
	} else {
		condition += " AND CRBRCR NOT LIKE ?"
	}
	db = db.Where(condition, branch[:1]+"%")
	err = db.Find(&p.VcardList).Error
	return err
}
