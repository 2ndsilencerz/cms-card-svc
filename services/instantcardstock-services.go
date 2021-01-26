package services

import (
	"fmt"

	"github.com/2ndsilencerz/cms-card-svc/models"

	"github.com/2ndsilencerz/cms-card-svc/configs/utils"
	"github.com/2ndsilencerz/cms-card-svc/models/pb"
	"github.com/2ndsilencerz/cms-card-svc/repository"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func setVCardTypeList(list []models.VCardType) *pb.VCardTypeList {
	result := new(pb.VCardTypeList)
	for _, v := range list {
		vtype := pb.VCardType{
			TypeCode:    v.TypeCode,
			Description: v.Description,
			MainType:    v.MainType,
		}
		result.VCardType = append(result.VCardType, &vtype)
	}
	return result
}

func setInstantCardStockList(list []models.InstantStock) *pb.InstantStockList {
	result := new(pb.InstantStockList)
	for _, v := range list {
		instantStock := pb.InstantStock{
			Id:           v.ID,
			EmbossDate:   timestamppb.New(v.EmbossDate),
			EndNumber:    v.EndNumber,
			GenerateDate: timestamppb.New(v.GenerateDate),
			Print:        int32(v.Print),
			ServiceCode:  int32(v.ServiceCode),
			StartNumber:  v.StartNumber,
			Status:       v.Status,
			Total:        int32(v.Total),
			Type:         v.Type,
			UserEmboss:   v.UserEmboss,
			UserGenerate: v.UserGenerate,
		}
		result.InstantStock = append(result.InstantStock, &instantStock)
	}
	return result
}

// GetInstantCardType used in /card/stock (Stock Kartu Instan -> Generate Stock)
func (s *Server) GetInstantCardType(ctx context.Context, empty *emptypb.Empty) (*pb.VCardTypeList, error) {
	utils.LogToFile(fmt.Sprint("Request: GetInstantCardType"))

	repo := &repository.VCardTypeRepository{
		Ctx: ctx,
	}
	err := repo.GetInstantCardType()
	if err != nil {
		return nil, err
	}
	result := setVCardTypeList(repo.CardTypeList)
	utils.LogToFile(fmt.Sprintf("Response: %T", result))
	return result, nil
}

// GetDataForEmboss used in /card/stock/emboss (Stock Kartu Instant -> Cetak Stock Kartu)
func (s *Server) GetDataForEmboss(ctx context.Context, in *pb.EmbossRequest) (*pb.InstantStockList, error) {
	utils.LogToFile(fmt.Sprintf("Request: %v", in))

	repo := &repository.InstantStockRepository{
		Setting: repository.Setting{
			Ctx:   ctx,
			Limit: in.Page.Limit,
			Page:  in.Page.Page,
		},
		StartDate: in.StartDate.AsTime(),
		EndDate:   in.EndDate.AsTime(),
		Type:      in.Type,
	}
	err := repo.GetDataForEmboss()
	if err != nil {
		return nil, err
	}
	result := setInstantCardStockList(repo.InstantStockList)
	utils.LogToFile(fmt.Sprintf("Response: %T", result))
	return result, nil
}
