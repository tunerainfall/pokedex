package pokeservice

import (
	"encoding/json"
	"fmt"

	"github.com/beevk/pokedex/internal/pokeapi"
	"github.com/beevk/pokedex/internal/pokecache"
)

type Service struct {
	client pokeapi.Client
	cache  pokecache.Cache
}

func NewService(client pokeapi.Client, cache pokecache.Cache) Service {
	return Service{
		client: client,
		cache:  cache,
	}
}

func (s Service) GetArea(url string) (pokeapi.Area, error) {
	if data, ok := s.cache.Get(url); ok {
		area := pokeapi.Area{}
		if err := json.Unmarshal(data, &area); err != nil {
			return pokeapi.Area{}, err
		}
		fmt.Printf("cache found for %s\n", url)

		return area, nil
	}

	area, err := s.client.GetLocation(url)
	if err != nil {
		return pokeapi.Area{}, fmt.Errorf("error from api: %w", err)
	}

	areaAsBytes, err := json.Marshal(area)
	if err != nil {
		return pokeapi.Area{}, err
	}

	if err := s.cache.Add(url, areaAsBytes); err != nil {
		return pokeapi.Area{}, err
	}

	return area, nil
}
