package repo_impl

import (
	"context"
	"footcer-backend/db"
	"footcer-backend/log"
	"footcer-backend/message"
	"footcer-backend/model"
	"footcer-backend/repository"
)

type ReviewRepoImpl struct {
	sql *db.Sql
}

func NewReviewRepo(sql *db.Sql) repository.ReviewRepository {
	return &ReviewRepoImpl{sql: sql}
}

func (r ReviewRepoImpl) AddReview(context context.Context, review model.Review) (model.Review, error) {
	query := `INSERT INTO public.review(
	review_id, content, rate, stadium_id, user_id, created_at, updated_at)
	VALUES (:review_id, :content, :rate, :stadium_id, :user_id, :created_at,:updated_at );`

	_, err := r.sql.Db.NamedExecContext(context, query, review)
	if err != nil {
		log.Error(err.Error())
		return review, message.SomeWentWrong
	}
	return review, nil

}
