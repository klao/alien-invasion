package invasion_test

import (
	"strings"
	"testing"

	"github.com/klao/alien-invasion/invasion"

	"github.com/stretchr/testify/require"
)

func TestMoveAlien(t *testing.T) {
	// Trapped alien
	planet, err := invasion.ParsePlanet(strings.NewReader("A"))
	require.NoError(t, err)

	planet.Aliens = make(map[int]*invasion.Alien)
	alien := &invasion.Alien{Id: 0, City: planet.Cities[0]}
	planet.Aliens[0] = alien

	sim := invasion.Simulation{planet}
	log := invasion.EventCollector{}
	sim.MoveAlien(alien, &log)

	require.Equal(t, 1, len(log.Events))
	require.Equal(t, "alien 1 is trapped in A", log.Events[0].String())
	require.Equal(t, 0, len(planet.Aliens))

	// Alien freely moving
	planet, err = invasion.ParsePlanet(strings.NewReader("A north=B\nB south=A\n"))
	require.NoError(t, err)

	planet.Aliens = make(map[int]*invasion.Alien)
	alien = &invasion.Alien{Id: 0, City: planet.Cities[0]}
	planet.Aliens[0] = alien

	sim = invasion.Simulation{planet}
	log = invasion.EventCollector{}
	sim.MoveAlien(alien, &log)

	require.Equal(t, 1, len(log.Events))
	require.Equal(t, "alien 1 moves from A to B", log.Events[0].String())
	require.Equal(t, planet.Cities[1], alien.City)
	require.Equal(t, alien, planet.Cities[1].Visitor)
}
