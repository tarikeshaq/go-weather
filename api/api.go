package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const weatherEndPoint = `http://api.openweathermap.org/data/2.5/weather?q=`

type WeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   float64 `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

func GetTemp(city string) (string, error) {
	weather, err := GetWeatherInfo(city)
	if err != nil {
		return "", err
	}
	temp := weather.Main.Temp
	val := strconv.FormatFloat(temp, 'f', -1, 64)
	if err != nil {
		return "", fmt.Errorf("Could not parse tempreature: %v", err)
	}
	return val, nil
}

func GetWind(city string) (string, error) {
	weather, err := GetWeatherInfo(city)
	if err != nil {
		return "", err
	}
	wind := weather.Wind.Speed
	val := strconv.FormatFloat(wind, 'f', -1, 64)
	if err != nil {
		return "", fmt.Errorf("Could not parse wind speed: %v", err)
	}
	return val, nil
}

func GetMain(city string) (string, error) {
	weather, err := GetWeatherInfo(city)
	if err != nil {
		return "", err
	}

	if len(weather.Weather) < 1 {
		return "", fmt.Errorf("No main weather information available")
	}

	return weather.Weather[0].Main, nil
}

func GetWeatherInfo(city string) (*WeatherResponse, error) {
	resp, err := http.Get(weatherEndPoint + city + "&APPID=" + os.Getenv("WEATHER_KEY"))
	if err != nil {
		return nil, fmt.Errorf("Could not get from api: %v", err)
	}
	defer resp.Body.Close()
	var weather WeatherResponse
	decoder := json.NewDecoder(resp.Body)
	decoder.Decode(&weather)
	return &weather, nil
}
