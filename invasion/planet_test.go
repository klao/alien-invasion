package invasion_test

import (
	"testing"

	"github.com/klao/alien-invasion/invasion"

	"github.com/stretchr/testify/require"
)

// Test the second stage of the parsing process
func TestPlanetFromPreParsedPlanet(t *testing.T) {
	preParsedPlanet := invasion.PreParsedPlanet{
		{Name: "A", Neighbors: []invasion.Connection{{"north", "B"}, {"east", "C"}}},
		{Name: "B", Neighbors: []invasion.Connection{{"south", "A"}}},
		{Name: "C", Neighbors: []invasion.Connection{{"west", "A"}}},
	}
	planet := invasion.PlanetFromPreParsedPlanet(preParsedPlanet)

	require.Equal(t, preParsedPlanet, planet.PlanetToPreParsedPlanet())
	require.Equal(t, "A north=B east=C\nB south=A\nC west=A\n", planet.ToString())
}
