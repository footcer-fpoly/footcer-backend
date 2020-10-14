package repository

import (
	"context"
	"footcer-backend/model"
)

type StadiumRepository interface {
	StadiumInfo(context context.Context, userId string) (interface{}, error)
	StadiumInfoForID(context context.Context, stadiumID string) (interface{}, error)
	StadiumUpdate(context context.Context, stadium model.Stadium) (model.Stadium, error)
	StadiumCollageUpdate(context context.Context, stadiumColl model.StadiumCollage) (model.StadiumCollage, error)
	StadiumCollageAdd(context context.Context, stadiumColl model.StadiumCollage) (model.StadiumCollage, error)
	StadiumDetailsAdd(context context.Context, stadiumDetails model.StadiumDetails) (model.StadiumDetails, error)
	SearchStadiumLocation(context context.Context, latitude string, longitude string) ([]model.Stadium, error)
	SearchStadiumName(context context.Context, name string) ([]model.Stadium, error)
}
