package main

import (
	"encoding/json"
	"fmt"
	"goweather/notify"
	"io/ioutil"
	"net/http"
	"os"
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

func main() {
	// Get request on openweather api
	response, err := http.Get("http://api.openweathermap.org/data/2.5/weather?q=adelaide&units=metric&appid=" + key)
	// Funky golang error handling
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	data, _ := ioutil.ReadAll(response.Body)
	// Setting var result of type weatherInfo struct and then unmarshalling json result into the variable
	var result weatherInfo
	json.Unmarshal([]byte(data), &result)
	//fmt.Printf("The weather in %s is %f degrees celcius", result.Name, result.Main["feels_like"])
	error := notify.SendLinux("Weather Alert", "The weather in "+result.Name+" is "+fmt.Sprintf("%.2f", result.Main["temp"])+" degrees celcius.", "icons/cloud.png", "normal", "5000")
	if error != nil {
		fmt.Println(error)
	}

}
