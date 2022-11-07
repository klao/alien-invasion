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

	// No superfluous spaces at the end, even if there are no neighbors
	city.Neighbors = nil
	require.Equal(t, "A", city.ToString())
}
