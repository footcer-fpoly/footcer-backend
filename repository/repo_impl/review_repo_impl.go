package repo_impl

import (
	"context"
	"database/sql"
	"footcer-backend/db"
	"footcer-backend/log"
	"footcer-backend/message"
	"footcer-backend/model"
	"footcer-backend/repository"
	"time"
)

type ReviewRepoImpl struct {
	sql *db.Sql
}

func NewReviewRepo(sql *db.Sql) repository.ReviewRepository {
	return &ReviewRepoImpl{sql: sql}
}

func (r ReviewRepoImpl) AddReview(context context.Context, review model.Review) (model.Review, error) {
	queryExits := `SELECT review_id
	FROM public.review WHERE user_id = $1 AND stadium_id = $2`

	var userId = ""
	errExits := r.sql.Db.GetContext(context, &userId, queryExits, review.User.UserId, review.StadiumId)
	if errExits != nil {

		if errExits == sql.ErrNoRows {
			query := `INSERT INTO public.review(
	review_id, content, rate, stadium_id, user_id, created_at_rv, updated_at_rv)
	VALUES (:review_id, :content, :rate, :stadium_id, :user_id, :created_at_rv,:updated_at_rv);`

			review.CreatedAt = time.Now()
			review.UpdatedAt = time.Now()

			_, err := r.sql.Db.NamedExecContext(context, query, review)
			if err != nil {
				log.Error(err.Error())
				return review, message.SomeWentWrong
			}
			return review, nil
		}
	}
	queryUpdate := `UPDATE public.review
			SET
			content = :content,
			rate = :rate,
			updated_at_rv 	  = COALESCE (:updated_at_rv, updated_at_rv)
			WHERE user_id = :user_id AND stadium_id = :stadium_id`

	review.UpdatedAt = time.Now()
	review.UserId = review.User.UserId

	_, err := r.sql.Db.NamedExecContext(context, queryUpdate, review)
	if err != nil {
		log.Error(err.Error())
		return review, message.SomeWentWrong
	}

	return review, nil
}
