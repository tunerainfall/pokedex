package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocation(url string) (AreaResponse, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return AreaResponse{}, err
	}

	// req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := c.httpClient.Do(req)
	if getErr != nil {
		return AreaResponse{}, getErr
	}

	rawBody, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, rawBody)
	}
	if err != nil {
		log.Fatal(err)
	}

	area := AreaResponse{}
	if err := json.Unmarshal(rawBody, &area); err != nil {
		return AreaResponse{}, err
	}

	// fmt.Println(area.Count)
	return area, nil
}

func (c *Client) GetPokemonsAtLocation(url string) (PokemonAreaResponse, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return PokemonAreaResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonAreaResponse{}, err
	}

	if res.StatusCode > 299 {
		return PokemonAreaResponse{}, fmt.Errorf("something wrong with request: Status %d", res.StatusCode)
	}

	rawBody, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return PokemonAreaResponse{}, err
	}

	pokemon := PokemonAreaResponse{}
	if err := json.Unmarshal(rawBody, &pokemon); err != nil {
		return PokemonAreaResponse{}, err
	}

	return pokemon, nil
}

func (c *Client) CatchPokemon(url string) (PokemonResponse, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return PokemonResponse{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonResponse{}, err
	}

	pokemon := PokemonResponse{}

	// unmarshal body to pokemon
	rawBody, err := io.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		return PokemonResponse{}, err
	}

	if err := json.Unmarshal(rawBody, &pokemon); err != nil {
		return PokemonResponse{}, err
	}

	return pokemon, nil
}
