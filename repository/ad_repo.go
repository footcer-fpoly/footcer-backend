package repository

import (
	"context"
	"footcer-backend/model"
)

type AdRepository interface {
	AddAd(context context.Context, ad model.Ad) (model.Ad, error)
}
