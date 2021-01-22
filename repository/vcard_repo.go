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

// CardNoFailedToParseMessage ...
const CardNoFailedToParseMessage = "parameter cannot be parsed or isn't defined"
const CardNoFailedToFindInfo = "failed to find info with this cardNo"

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
		return utils.NewError(CardNoFailedToParseMessage)
	}
	err = db.Find(&p.VcardList).Error

	return err
}

// GetVCardToMaintenance ...
func (p *VCardRepository) GetVCardToMaintenance(action string, branch string) error {

	res := utils.StrToInt(p.FilterValue)
	if res == 0 {
		return utils.NewError(CardNoFailedToParseMessage)
	}

	cardFlag := &constCardFlag{}
	cardFlag.set(action)

	var condition string

	db := database.InitDB()
	defer database.CloseDB(db)

	var err error
	db = db.WithContext(p.Ctx).Limit(p.Limit).Offset(p.Offsets)

	condition = "CRSTS = ?"
	var cardStatus string
	// set filter status
	if cardFlag.action == CardActionActivate {
		cardStatus = utils.IntToStr(cardStatusBlock)
	} else if cardFlag.action == CardActionBlock || cardFlag.action == CardActionChange {
		cardStatus = utils.IntToStr(cardStatusActive)
	} else if cardFlag.action == CardActionClose {
		cardStatus = utils.IntToStr(cardStatusClose)
	}

	// set filter cardNo or accFlag
	if p.FilterType == filterTypeCardNo {
		condition += " AND CRDNO = ?"
	} else if p.FilterType == filterTypeAccFlag {
		condition += " AND ACCFLAG = ?"
	}

	// set if
	if branch != "" && branch[:1] == syariahBranchPre {
		condition += " AND CRBRCR LIKE ?"
	} else {
		condition += " AND CRBRCR NOT LIKE ?"
	}
	db = db.Where(condition, cardStatus, p.FilterValue, branch[:1]+"%")
	err = db.Find(&p.VcardList).Error
	return err
}

type CardDetails struct {
	models.VCard
	models.VCardType
	models.Branch
}

// GetDetails of card inputted
func (v *VCardRepository) GetDetails(card models.VCard) ([]CardDetails, error) {

	db := database.InitDB()
	defer database.CloseDB(db)
	var err error

	if utils.StrToInt(card.AccFlag) == 0 {
		return nil, utils.NewError(CardNoFailedToParseMessage)
	}

	res := []CardDetails{}

	condition := "v_cards.ACCFLAG IS NOT NULL"
	condition += " AND v_cards.ACCFLAG != '-'"
	condition += " AND v_cards.ACCFLAG = ?"
	db = db.WithContext(v.Ctx).Model(&models.VCard{}).Select("*")
	db = db.Joins("left join branches on v_cards.CRBRCR = branches.BRCCODE")
	db = db.Joins("left join v_card_types on v_cards.CRTYPE = v_card_types.CTTYPE")
	db = db.Where(condition, card.AccFlag)
	err = db.Find(&res).Error
	return res, err
}
