package pokeapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func (c *Client) GetLocation(url string) (Area, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := c.httpClient.Do(req)
	if getErr != nil {
		return Area{}, getErr
	}

	// var body Area

	rawBody, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, rawBody)
	}
	if err != nil {
		log.Fatal(err)
	}

	area := Area{}
	if err := json.Unmarshal(rawBody, &area); err != nil {
		return Area{}, err
	}

	// fmt.Println(area.Count)
	return area, nil
}

func (c *Client) BytesToArea(b []byte) (Area, error) {
	return Area{}, nil
}
