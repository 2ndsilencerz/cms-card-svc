package repository

import (
	"github.com/2ndsilencerz/cms-card-svc/configs/utils"
	"golang.org/x/net/context"
)

// Setting for context, limit, and filter
type Setting struct {
	Ctx         context.Context
	FilterType  string
	FilterValue string
	Limit       string
	Page        string
	LimitInt    int
	PageInt     int
	Total       int64
}

// Offsets calculated by Page and Limit
func (s *Setting) Offsets() (int, error) {
	s.LimitInt = utils.StrToInt(s.Limit)
	s.PageInt = utils.StrToInt(s.Page)
	if s.LimitInt == 0 && s.PageInt == 0 {
		return 0, utils.NewError("Failed to parse limit and page number")
	}
	offsets := (s.PageInt-1)*s.LimitInt - 1
	if offsets == -1 {
		offsets = 0
	}
	return offsets, nil
}

// TotalPage for paging
func (s *Setting) TotalPage() int32 {
	return int32(s.Total / 10)
}
