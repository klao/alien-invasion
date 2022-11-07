package invasion_test

import (
	"strings"
	"testing"

	"github.com/klao/alien-invasion/invasion"

	"github.com/stretchr/testify/require"
)

const (
	input1 = "A north=B east=C\nB south=A\nC west=A\n"
)

var preParsedPlanet1 = invasion.PreParsedPlanet{
	{Name: "A", Neighbors: []invasion.Connection{{"north", "B"}, {"east", "C"}}},
	{Name: "B", Neighbors: []invasion.Connection{{"south", "A"}}},
	{Name: "C", Neighbors: []invasion.Connection{{"west", "A"}}},
}

// Test the second stage of the parsing process
func TestPlanetFromPreParsedPlanet(t *testing.T) {
	preParsedPlanet := preParsedPlanet1
	planet := invasion.PlanetFromPreParsedPlanet(preParsedPlanet)

	require.Equal(t, preParsedPlanet, planet.PlanetToPreParsedPlanet())
	require.Equal(t, input1, planet.ToString())

	// Test unidirectional connections
	preParsedPlanet = invasion.PreParsedPlanet{
		{Name: "A", Neighbors: []invasion.Connection{{"north", "B"}}},
		{Name: "B", Neighbors: []invasion.Connection{}},
	}
	planet = invasion.PlanetFromPreParsedPlanet(preParsedPlanet)
	require.Equal(t, preParsedPlanet, planet.PlanetToPreParsedPlanet())
	require.Equal(t, "A north=B\nB\n", planet.ToString())
}

// Test the first stage of the parsing process
func TestPreParsePlanet(t *testing.T) {
	preParsedPlanet, err := invasion.PreParsePlanet(strings.NewReader(input1))
	require.NoError(t, err)
	require.Equal(t, preParsedPlanet1, preParsedPlanet)
}

func TestFormatRemovesDestroyed(t *testing.T) {
	planet := invasion.PlanetFromPreParsedPlanet(preParsedPlanet1)
	planet.Cities[0].Destroyed = true
	require.Equal(t, "B\nC\n", planet.ToString())
}
