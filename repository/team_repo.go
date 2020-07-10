package repository

import (
	"context"
	"footcer-backend/model"
)

type TeamRepository interface {
	AddTeam(context context.Context, team model.Team) (model.Team, error)
	SearchWithPhoneMemberTeam(context context.Context, phone string) (model.User, error)
	AddMemberTeam(context context.Context, teamDetails model.TeamDetails) (model.TeamDetails, error)
	GetTeam(context context.Context, teamId string) (interface{}, error)
}
