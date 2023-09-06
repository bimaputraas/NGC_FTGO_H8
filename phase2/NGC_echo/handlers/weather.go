package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ngc_echo/models"
)

func GetWeather(city string) models.Weather {
	url := fmt.Sprintf("https://weather-by-api-ninjas.p.rapidapi.com/v1/weather?city=%s",city)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	req.Header.Add("X-RapidAPI-Key", "f5ef8d5bdamsh6881ca2f8bd4bc9p15ba10jsn27e1c43df0d3")
	req.Header.Add("X-RapidAPI-Host", "weather-by-api-ninjas.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	// responseBody, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	panic(err)
	// }

	// helpers.WriteResponseWithData(c,200,"Success get weather",string(responseBody))
	var responseBody models.Weather
	if err := json.NewDecoder(res.Body).Decode(&responseBody) ; err!=nil{
		panic(err)
	}

	return responseBody
}

// type Weather struct {
// 	CloudPct      float64   `json:"cloud_pct"`
// 	Temp  			float64 `json:"temp"`
// 	FeelsLike    float64 `json:"feels_like"`
// 	Humidity        float64 `json:"humidity"`
// 	MinTemp     float64  `json:"min_temp"`
// 	MaxTemp 	float64    `json:"max_temp"`
// 	WindSpeed      float64   `json:"wind_speed"`
// 	WindDegrees       float64 `json:"wind_degrees"`
// 	Sunrise       float64 `json:"sunrise"`
// 	Sunset       float64 `json:"sunset"`
// }