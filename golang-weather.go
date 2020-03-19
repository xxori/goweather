package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const key string = "7adac862fe4f6f846740870350185838"

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
	Clouds map[string]int
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

func main() {
	// Get request on openweather api
	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=hobart&units=metric&appid=" + key)
	// Funky golang error handling
	if err != nil {
		fmt.Printf("Weather request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		var result weatherInfo //map[string]interface{}
		json.Unmarshal([]byte(data), &result)
		//fmt.Printf("The weather in %s is %f degrees celcius", result.Name, result.Main["feels_like"])
		fmt.Println(result)

	}

}
