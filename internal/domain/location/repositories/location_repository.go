package repositories

import "github.com/rodrigorahman/estados_cidades_api_go_codebase/internal/domain/location/entities"

type LocationRepositoryInterface interface {
	GetStates() ([]entities.StateEntity, error)
	GetCitiesByState(stateAcronym string) ([]entities.CityEntity, error)
}
