package test

import (
	"fmt"
	"testing"

	"github.com/2ndsilencerz/cms-card-svc/models/pb"

	"github.com/2ndsilencerz/cms-card-svc/configs/utils"
	"github.com/2ndsilencerz/cms-card-svc/services"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestGetInstantCardType(t *testing.T) {

	// conn, err := grpc.Dial(":9991", grpc.WithInsecure())
	// if err != nil {
	// 	t.Error(err)
	// 	return
	// }
	// client := pb.NewCardListClient(conn)
	server := &services.Server{}
	empty := new(emptypb.Empty)
	cardTypeList, err := server.GetInstantCardType(ctx, empty)
	if err != nil {
		t.Error(err)
		return
	}

	if len(cardTypeList.VCardType) <= 0 {
		t.Error("no data found")
		return
	}

	utils.LogToFile(fmt.Sprintf("TestGetInstantCardType result: %v", cardTypeList.VCardType))
}

func TestGetDataForEmboss(t *testing.T) {

	server := &services.Server{}
	param := &pb.EmbossRequest{}
	instantStockList, err := server.GetDataForEmboss(ctx, param)
	if err != nil {
		t.Error(err)
		return
	}

	if len(instantStockList.InstantStock) <= 0 {
		t.Error("no data found")
		return
	}

	utils.LogToFile(fmt.Sprintf("TestGetDataForEmboss result: %v", instantStockList.InstantStock))
}
