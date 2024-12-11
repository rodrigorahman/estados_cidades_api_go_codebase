package repositories

import (
	"fmt"
	"github.com/rodrigorahman/estados_cidades_api_go_codebase/internal/domain/location/entities"
	"github.com/rodrigorahman/estados_cidades_api_go_codebase/internal/domain/location/repositories"
	"github.com/rodrigorahman/estados_cidades_api_go_codebase/internal/infraestructure/datasources"
)

type LocationRepository struct {
	locationsApi   *datasources.LocationBrasilApiDatasource
	locationsCache *datasources.LocationCacheDatasource
}

func NewLocationRepository(locationsApi *datasources.LocationBrasilApiDatasource, locationsCache *datasources.LocationCacheDatasource) repositories.LocationRepositoryInterface {
	return &LocationRepository{
		locationsApi:   locationsApi,
		locationsCache: locationsCache,
	}
}

func (l LocationRepository) GetStates() ([]entities.StateEntity, error) {
	states := l.locationsCache.GetStates()
	var err error
	if states == nil {
		states, err = l.locationsApi.GetStates()
		if err != nil {
			fmt.Println("Erro ao buscar estados: ", err)
			return nil, err
		}
		err = l.locationsCache.SetStates(states)
		if err != nil {
			fmt.Println("Erro ao salvar estados no cache: ", err)
		}
	}

	return states, nil
}

func (l LocationRepository) GetCitiesByState(stateAcronym string) ([]entities.CityEntity, error) {
	cities := l.locationsCache.GetCitiesByState(stateAcronym)
	var err error
	if cities == nil {
		cities, err = l.locationsApi.GetCitiesByState(stateAcronym)
		if err != nil {
			fmt.Println("Erro ao buscar cidades: ", err)
			return nil, err
		}
		err = l.locationsCache.SetCitiesByState(stateAcronym, cities)
		if err != nil {
			fmt.Println("Erro ao salvar cidades no cache: ", err)
		}
	}
	return cities, nil
}
