package test

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/2ndsilencerz/cms-card-svc/configs/utils"
	"github.com/2ndsilencerz/cms-card-svc/models/pb"
	"github.com/2ndsilencerz/cms-card-svc/repository"
	"github.com/2ndsilencerz/cms-card-svc/services"
)

var ctx context.Context = context.Background()

func TestGetCardList(t *testing.T) {

	// conn := ConnectionConfig()
	// client := pb.NewCardListClient(conn)
	server := &services.Server{}

	page := &pb.Page{
		FilterType:  "cardNo",
		FilterValue: "6274860010001736",
		Page:        "1",
		Limit:       "10",
	}

	// cardList, err := client.GetCardList(ctx, page)
	cardList, err := server.GetCardList(ctx, page)
	if err != nil {
		t.Error(err)
		return
	}

	if len(cardList.GetVcard()) <= 0 {
		t.Error("no data found")
		return
	}

	res := cardList.GetVcard()[0]
	utils.LogToFile(fmt.Sprintf("TestGetCardList result : %v", res))
	if res.CardNo != page.FilterValue {
		t.Error("result not match with test")
	}
}

func TestGetBlockedCard(t *testing.T) {

	// conn := ConnectionConfig()
	// client := pb.NewCardListClient(conn)
	server := &services.Server{}

	page := &pb.Page{
		FilterType:  "accFlag",
		FilterValue: "61003010013579",
		Page:        "1",
		Limit:       "10",
	}
	blockPage := &pb.BlockPage{
		Page:     page,
		Branch:   "001",
		OcBranch: "",
	}

	// cardList, err := client.GetCardBlockedList(ctx, blockPage)
	cardList, err := server.GetCardBlockedList(ctx, blockPage)

	if err != nil {
		t.Error(err)
		return
	}

	if len(cardList.GetVcard()) <= 0 {
		t.Error("no data found")
		return
	}

	res := cardList.GetVcard()[0]
	utils.LogToFile(fmt.Sprintf("TestGetBlockedCard result : %v", res))
	if res.CardNo == "" {
		t.Error("result not match with test")
	}
}

func TestSQLInject(t *testing.T) {

	// conn := ConnectionConfig()
	// client := pb.NewCardListClient(conn)
	server := &services.Server{}

	page := &pb.Page{
		FilterType:  "accFlag",
		FilterValue: "0 AND 1=1 || delete * from this;",
		Page:        "1",
		Limit:       "10",
	}

	// _, err = client.GetCardList(ctx, page)
	_, err := server.GetCardList(ctx, page)
	utils.LogToFile(fmt.Sprint(err))
	if err != nil && strings.Contains(err.Error(), repository.CardNoFailedToParseMessage) {
		return
	}

	t.Error("this test has failed")
}

func TestGetCardDetails(t *testing.T) {

	// conn := ConnectionConfig()
	// client := pb.NewCardListClient(conn)
	server := &services.Server{}
	vcard := &pb.VCard{
		CardNo:  "6274860010001736",
		AccFlag: "61003010013579",
	}

	// cardList, err := client.GetCardDetails(ctx, vcard)
	cardList, err := server.GetCardDetails(ctx, vcard)

	if err != nil {
		t.Error(err)
		return
	}

	if len(cardList.GetVcard()) <= 0 {
		t.Error("no data found")
		return
	}

	utils.LogToFile(fmt.Sprintf("TestGetCardDetails result : %v", cardList.GetVcard()))
	for _, v := range cardList.GetVcard() {
		if v.CardNo == vcard.CardNo {
			return
		}
	}

	t.Error("result not match with test")
}
