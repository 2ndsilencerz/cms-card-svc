package services

import (
	"context"
	"testing"

	"github.com/2ndSilencerz/cms-card-svc/models/pb"
	"google.golang.org/grpc"
)

var ctx context.Context = context.Background()

func TestGetCardList(t *testing.T) {

	conn, err := grpc.Dial(":9991", grpc.WithInsecure())
	if err != nil {
		t.Error(err)
		return
	}
	client := pb.NewCardListClient(conn)
	cardList, err := client.GetCardList(ctx,
		&pb.Page{
			FilterType:  "No Kartu",
			FilterValue: " ",
			Page:        "1",
			Limit:       "10",
		})
	if err != nil {
		t.Error(err)
		return
	}

	if len(cardList.GetVcard()) < 0 {
		t.Error("no data found")
		return
	}
}
