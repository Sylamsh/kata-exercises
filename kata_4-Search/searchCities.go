package kata4search

import "strings"

var allCities = []string{"Paris", "Budapest", "Skopje", "Rotterdam", "Valencia", "Vancouver", "Amsterdam", "Vienna", "Sydney", "New York City", "London", "Bangkok", "Hong Kong", "Dubai", "Rome", "Istanbul"}

func SearchCities(searchQuery string) []string {
	var selectedCities []string
	if searchQuery == "*" {
		return allCities
	}
	if len(searchQuery) < 2 {
		return selectedCities
	}
	for _, city := range allCities {
		if strings.Contains(strings.ToLower(city), strings.ToLower(searchQuery)) {
			selectedCities = append(selectedCities, city)
		}
	}
	return selectedCities
}
