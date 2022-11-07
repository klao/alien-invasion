package invasion_test

import (
	"testing"

	"github.com/klao/alien-invasion/invasion"

	"github.com/stretchr/testify/require"
)

func TestCityToString(t *testing.T) {
	city := invasion.City{
		Name: "A",
		Neighbors: []*invasion.City{
			{Name: "B"},
			{Name: "C"},
		},
		NeighborDirections: map[string]string{
			"B": "north",
			"C": "east",
		},
	}
	require.Equal(t, "A north=B east=C", city.ToString())

	// Test for destroyed neighbors
	city.Neighbors[0].Destroyed = true
	require.Equal(t, "A east=C", city.ToString())

	// No superfluous spaces at the end, even if there are no neighbors
	city.Neighbors = nil
	require.Equal(t, "A", city.ToString())
}

func TestToPreParseCity(t *testing.T) {
	preParsedCity, err := invasion.PreParseCity("A north=B east=C")
	require.NoError(t, err)
	require.Equal(t, invasion.PreParsedCity{
		Name: "A",
		Neighbors: []invasion.Connection{
			{"north", "B"},
			{"east", "C"},
		},
	}, preParsedCity)

	// Test for invalid input
	_, err = invasion.PreParseCity("A north=B east=C west")
	require.Error(t, err)
	require.Equal(t, "invalid connection \"west\"", err.Error())
}
