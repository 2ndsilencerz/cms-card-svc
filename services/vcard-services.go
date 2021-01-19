package services

import (
	"fmt"
	"strconv"

	"github.com/2ndsilencerz/cms-card-svc/configs/database"
	"github.com/2ndsilencerz/cms-card-svc/configs/utils"
	"github.com/2ndsilencerz/cms-card-svc/models"
	"github.com/2ndsilencerz/cms-card-svc/models/pb"
	"golang.org/x/net/context"
)

// GetCardList ...
func (s *Server) GetCardList(ctx context.Context, in *pb.Page) (*pb.VCardList, error) {
	utils.LogToFile(fmt.Sprintf("Request: %T", in))
	db := database.InitDB()
	defer database.CloseDB(db)

	vcardList := []models.VCard{}

	limit, err := strconv.Atoi(in.Limit)
	if err != nil {
		return nil, err
	}
	page, err := strconv.Atoi(in.Page)
	if err != nil {
		return nil, err
	}
	offsets := (page-1)*limit - 1
	if offsets == -1 {
		offsets = 0
	}
	filterType := in.FilterType
	filterValue := in.FilterValue
	if filterType == "No Kartu" && filterValue != " " {
		err = db.WithContext(ctx).Where("CRDNO = ?", filterValue).Limit(limit).Offset(offsets).Find(&vcardList).Error
	} else if filterType == "No Rekening" && filterValue != " " {
		err = db.WithContext(ctx).Where("CRACIF = ?", filterValue).Limit(limit).Offset(offsets).Find(&vcardList).Error
	} else {
		err = db.WithContext(ctx).Limit(limit).Offset(offsets).Find(&vcardList).Error
	}
	if err != nil {
		return nil, err
	}

	result := new(pb.VCardList)
	for _, v := range vcardList {
		card := pb.VCard{
			CardNo:     v.CardNo,
			CardType:   v.CardType,
			NameOnCard: v.NameOnCard,
			CifName:    v.CifName,
			Cif:        v.Cif,
			CardBranch: v.CardBranch,
			Status:     v.Status,
			InstantNon: v.InstantNon,
		}
		result.Vcard = append(result.Vcard, &card)
	}

	utils.LogToFile(fmt.Sprintf("Response: %T", result))
	return result, nil
}
