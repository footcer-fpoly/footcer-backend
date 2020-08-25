package repository

import (
	"context"
	"footcer-backend/model"
)

type TeamRepository interface {
	AddTeam(context context.Context, team model.Team) (model.Team, error)
	SearchWithPhoneMemberTeam(context context.Context, phone string) (model.User, error)
	AddMemberTeam(context context.Context, teamDetails model.TeamDetails) (model.TeamDetails, error)

	GetTeamForUser(context context.Context, userId string) (interface{}, error)
	GetTeamForID(context context.Context, teamId string) (interface{}, error)

	DeleteMember(context context.Context, userID string) error
	DeleteTeam(context context.Context, teamID string) error

	UpdateTeam(context context.Context, team model.Team) (model.Team, error)
	AcceptInvite(context context.Context, teamDetails model.TeamDetails) error
}
