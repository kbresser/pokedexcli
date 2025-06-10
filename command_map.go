package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type locationArea struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(page *config) error {
	startUrl := ""

	if page.Next == nil {
		startUrl = "https://pokeapi.co/api/v2/location-area?limit=20&offset=0"
	} else {
		startUrl = *page.Next
	}

	res, err := http.Get(startUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var locations locationArea
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return err
	}

	page.Next = &locations.Next
	page.Previous = locations.Previous

	for i := 0; i < len(locations.Results); i++ {
		fmt.Println(locations.Results[i].Name)
	}

	return nil
}

func commandMapb(page *config) error {
	startUrl := ""

	if page.Previous == nil {
		return fmt.Errorf("you're on the first page")
	} else {
		startUrl = *page.Previous
	}

	res, err := http.Get(startUrl)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var locations locationArea
	decoder := json.NewDecoder(res.Body)
	err = decoder.Decode(&locations)
	if err != nil {
		return err
	}

	page.Next = &locations.Next
	page.Previous = locations.Previous

	for i := 0; i < len(locations.Results); i++ {
		fmt.Println(locations.Results[i].Name)
	}

	return nil
}
