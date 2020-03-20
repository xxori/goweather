package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
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
	// Get request on openweather api, docs at https://openweathermap.org/current
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
	// Unmarshalling the json data and passing a pointer to the empty weatherInfo struct result
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
	notif := flag.Bool("n", false, "Send a notification of output or print to console")
	respType := flag.String("t", "temp", "Determines response type (temp, weather)")
	flag.Parse()
	if len(flag.Args()) == 0 {
		exit("Please input a target city or place")
	}

	resp, err := getWeather(flag.Args()[0])
	if err != nil {
		exit(err.Error())
	}

	if resp.Cod == 0 {
		exit("Invalid target place")
	}
	var title, icon, text string
	switch *respType {
	case "temp":
		title = "Temperature in " + resp.Name
		text = fmt.Sprintf("Current: %.2f\nMax: %.2f\nMin: %.2f\nFeel Like: %.2f", resp.Main["temp"], resp.Main["temp_max"], resp.Main["temp_min"], resp.Main["feels_like"])
		icon = "icons/thermometer-half.png"
	case "weather":
		title = "Weather in " + resp.Name
		text = resp.Weather[0].Description
		icon = "icons/cloud.png"
	default:
		exit("Invalid response information")
	}
	if *notif {
		notifSend(title, text, icon, "normal", "5000")
	} else {
		fmt.Println(resp)
		fmt.Println(*respType)
	}
}

func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
