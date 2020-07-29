package repo_impl

import (
	"context"
	"footcer-backend/db"
	"footcer-backend/log"
	"footcer-backend/message"
	"footcer-backend/model"
	"footcer-backend/repository"
)

type ServiceRepoImpl struct {
	sql *db.Sql
}

func NewServiceRepo(sql *db.Sql) repository.ServiceRepository {
	return &ServiceRepoImpl{sql: sql}
}

func (s *ServiceRepoImpl) AddService(context context.Context, service model.Service) (model.Service, error) {
	query := `INSERT INTO public.service(
	service_id, name_service, price_service, image, stadium_id)
	VALUES (:service_id, :name_service, :price_service, :image, :stadium_id);`

	_, err := s.sql.Db.NamedExecContext(context, query, service)
	if err != nil {
		log.Error(err.Error())
		return service, message.SomeWentWrong
	}

	return service, nil
}

func (s *ServiceRepoImpl) DeleteService(context context.Context, serviceId string) error {
	queryDelete := `DELETE FROM public.service
	WHERE service_id = $1;`
	_, errDeleteTeam := s.sql.Db.ExecContext(context, queryDelete, serviceId)
	if errDeleteTeam != nil {
		log.Error(errDeleteTeam.Error())
		return message.SomeWentWrong
	}
	return nil
}
