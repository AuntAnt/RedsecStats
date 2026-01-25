package main

import (
	"fmt"
	"log"
	"slices"
)

const (
	baseUrl = "https://api.gametools.network/bf6/stats/?name=%s&platform=%s&raw=true"
)

func main() {
	platform := setPlatform()
	username := setUsername()

	log.Printf("Fetching statistic for %s, platform: %s\n", username, platform)

	url := fmt.Sprintf(baseUrl, username, platform)
	fetchRSReviveStat(url)
}

func setUsername() string {
	var username string
	fmt.Println("Enter username:")
	fmt.Scan(&username)
	return username
}

func setPlatform() string {
	var platform int

	availablePlatforms := map[int]string{
		0: "pc",
		1: "psn",
		2: "ps5",
		3: "ps4",
		4: "xbox",
		5: "xboxone",
		6: "xboxseries",
	}

	fmt.Println("Select platform from available (enter only digit)")
	keys := getKeys(availablePlatforms)

	for key := range keys {
		val := availablePlatforms[key]
		fmt.Printf(" %d: %s\n", key, val)
	}
	fmt.Scan(&platform)

	result, found := availablePlatforms[platform]

	if !found {
		fmt.Println("You select unsupported platform, setted PC by default")
	}
	return result
}

func getKeys(m map[int]string) []int {
	keys := make([]int, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}
	slices.Sort(keys)
	return keys
}
