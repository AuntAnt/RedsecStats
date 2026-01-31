package models

import "strings"

type RawData struct {
	Stats []PlayerStats `json:"playerStats"`
}

type PlayerStats struct {
	Categories []Category `json:"categories"`
}

type Category struct {
	Name   string          `json:"catName"`
	Fields []CategoryField `json:"catFields"`
}

type CategoryField struct {
	Name   string   `json:"name"`
	Value  int      `json:"value"`
	Fields []Fields `json:"fields"`
}

type Fields struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (f Fields) CheckIfGranite() bool {
	return strings.Contains(f.Value, "Granite")
}
