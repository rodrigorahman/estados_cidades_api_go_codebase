package datasources

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/rodrigorahman/estados_cidades_api_go_codebase/internal/core"
	"github.com/rodrigorahman/estados_cidades_api_go_codebase/internal/domain/location/entities"
	"time"
)

const (
	StateKey = "states"
	CityKey  = "cities"
)

type LocationCacheDatasource struct {
	redisClient *core.RedisClient
}

func NewLocationCacheDatasource(redisClient *core.RedisClient) *LocationCacheDatasource {
	return &LocationCacheDatasource{redisClient: redisClient}
}

func (l LocationCacheDatasource) GetStates() []entities.StateEntity {
	statesCache, err := l.redisClient.Get(StateKey)

	if err != nil {
		if errors.Is(redis.Nil, err) {
			return nil
		}
		fmt.Printf("Error getting states from cache: %v", err)
		return nil
	}

	var states []entities.StateEntity
	err = json.Unmarshal([]byte(statesCache), &states)
	if err != nil {
		return nil
	}
	// 14343768
	return states
}

func (l LocationCacheDatasource) GetCitiesByState(state string) []entities.CityEntity {
	citiesCache, err := l.redisClient.Get(fmt.Sprintf("%s:%s", CityKey, state))
	if err != nil {
		if errors.Is(redis.Nil, err) {
			return nil
		}
		return nil
	}

	var cities []entities.CityEntity
	err = json.Unmarshal([]byte(citiesCache), &cities)
	if err != nil {
		return nil
	}
	return cities
}

func (l LocationCacheDatasource) SetStates(states []entities.StateEntity) error {
	statesBytes, err := json.Marshal(states)
	if err != nil {
		return err
	}
	err = l.redisClient.Set(StateKey, string(statesBytes), 180*24*time.Hour)
	if err != nil {
		return err
	}
	return nil
}

func (l LocationCacheDatasource) SetCitiesByState(state string, cities []entities.CityEntity) error {
	citiesBytes, err := json.Marshal(cities)
	if err != nil {
		return err
	}
	err = l.redisClient.Set(fmt.Sprintf("%s:%s", CityKey, state), string(citiesBytes), 90*24*time.Hour)
	if err != nil {
		return err
	}
	return nil
}
