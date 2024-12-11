package datasources

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/rodrigorahman/estados_cidades_api_go_codebase/internal/domain/location/entities"
	"github.com/rodrigorahman/estados_cidades_api_go_codebase/internal/infraestructure/repositories/location/dto"
	"log"
	"net/http"
	"time"
)

type LocationBrasilApiDatasource struct {
}

func NewLocationBrasilApiDatasource() *LocationBrasilApiDatasource {
	return &LocationBrasilApiDatasource{}
}

func (l LocationBrasilApiDatasource) GetStates() ([]entities.StateEntity, error) {
	httpClient := http.Client{
		Timeout: 20 * time.Second,
	}
	req, err := http.NewRequest("GET", "https://brasilapi.com.br/api/ibge/uf/v1", nil)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req = req.WithContext(ctx)

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao buscar estados: ", err)
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Erro ao fechar o Body de busca de estados: %v", err)
		}
	}()

	var brResp []dto.BrasilApiStateResponse

	err = json.NewDecoder(resp.Body).Decode(&brResp)
	if err != nil {
		fmt.Println("Erro ao decodificar estados: ", err)
		return nil, err
	}

	var states []entities.StateEntity
	for _, s := range brResp {
		states = append(states, entities.StateEntity{
			Acronym: s.Sigla,
			Name:    s.Nome,
		})
	}

	return states, nil

}

func (l LocationBrasilApiDatasource) GetCitiesByState(stateAcronym string) ([]entities.CityEntity, error) {
	httpClient := http.Client{
		Timeout: 20 * time.Second,
	}
	req, err := http.NewRequest(
		"GET",
		fmt.Sprintf(
			"https://brasilapi.com.br/api/ibge/municipios/v1/%v?providers=dados-abertos-br,gov,wikipedia",
			stateAcronym,
		),
		nil,
	)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	req = req.WithContext(ctx)

	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println("Erro ao buscar cidades por estado: ", err)
		return nil, err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("Erro ao fechar o Body de busca de cidades: %v", err)
		}
	}()

	var brResp []dto.BrasilApiCityResponse

	err = json.NewDecoder(resp.Body).Decode(&brResp)
	if err != nil {
		fmt.Println("Erro ao decodificar estados: ", err)
		return nil, err
	}

	var cities []entities.CityEntity
	for _, s := range brResp {
		cities = append(cities, entities.CityEntity{
			Name: s.Nome,
			Ibge: s.Ibge,
		})
	}

	return cities, nil
}
