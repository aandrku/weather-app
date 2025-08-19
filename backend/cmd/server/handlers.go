package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (a *application) getWeather(w http.ResponseWriter, r *http.Request) {
	type APIResponse struct {
		Coord struct {
			Lon int `json:"lon"`
			Lat int `json:"lat"`
		} `json:"coord"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Base string `json:"base"`
		Main struct {
			Temp      float64 `json:"temp"`
			FeelsLike float64 `json:"feels_like"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  int     `json:"pressure"`
			Humidity  int     `json:"humidity"`
			SeaLevel  int     `json:"sea_level"`
			GrndLevel int     `json:"grnd_level"`
		} `json:"main"`
		Visibility int `json:"visibility"`
		Wind       struct {
			Speed float64 `json:"speed"`
			Deg   int     `json:"deg"`
			Gust  float64 `json:"gust"`
		} `json:"wind"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Dt  int `json:"dt"`
		Sys struct {
			Country string `json:"country"`
			Sunrise int    `json:"sunrise"`
			Sunset  int    `json:"sunset"`
		} `json:"sys"`
		Timezone int    `json:"timezone"`
		ID       int    `json:"id"`
		Name     string `json:"name"`
		Cod      int    `json:"cod"`
	}

	type Response struct {
		Main        string  `json:"main"`
		Description string  `json:"description"`
		Icon        string  `json:"icon"`
		Temp        float64 `json:"temp"`
		FeelsLike   float64 `json:"feels_like"`
		TempMin     float64 `json:"temp_min"`
		TempMax     float64 `json:"temp_max"`
		City        string  `json:"name"`
	}

	q := r.URL.Query()

	lat := q.Get("lat")
	lon := q.Get("lon")

	req, err := http.NewRequest("GET", a.config.apiURL, nil)
	if err != nil {
		log.Printf("failed to create request: %v", err)
		return
	}

	params := req.URL.Query()
	params.Add("lat", lat)
	params.Add("lon", lon)
	params.Add("appid", a.config.weatherAPIKey)
	params.Add("units", "imperial")
	req.URL.RawQuery = params.Encode()

	res, err := a.client.Do(req)
	if err != nil {
		log.Printf("failed to make request to the api: %v", err)
		return
	}
	defer res.Body.Close()

	var apiResponse APIResponse
	dec := json.NewDecoder(res.Body)

	if err = dec.Decode(&apiResponse); err != nil {
		log.Printf("failed unmarshal api response: %v", err)
		return
	}

	resp := Response{
		Main:        apiResponse.Weather[0].Main,
		Description: apiResponse.Weather[0].Description,
		Icon:        apiResponse.Weather[0].Icon,
		Temp:        apiResponse.Main.Temp,
		FeelsLike:   apiResponse.Main.FeelsLike,
		TempMin:     apiResponse.Main.TempMin,
		TempMax:     apiResponse.Main.TempMax,
		City:        apiResponse.Name,
	}

	data, err := json.Marshal(resp)
	if err != nil {
		log.Printf("faield to marshar response: %v", err)
		return
	}

	w.Write(data)
}
