package weather

import "strings"

func GetWeatherEmoji(condition string) string {
	condition = strings.TrimSpace(condition)
	switch condition {
	case "Sunny":
		return "☀️"
	case "Clear":
		return "🌞"
	case "Partly cloudy", "Partly Cloudy":
		return "🌤️"
	case "Cloudy":
		return "☁️"
	case "Overcast":
		return "🌥️"
	case "Mist", "Fog", "Haze":
		return "🌫️"
	case "Patchy rain possible", "Patchy light drizzle", "Light rain", "Patchy rain nearby", "Light rain shower", "Light drizzle":
		return "🌦️"
	case "Rain", "Heavy rain":
		return "🌧️"
	case "Thunderstorm":
		return "⛈️"
	case "Snow":
		return "❄️"
	default:
		return ""
	}
}
