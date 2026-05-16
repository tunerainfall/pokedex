package config

import (
	"github.com/beevk/pokedex/internal/pokeapi"
)

type AreaGetter interface {
	GetArea(url string) (pokeapi.Area, error)
}

type Config struct {
	Areas    AreaGetter
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
}
