package services

import (
	"fmt"

	"github.com/2ndsilencerz/cms-card-svc/repository"

	"github.com/2ndsilencerz/cms-card-svc/configs/utils"
	"github.com/2ndsilencerz/cms-card-svc/models"
	"github.com/2ndsilencerz/cms-card-svc/models/pb"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func setCardList(cards []models.VCard) *pb.VCardList {
	result := new(pb.VCardList)
	for _, v := range cards {
		card := pb.VCard{
			CardNo:       v.CardNo,
			CardType:     v.CardType,
			NameOnCard:   v.NameOnCard,
			CifName:      v.CifName,
			Cif:          v.Cif,
			CardBranch:   v.CardBranch,
			Status:       v.Status,
			InstantNon:   v.InstantNon,
			AccFlag:      v.AccFlag,
			BirthDate:    timestamppb.New(v.BirthDate),
			BirthPlace:   v.BirthPlace,
			MotherName:   v.MotherName,
			PhoneNum:     v.PhoneNum,
			Address:      v.Address,
			CreateDate:   timestamppb.New(v.CreateDate),
			LastUpdate:   timestamppb.New(v.LastUpdate),
			ExpireOnCMS:  timestamppb.New(v.ExpireOnCMS),
			ExpireOnCard: timestamppb.New(v.ExpireOnCard),
			ExpirePin:    timestamppb.New(v.ExpirePin),
		}
		result.Vcard = append(result.Vcard, &card)
	}
	return result
}

func setCardListOfCardDetails(vcards []repository.CardDetails) *pb.VCardList {
	result := new(pb.VCardList)
	for _, v := range vcards {
		card := pb.VCard{
			CardNo:       v.VCard.CardNo,
			CardType:     v.VCard.CardType,
			NameOnCard:   v.VCard.NameOnCard,
			CifName:      v.VCard.CifName,
			Cif:          v.VCard.Cif,
			CardBranch:   v.VCard.CardBranch,
			Status:       v.VCard.Status,
			InstantNon:   v.VCard.InstantNon,
			BranchName:   v.Branch.BranchName,
			CardTypeDesc: v.VCardType.Description,
			AccFlag:      v.VCard.AccFlag,
			BirthDate:    timestamppb.New(v.VCard.BirthDate),
			BirthPlace:   v.VCard.BirthPlace,
			MotherName:   v.VCard.MotherName,
			PhoneNum:     v.VCard.PhoneNum,
			Address:      v.VCard.Address,
			CreateDate:   timestamppb.New(v.VCard.CreateDate),
			LastUpdate:   timestamppb.New(v.VCard.LastUpdate),
			ExpireOnCMS:  timestamppb.New(v.VCard.ExpireOnCMS),
			ExpireOnCard: timestamppb.New(v.VCard.ExpireOnCard),
			ExpirePin:    timestamppb.New(v.VCard.ExpirePin),
		}
		result.Vcard = append(result.Vcard, &card)
	}
	return result
}

// GetCardList used in /card/listKartu (Menu Daftar Kartu -> Daftar Kartu)
func (s *Server) GetCardList(ctx context.Context, in *pb.Page) (*pb.VCardList, error) {
	utils.LogToFile(fmt.Sprintf("Request: %T", in))

	repo := &repository.VCardRepository{
		Setting: repository.Setting{
			Ctx:         ctx,
			FilterType:  in.FilterType,
			FilterValue: in.FilterValue,
			Limit:       in.Limit,
			Page:        in.Page,
		}}
	err := repo.GetCardList()
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

	repo := &repository.VCardRepository{
		Setting: repository.Setting{
			Ctx:         ctx,
			FilterType:  in.Page.FilterType,
			FilterValue: in.Page.FilterValue,
			Limit:       in.Page.Limit,
			Page:        in.Page.Page,
		}}

	err := repo.GetVCardToMaintenance(repository.CardActionBlock, in.Branch)
	if err != nil {
		return nil, err
	}

	result := setCardList(repo.VcardList)
	utils.LogToFile(fmt.Sprintf("Response: %T", result))
	return result, nil
}

// GetCardDetails used in /card/detail/{cardNo}
func (s *Server) GetCardDetails(ctx context.Context, in *pb.VCard) (*pb.VCardList, error) {
	utils.LogToFile(fmt.Sprintf("Request: %T", in))

	repo := &repository.VCardRepository{
		Setting: repository.Setting{
			Ctx: ctx,
		}}

	vcards, err := repo.GetDetails(models.VCard{
		CardNo:  in.CardNo,
		AccFlag: in.AccFlag,
	})
	if err != nil {
		return nil, err
	}

	result := setCardListOfCardDetails(vcards)
	utils.LogToFile(fmt.Sprintf("Response: %T", result))
	return result, nil
}
