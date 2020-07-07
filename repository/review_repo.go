package repository

import (
	"context"
	"footcer-backend/model"
)

type ReviewRepository interface {
	AddReview(context context.Context, review model.Review) (model.Review, error)
}
