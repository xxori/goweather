package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const key string = "7adac862fe4f6f846740870350185838"

//Struct representing openweathermap api response
type weatherInfo struct {
	Coord   map[string]float64
	Weather []struct {
		ID          int
		Main        string
		Description string
		Icon        string
	}
	Main   map[string]float64
	Wind   map[string]float64
	Clouds map[string]interface{}
	Rain   map[string]interface{}
	Snow   map[string]interface{}
	Dt     int
	Sys    struct {
		Type    int
		ID      int
		Message float64
		Country string
		Sunrise int
		Sunset  int
	}
	Timezone int
	Name     string
	Cod      int
}

func getWeather(location string) (weatherInfo, error) {
	// Get request on openweather api
	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=" + location + "&units=metric&appid=" + key)
	// Funky golang error handling
	if err != nil {
		return weatherInfo{}, err
	}
	data, err := ioutil.ReadAll(response.Body)

	// Yet more funy go error handling
	if err != nil {
		return weatherInfo{}, err
	}

	// Setting var result of type weatherInfo struct and then unmarshalling json result into the variable
	var result weatherInfo
	json.Unmarshal([]byte(data), &result)
	return result, nil
}

// Linux notification sending using libnotf
func notifSend(title string, message string, icon string, urgency string, delay string) error {
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	dir := currentDir + "/" + icon

	cmd := exec.Command("notify-send", "-i", dir, "-u", urgency, "-t", delay, title, message)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	read := bufio.NewReader(os.Stdin)
	location, _ := read.ReadString('\n')
	location = strings.Replace(location, "\n", "", -1)
	resp, err := getWeather(location)
	if err != nil {
		panic(err)
	}
	if resp.Cod == 0 {
		panic("Invalid city")
	}
	fmt.Println(resp)
}
