package kata4search

import (
	"fmt"
	"reflect"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchCities_ReturnTypeStringSlice(t *testing.T) {
	typeOfSearchCities := reflect.TypeOf(SearchCities("ascea"))
	assert.Equal(t,
		fmt.Sprint(typeOfSearchCities), "[]string",
		fmt.Sprintf("Returning wrong Type: %v", typeOfSearchCities),
	)
}

func TestSearchCities_LessThan2CharQuery(t *testing.T) {
	lenOfSearchCities := len(SearchCities(""))
	assert.Equal(t,
		lenOfSearchCities, 0,
		fmt.Sprintf("Search is Not Optimised returning %d options for empty query", lenOfSearchCities),
	)
}

func TestSearchCities_SearchCorrectQueryStart(t *testing.T) {
	searchOptions := SearchCities("Va")
	expectedOptions := []string{"Vancouver", "Valencia"}
	sort.Strings(searchOptions)
	sort.Strings(expectedOptions)
	assert.Equal(t, searchOptions, expectedOptions)
}

func TestSearchCities_SearchCorrectQueryCaseInsensitive(t *testing.T) {
	searchOptions := SearchCities("va")
	expectedOptions := []string{"Vancouver", "Valencia"}
	sort.Strings(searchOptions)
	sort.Strings(expectedOptions)
	assert.Equal(t, searchOptions, expectedOptions)
}

func TestSearchCities_SearchCorrectQueryMid(t *testing.T) {
	searchOptions := SearchCities("ape")
	expectedOptions := []string{"Budapest"}
	assert.Equal(t, searchOptions, expectedOptions)
}

func TestSearchCities_AsteriskAllCities(t *testing.T) {
	searchOptions := SearchCities("*")
	assert.Equal(t, searchOptions, allCities)
}
