package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/AuntAnt/RedsecStats/src/models"
)

const (
	playedDuo = "matches_gm_graniteDuo"
	playedBr  = "matches_gm_brsquad"
	stuns     = "tp_gad_gren_stun"
	revives   = "revives_gm_granite"
)

func fetchRSReviveStat(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Response code: %s\n", resp.Status)

	if resp.StatusCode == 404 {
		log.Fatalln("Player not found")
	} else if resp.StatusCode != 200 {
		log.Fatalln("Something went wrong")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	result := unmarshalPlayerStats(body)

	fmt.Println("\nRedSec statistic")
	fmt.Printf("  BR Played  |  %d  \n", result.played)
	fmt.Printf("  Stuns      |  %d  \n", result.stuns)
	fmt.Printf("  Revives    |  %d  \n", result.revives)
}

func unmarshalPlayerStats(data []byte) StatResult {
	var playerStats models.RawData

	err := json.Unmarshal(data, &playerStats)
	if err != nil {
		log.Fatalln(err)
	}

	fields := playerStats.Stats[0].Categories[0].Fields
	if len(fields) == 0 {
		log.Fatalln("Nothing found")
	}

	var totalRSRevives int
	var totalStuns int
	var totalBrPlayed int

	for _, field := range fields {
		switch field.Name {
		case revives:
			totalRSRevives = field.Value
		case stuns:
			for _, f := range field.Fields {
				if f.CheckIfGranite() {
					totalStuns += field.Value
				}
			}
		case playedDuo, playedBr:
			totalBrPlayed += field.Value
		}
	}

	return StatResult{
		revives: totalRSRevives,
		stuns:   totalStuns,
		played:  totalBrPlayed,
	}
}

type StatResult struct {
	revives int
	stuns   int
	played  int
}
