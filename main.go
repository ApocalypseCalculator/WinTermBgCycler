package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type Settings struct {
	WinTermSettings string `json:"wintermsettings"`
	PicsPathFile    string `json:"picspathfile"`
	PicsFolder      string `json:"picsfolder"`
	Interval        int    `json:"interval"`
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initialize() Settings {
	jsonFile, err := os.Open("config.json")
	check(err)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	settings := Settings{}
	_ = json.Unmarshal(byteValue, &settings)
	//fmt.Println("Successfully loaded config.json")
	return settings
}

func validImage(text string) bool {
	lower := strings.ToLower(text)
	if !strings.HasPrefix(lower, "&") && (strings.HasSuffix(lower, ".png") || strings.HasSuffix(lower, ".jpg") || strings.HasSuffix(lower, ".jpeg") || strings.HasSuffix(lower, ".webp")) {
		if _, err := os.Stat(text); err == nil {
			return true
		}
	}
	return false
}

func loadImages(settings Settings) []string {
	pathfile, err := os.Open(settings.PicsPathFile)
	check(err)
	defer pathfile.Close()
	scanner := bufio.NewScanner(pathfile)
	paths := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if validImage(text) {
			paths = append(paths, text)
		}
	}
	//fmt.Println("Successfully loaded pics path file")
	folder, err2 := ioutil.ReadDir(settings.PicsFolder)
	check(err2)
	for _, file := range folder {
		name, _ := filepath.Abs(path.Join(settings.PicsFolder, file.Name()))
		if validImage(name) {
			paths = append(paths, name)
		}
	}
	return paths
}

func setNewBg(settings Settings, image string) {
	defaultloc, _ := os.UserCacheDir()
	if settings.WinTermSettings == "" {
		defaultloc += "\\Packages\\Microsoft.WindowsTerminal_8wekyb3d8bbwe\\LocalState\\settings.json"
	} else {
		defaultloc = settings.WinTermSettings
	}
	jsonFile, err := os.Open(defaultloc)
	check(err)
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var result map[string]interface{}
	_ = json.Unmarshal(byteValue, &result)
	//fmt.Println("Successfully loaded Windows Terminal settings file")
	//POV: you are me and don't know how to write Go properly
	for i, _ := range result["profiles"].(map[string]interface{})["list"].([]interface{}) {
		result["profiles"].(map[string]interface{})["list"].([]interface{})[i].(map[string]interface{})["backgroundImage"] = image
	}
	raw, _ := json.MarshalIndent(result, "", "\t")
	os.WriteFile(defaultloc, raw, 0644)
	fmt.Println("Set new background as ", image)
	return
}

func interval(settings Settings, indx int) {
	paths := loadImages(settings)
	if len(paths) == 0 {
		fmt.Println("No images found, please add at least 1")
	} else {
		setNewBg(settings, paths[indx%len(paths)])
		indx++
	}
	time.Sleep(time.Duration(settings.Interval * int(time.Second)))
	interval(settings, indx)
	return
}

func main() {
	settings := initialize()
	interval(settings, rand.Intn(1000))
	return
}
