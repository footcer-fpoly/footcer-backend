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

	uuid "github.com/satori/go.uuid"
)

type TeamRepoImpl struct {
	sql *db.Sql
}

func NewTeamRepo(sql *db.Sql) repository.TeamRepository {
	return &TeamRepoImpl{
		sql: sql,
	}
}

func (t TeamRepoImpl) AddTeam(context context.Context, team model.Team) (model.Team, error) {
	queryCreate := `INSERT INTO public.team(
	team_id, name, level, place, description, avatar, background, leader_id, created_at, updated_at)
	VALUES (:team_id, :name, :level, :place, :description, :avatar, :background, :leader_id, :created_at, :updated_at);`

	_, err := t.sql.Db.NamedExecContext(context, queryCreate, team)
	if err != nil {
		log.Error(err.Error())
		return team, message.SomeWentWrong
	}
	var memberTeam = model.TeamDetails{
		TeamDetailsId: uuid.NewV1().String(),
		TeamId:        team.TeamId,
		UserId:        team.LeaderId,
		Accept:        "1",
		Role:          "1",
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	queryCreateMemberTeam := `INSERT INTO public.team_details(
	team_details_id, team_id, user_id, accept,role, created_at, updated_at)
	VALUES (:team_details_id, :team_id, :user_id, :accept,:role, :created_at, :updated_at);`
	_, errMemberTeam := t.sql.Db.NamedExecContext(context, queryCreateMemberTeam, memberTeam)
	if errMemberTeam != nil {
		log.Error(errMemberTeam.Error())
		return team, message.SomeWentWrong
	}
	if len(team.MemberList) > 0 {
		for _, element := range team.MemberList {
			memberTeam = model.TeamDetails{
				TeamDetailsId: uuid.NewV1().String(),
				TeamId:        team.TeamId,
				UserId:        element,
				Accept:        "0",
				Role:          "0",
				CreatedAt:     time.Now(),
				UpdatedAt:     time.Now(),
			}
			queryCreateMemberTeam := `INSERT INTO public.team_details(
				team_details_id, team_id, user_id, accept,role, created_at, updated_at)
				VALUES (:team_details_id, :team_id, :user_id, :accept,:role, :created_at, :updated_at);`
			_, errMemberTeam := t.sql.Db.NamedExecContext(context, queryCreateMemberTeam, memberTeam)
			if errMemberTeam != nil {
				log.Error(errMemberTeam.Error())
				return team, message.SomeWentWrong
			}
		}
	}

	return team, nil
}

func (t TeamRepoImpl) SearchWithPhoneMemberTeam(context context.Context, phone string) (model.User, error) {
	var user = model.User{}
	queryUser := `SELECT * FROM users WHERE users.phone = $1 AND role = $2`

	err := t.sql.Db.GetContext(context, &user, queryUser, phone, "0")
	if err != nil {
		if err == sql.ErrNoRows {
			return user, message.UserNotFound
		}
		log.Error(err.Error())
		return user, message.SomeWentWrong
	}

	return user, nil
}

func (t TeamRepoImpl) AddMemberTeam(context context.Context, teamDetails model.TeamDetails) (model.TeamDetails, error) {
	var user = model.User{}
	queryMemberExits := `SELECT user_id FROM public.team_details WHERE user_id = $1 AND team_id = $2`

	err := t.sql.Db.GetContext(context, &user, queryMemberExits, teamDetails.UserId, teamDetails.TeamId)

	if err != nil {
		if err == sql.ErrNoRows {
			queryCreateMemberTeam := `INSERT INTO public.team_details(
				team_details_id, team_id, user_id, accept, created_at, updated_at)
				VALUES (:team_details_id, :team_id, :user_id, :accept, :created_at, :updated_at);`
			_, errMemberTeam := t.sql.Db.NamedExecContext(context, queryCreateMemberTeam, teamDetails)
			if errMemberTeam != nil {
				log.Error(errMemberTeam.Error())
				return teamDetails, message.SomeWentWrong
			}
			return teamDetails, nil
		}
		log.Error(err.Error())
		return teamDetails, message.SomeWentWrong

	}
	return teamDetails, message.TeamMemberExits

}

func (t TeamRepoImpl) GetTeamForUser(context context.Context, userId string) (interface{}, error) {

	type teamTemp struct {
		model.Team
		model.TeamDetails `json:"member"`
	}

	var team = []teamTemp{}

	queryTeam := `SELECT team.*, team_details.role
	FROM public.team INNER JOIN public.team_details ON team_details.team_id = team.team_id WHERE team_details.user_id = $1;`

	err := t.sql.Db.SelectContext(context, &team,
		queryTeam, userId)
	if err != nil {
		log.Error(err.Error())
		return team, err
	}

	return team, nil
}
func (t TeamRepoImpl) GetTeamForID(context context.Context, teamId string) (interface{}, error) {

	type userMembertemp struct {
		model.TeamDetails 
		model.User `json:"user"`
	}
	type MemberList []userMembertemp


	type teamTemp struct {
		model.Team
		MemberList `json:"member"`
	}

	var team = teamTemp{}

	queryTeam := `SELECT team.* 
	FROM public.team WHERE team.team_id = $1;`

	err := t.sql.Db.GetContext(context, &team,
		queryTeam, teamId)
	if err != nil {
		log.Error(err.Error())
		return team, err
	}

	var memberList = []userMembertemp{}
	queryMemberTeam := `SELECT team_details.*, users.display_name, users.avatar
	FROM public.team_details INNER JOIN users ON users.user_id = team_details.user_id WHERE team_details.team_id = $1;`
	errMember := t.sql.Db.SelectContext(context, &memberList, queryMemberTeam, teamId)
	if errMember != nil {
		log.Error(errMember.Error())
		return memberList, errMember
	}
	team.MemberList = memberList


	return team, nil
}
