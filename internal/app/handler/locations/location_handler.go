package locations

import (
	"github.com/gin-gonic/gin"
	"github.com/rodrigorahman/estados_cidades_api_go_codebase/internal/app/handler/locations/dto"
	"github.com/rodrigorahman/estados_cidades_api_go_codebase/internal/domain/location/repositories"
)

type LocationHandler struct {
	locationRepository repositories.LocationRepositoryInterface
}

func NewLocationHandler(locationRepository repositories.LocationRepositoryInterface) *LocationHandler {
	return &LocationHandler{locationRepository: locationRepository}
}
func (l *LocationHandler) FindAllStates(c *gin.Context) {
	states, err := l.locationRepository.GetStates()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var statesResponse []dto.StateResponse

	for _, s := range states {
		statesResponse = append(statesResponse, dto.StateResponse{
			Acronym: s.Acronym,
			Name:    s.Name,
		})
	}
	c.JSON(200, statesResponse)
}

func (l *LocationHandler) FindCitiesByState(c *gin.Context) {
	stateAcronym := c.Param("state")
	cities, err := l.locationRepository.GetCitiesByState(stateAcronym)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var citiesResponse []dto.CityResponse

	for _, c := range cities {
		citiesResponse = append(citiesResponse, dto.CityResponse{
			Name: c.Name,
			Ibge: c.Ibge,
		})
	}
	c.JSON(200, citiesResponse)
}
