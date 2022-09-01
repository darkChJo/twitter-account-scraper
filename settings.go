package main

import (
	"encoding/json"
	"os"
)

type SettingsData struct {
	Path     string
	Creators []string
}

func getSettings() *SettingsData {

	var settings SettingsData

	file, err := os.ReadFile("./settings.json")
	if err != nil {
		f, err := os.Create("./settings.json")
		if err != nil {
			panic(err)
		}
		data, err := json.MarshalIndent(settings, "", "\t")
		if err != nil {
			panic(err)
		}
		_, err2 := f.WriteString(string(data))
		if err2 != nil {
			panic(err)
		}
		f.Close()
	} else {

		json.Unmarshal(file, &settings)
	}
	return &settings
}
