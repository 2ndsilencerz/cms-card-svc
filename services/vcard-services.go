package services

import (
	"fmt"

	"github.com/2ndsilencerz/cms-card-svc/repository"

	"github.com/2ndsilencerz/cms-card-svc/configs/database"
	"github.com/2ndsilencerz/cms-card-svc/configs/utils"
	"github.com/2ndsilencerz/cms-card-svc/models"
	"github.com/2ndsilencerz/cms-card-svc/models/pb"
	"golang.org/x/net/context"
)

func setOffset(limitStr, pageStr string) (int, error) {
	limit := utils.StrToInt(limitStr)
	page := utils.StrToInt(pageStr)
	if limit == 0 && page == 0 {
		return 0, utils.NewError("Failed to parse limit and page number")
	}
	offsets := (page-1)*limit - 1
	if offsets == -1 {
		offsets = 0
	}
	return offsets, nil
}

func setCardList(cards []models.VCard) *pb.VCardList {
	result := new(pb.VCardList)
	for _, v := range cards {
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
	return result
}

// GetCardList used in /card/listKartu (Menu Daftar Kartu -> Daftar Kartu)
func (s *Server) GetCardList(ctx context.Context, in *pb.Page) (*pb.VCardList, error) {
	utils.LogToFile(fmt.Sprintf("Request: %T", in))

	limit := utils.StrToInt(in.Limit)
	offsets, err := setOffset(in.Limit, in.Page)
	if err != nil {
		return nil, err
	}

	filterType := in.FilterType
	filterValue := in.FilterValue

	repo := &repository.VCardRepository{
		Ctx:         ctx,
		FilterType:  filterType,
		FilterValue: filterValue,
		Limit:       limit,
		Offsets:     offsets,
	}
	err = repo.GetCardList()
	if err != nil {
		return nil, err
	}

	result := setCardList(repo.VcardList)

	utils.LogToFile(fmt.Sprintf("Response: %T", result))
	return result, nil
}

// GetCardBlockedList used in /card/list/block (Menu Daftar Kartu -> Blokir Kartu)
func (s *Server) GetCardBlockedList(ctx context.Context, in *pb.BlockPage) (*pb.VCardList, error) {
	utils.LogToFile(fmt.Sprintf("Request: %T", in))
	db := database.InitDB()
	defer database.CloseDB(db)

	limit := utils.StrToInt(in.Page.Limit)
	offsets, err := setOffset(in.Page.Limit, in.Page.Page)
	if err != nil {
		return nil, err
	}

	filterType := in.Page.FilterType
	filterValue := in.Page.FilterValue

	repo := repository.VCardRepository{
		Ctx:         ctx,
		FilterType:  filterType,
		FilterValue: filterValue,
		Limit:       limit,
		Offsets:     offsets,
	}

	err = repo.GetVCardToMaintenance(repository.CardActionBlock, in.Branch)
	if err != nil {
		return nil, err
	}

	result := setCardList(repo.VcardList)
	utils.LogToFile(fmt.Sprintf("Response: %T", result))
	return result, nil
}
