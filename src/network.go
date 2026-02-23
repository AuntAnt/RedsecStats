package main

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"slices"

	"github.com/AuntAnt/RedsecStats/src/models"
	"github.com/AuntAnt/RedsecStats/src/utils"
)

const (
	playedDuo = "matches_gm_graniteDuo"
	playedBr  = "matches_gm_brsquad"
	stuns     = "tp_gad_gren_stun"
	revives   = "revives_gm_granite"
)

func fetchStatistic(url string) (*StatResult, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	log.Printf("Response code: %s\n", resp.Status)

	if resp.StatusCode == 404 {
		return nil, errors.New("Player not found")
	} else if resp.StatusCode != 200 {
		return nil, errors.New("Something went wrong")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// write full response body to file in debug mode
	if utils.Debug {
		writeResponseBody(body)
	}
	result, err := unmarshalPlayerStats(body)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func unmarshalPlayerStats(data []byte) (*StatResult, error) {
	var playerStats models.RawData

	err := json.Unmarshal(data, &playerStats)
	if err != nil {
		return nil, err
	}

	categoryFields := playerStats.Stats[0].Categories[0].Fields
	if len(categoryFields) == 0 {
		return nil, errors.New("Nothing found")
	}

	var totalRSRevives int
	var totalStuns int
	var totalBrPlayed int

	for _, catField := range categoryFields {
		switch catField.Name {
		case revives:
			totalRSRevives = catField.Value
		case stuns:
			if slices.ContainsFunc(catField.Fields, func(f models.Fields) bool { return f.CheckIfGranite() }) {
				totalStuns += catField.Value
			}
		case playedDuo, playedBr:
			totalBrPlayed += catField.Value
		}
	}

	result := StatResult{
		revives: totalRSRevives,
		stuns:   totalStuns,
		played:  totalBrPlayed,
	}

	return &result, nil
}

type StatResult struct {
	revives int
	stuns   int
	played  int
}
