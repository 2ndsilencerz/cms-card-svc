package repository

import (
	"fmt"
	"time"

	"github.com/2ndsilencerz/cms-card-svc/models"

	"github.com/2ndsilencerz/cms-card-svc/configs/database"
	"github.com/2ndsilencerz/cms-card-svc/configs/utils"

	"golang.org/x/net/context"
)

// InstantStockRepository ...
type InstantStockRepository struct {
	Ctx              context.Context
	Type             string
	StartDate        time.Time
	EndDate          time.Time
	InstantStockList []models.InstantStock
}

// GetDataForEmboss ...
func (p *InstantStockRepository) GetDataForEmboss() error {

	db := database.InitDB()
	defer database.CloseDB(db)
	var err error

	var condition string = "STATUS = ?"
	db = db.WithContext(p.Ctx)
	if len(p.Type) > 0 && utils.StrToInt(p.Type) != 0 {
		condition += " AND TYPE = ?"
	} else if len(p.Type) > 0 && utils.StrToInt(p.Type) == 0 {
		return utils.NewError(CardNoFailedToParseMessage)
	}

	// somehow the IsZero() return false if we don't set the value
	// which is technically wrong
	// (maybe because the code tested in windows the result was 1970-01-01 00:00:00 +0000 UTC).
	// we will use Nanosecond() instead to compare
	if p.StartDate.Nanosecond() != 0 {
		condition += " AND GENERATE_DATE >= ?"
	}
	if p.EndDate.Nanosecond() != 0 {
		condition += " AND GENERATE_DATE <= ?"
	}

	db = db.Where(condition, "G", p.Type, p.StartDate, p.EndDate)
	err = db.Find(&p.InstantStockList).Error
	utils.PrintToConsole(fmt.Sprint(p.InstantStockList))
	return err
}
