package main

import (
	"encoding/json"
	"flag"
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
	Clouds map[string]float64
	Rain   map[string]float64
	Snow   map[string]float64
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

	// Yet more funky go error handling
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
	// Declaring arguments for cli usage
	notif := flag.Bool("n", false, "Send a notification of output or print to console")
	respType := flag.String("t", "temp", "Determines response type (temp, weather, wind, rain, snow)")
	flag.Parse()
	// If user doesn't enter target city program exits
	if len(flag.Args()) == 0 {
		exit("Please input a target city or place")
	}

	resp, err := getWeather(strings.Join(flag.Args(), " "))
	if err != nil {
		exit(err.Error())
	}

	if resp.Cod == 0 {
		exit("Invalid target place")
	}
	var title, icon, text string

	// Switch for what the weather information the user asked for
	switch *respType {
	case "temp":
		title = "Temperature in " + resp.Name
		text = fmt.Sprintf("Current: %.2f\nMax: %.2f\nMin: %.2f\nFeel Like: %.2f", resp.Main["temp"], resp.Main["temp_max"], resp.Main["temp_min"], resp.Main["feels_like"])
		icon = "icons/thermometer-half.png"
	case "weather":
		title = "Weather in " + resp.Name
		text = resp.Weather[0].Description
		icon = "icons/cloud.png"
	case "wind":
		title = "Wind in " + resp.Name
		text = fmt.Sprintf("Speed: %.2fm/s\nDirection: %.2f degrees", resp.Wind["speed"], resp.Wind["deg"])
		icon = "icons/wind.png"
	case "rain":
		title = "Rain in " + resp.Name
		if resp.Rain["1h"] == 0 && resp.Rain["3h"] == 0 {
			text = "There is currently no rain in " + resp.Name
			icon = "icons/x.png"
		} else {
			text = fmt.Sprintf("Volume in last hour: %.2fmm\nVolume in last 3 hours: %.2fmm", resp.Rain["1h"], resp.Rain["3h"])
			icon = "icons/cloud-pouring.png"
		}
	case "snow":
		title = "Snow in " + resp.Name
		if resp.Snow["1h"] == 0 && resp.Snow["3h"] == 0 {
			text = "There is currently no snow in " + resp.Name
			icon = "icons/x.png"
		} else {
			text = fmt.Sprintf("Volume in last hour: %.2fmm\nVolume in last 3 hours: %.2fmm", resp.Snow["1h"], resp.Snow["3h"])
			icon = "icons/cloud-snow.png"
		}

	default:
		exit("Invalid response information")
	}
	if *notif {
		notifSend(title, text, icon, "normal", "10000")
	} else {
		fmt.Println(text)
	}
}

// Helper function
func exit(message string) {
	fmt.Println(message)
	os.Exit(1)
}
