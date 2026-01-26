package main

import (
	"encoding/json"
	"fmt"
	"github.com/AuntAnt/RedsecStats/src/models"
	"io"
	"log"
	"net/http"
)

func fetchRSReviveStat(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Response code: %d\n", resp.StatusCode)

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	result := unmarshalPlayerStats(body)
	fmt.Printf("Total revives in all RedSec modes: %d\n", result)
}

func unmarshalPlayerStats(data []byte) int {
	var playerStats models.RawData

	err := json.Unmarshal(data, &playerStats)
	if err != nil {
		log.Fatalln(err)
	}

	stats := playerStats.Stats
	if len(stats) == 0 {
		log.Fatalln("Nothing found")
	}

	fields := stats[0].Categories[0].Fields

	var totalRSRevives int
	for _, field := range fields {
		if field.Name == "revives_gm_granite" {
			totalRSRevives = field.Value
			break
		}
	}
	return totalRSRevives
}
