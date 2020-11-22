package repository

import (
	"context"
	"footcer-backend/model"
)

type StadiumRepository interface {
	StadiumInfo(context context.Context, userId string) (interface{}, error)
	StadiumInfoForID(context context.Context, stadiumID string) (interface{}, error)
	StadiumUpdate(context context.Context, stadium model.Stadium, role int8) (model.Stadium, error)

	StadiumCollageUpdate(context context.Context, stadiumColl model.StadiumCollage) (model.StadiumCollage, error)
	StadiumCollageAdd(context context.Context, stadiumColl model.StadiumCollage) (model.StadiumCollage, error)
	StadiumCollageDelete(context context.Context, idCollage string) error

	StadiumDetailsAdd(context context.Context, stadiumDetails model.StadiumDetails) (model.StadiumDetails, error)
	SearchStadiumLocation(context context.Context, latitude string, longitude string) ([]model.Stadium, error)
	SearchStadiumName(context context.Context, name string) ([]model.Stadium, error)

	StadiumDetailsInfoForStadiumCollage(context context.Context, stadiumCollageId string, date string) (interface{}, error)
	StadiumDetailsUpdateForStadiumCollage(context context.Context, details model.StadiumDetails) (interface{}, error)

	StadiumUploadImages(context context.Context, images model.Images) (interface{}, error)
}
