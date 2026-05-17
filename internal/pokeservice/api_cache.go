package pokeservice

import (
	"encoding/json"
	"fmt"

	"github.com/tunerainfall/pokedex/internal/pokeapi"
	"github.com/tunerainfall/pokedex/internal/pokecache"
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

func (s Service) GetArea(url string) (pokeapi.AreaResponse, error) {
	if data, ok := s.cache.Get(url); ok {
		area := pokeapi.AreaResponse{}
		if err := json.Unmarshal(data, &area); err != nil {
			return pokeapi.AreaResponse{}, err
		}
		fmt.Printf("cache found for %s\n", url)

		return area, nil
	}

	area, err := s.client.GetLocation(url)
	if err != nil {
		return pokeapi.AreaResponse{}, fmt.Errorf("error from api: %w", err)
	}

	areaAsBytes, err := json.Marshal(area)
	if err != nil {
		return pokeapi.AreaResponse{}, err
	}

	if err := s.cache.Add(url, areaAsBytes); err != nil {
		return pokeapi.AreaResponse{}, err
	}

	return area, nil
}

func (s Service) GetPokemons(city string) (pokeapi.PokemonAreaResponse, error) {
	url := pokeapi.LocationURL + "/" + city

	if data, ok := s.cache.Get(url); ok {
		pokemons := pokeapi.PokemonAreaResponse{}
		if err := json.Unmarshal(data, &pokemons); err != nil {
			return pokeapi.PokemonAreaResponse{}, err
		}
		fmt.Printf("cache found for %s\n", url)

		return pokemons, nil
	}

	pokemons, err := s.client.GetPokemonsAtLocation(url)
	if err != nil {
		return pokeapi.PokemonAreaResponse{}, err
	}

	pokemonsAsBytes, err := json.Marshal(pokemons)
	if err != nil {
		return pokeapi.PokemonAreaResponse{}, err
	}

	if err := s.cache.Add(url, pokemonsAsBytes); err != nil {
		return pokeapi.PokemonAreaResponse{}, err
	}

	return pokemons, nil
}

func (s Service) CatchPokemon(name string) (pokeapi.PokemonResponse, error) {
	url := pokeapi.PokemonURL + "/" + name

	// check cache
	if data, ok := s.cache.Get(url); ok {
		pokemon := pokeapi.PokemonResponse{}
		if err := json.Unmarshal(data, &pokemon); err != nil {
			return pokeapi.PokemonResponse{}, err
		}
		fmt.Printf("cache found for %s\n", url)

		return pokemon, nil
	}

	// call api
	pokemon, err := s.client.CatchPokemon(url)
	if err != nil {
		return pokeapi.PokemonResponse{}, err
	}

	// store in cache
	rawData, err := json.Marshal(pokemon)
	if err != nil {
		return pokeapi.PokemonResponse{}, err
	}
	if err := s.cache.Add(url, rawData); err != nil {
		return pokeapi.PokemonResponse{}, err
	}

	// return response
	return pokemon, nil
}
