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

	// in API ps4 and xboxone available as well,
	// but BF6 not availble on those platforms.
	// left them commented for a while
	availablePlatforms := map[int]string{
		0: "pc(ea)",
		1: "psn",
		2: "ps5",
		3: "xbox",
		4: "xboxseries",
		// 5: "ps4",
		// 6: "xboxone",
	}

	fmt.Println("  -- IMPORTANT --")
	fmt.Println("If your PSN/Xbox account name is different from your EA account name, select PC and enter your EA account")
	fmt.Println("\nSelect platform from available (enter only digit)")
	keys := getKeys(availablePlatforms)

	for key := range keys {
		val := availablePlatforms[key]
		fmt.Printf(" %d: %s\n", key, val)
	}
	fmt.Scan(&platform)

	result, found := availablePlatforms[platform]

	if !found {
		fmt.Println("You select unsupported platform, setted PC by default")
		return "pc"
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
