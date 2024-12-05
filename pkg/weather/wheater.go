package weather

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/fatih/color"
)

type Weather struct {
	Location struct {
		Name string `json:"name"`
	} `json:"location"`
	Current struct {
		TempC     float64 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
		} `json:"condition"`
	}
	Forecast struct {
		Forecastday []struct {
			Astro struct {
				Sunrise   string `json:"sunrise"`
				Sunset    string `json:"sunset"`
				MoonPhase string `json:"moon_phase"`
			} `json:"astro"`
			Hour []struct {
				TimeEpoch int64   `json:"time_epoch"`
				TempC     float64 `json:"temp_c"`
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
				ChanceOfRain float64 `json:"chance_of_rain"`
			} `json:"hour"`
		} `json:"forecastday"`
	} `json:"forecast"`
}

func GetWeather(loc string, apiKey string) (*Weather, error) {
	url := fmt.Sprintf("http://api.weatherapi.com/v1/forecast.json?key=%s&q=%s&days=1&aqi=yes&alerts=no", apiKey, loc)

	res, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("Weather API is down")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var weather Weather
	err = json.Unmarshal(body, &weather)
	if err != nil {
		return nil, err
	}

	return &weather, nil
}

func DisplayWeather(weather *Weather) {
	location, current, hours, sunrise, sunset := weather.Location, weather.Current, weather.Forecast.Forecastday[0].Hour, weather.Forecast.Forecastday[0].Astro.Sunrise, weather.Forecast.Forecastday[0].Astro.Sunset

	fmt.Printf("%s: %.0f C, %s %s\n", location.Name, current.TempC, current.Condition.Text, GetWeatherEmoji(current.Condition.Text))
	fmt.Printf("Sunrise: %s 🌄, Sunset: %s 🌇\n", sunrise, sunset)

	for _, hour := range hours {
		date := time.Unix(hour.TimeEpoch, 0)
		if date.Before(time.Now()) {
			continue
		}

		message := fmt.Sprintf("Time: %s, %.0fC, %s %s, Chance of Rain: %.0f%%\n", date.Format("15:04"), hour.TempC, hour.Condition.Text, GetWeatherEmoji(hour.Condition.Text), hour.ChanceOfRain)

		if hour.ChanceOfRain < 40 {
			fmt.Print(message)
		} else {
			color.Red(message)
		}
	}
}
