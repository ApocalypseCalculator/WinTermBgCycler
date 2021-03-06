package main

import (
	"fmt"
	"os"
)

func checkFiles() {
	//write default config
	if _, err := os.Stat("config.json"); err != nil {
		fmt.Println("Creating default configuration file")
		os.WriteFile("config.json", []byte("{\n\t\"wintermsettings\": \"\",\n\t\"picspathfile\": \"backgrounds/pics.txt\",\n\t\"picsfolder\": \"backgrounds\",\n\t\"interval\": 120\n}"), 0644)
	}
	settings := initialize()
	if settings.PicsPathFile == "backgrounds/pics.txt" {
		if _, err := os.Stat("backgrounds"); err != nil {
			fmt.Println("Creating default backgrounds folder")
			os.Mkdir("backgrounds", 0644)
		}
		if _, err := os.Stat(settings.PicsPathFile); err != nil {
			fmt.Println("Creating default backgrounds paths file")
			os.WriteFile(settings.PicsPathFile, []byte("& You can either place your pictures inside this folder or append the file path of the image to this file\n"), 0644)
		}
	}
}
