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

func TestPlaceAlien(t *testing.T) {
	planet, err := invasion.ParsePlanet(strings.NewReader("A\n"))
	require.NoError(t, err)

	sim := invasion.Simulation{planet}
	log := invasion.EventCollector{}
	sim.PlaceAlien(0, &log)

	require.Equal(t, 1, len(log.Events))
	require.Equal(t, "alien 1 descends on A", log.Events[0].String())
	require.Equal(t, planet.Cities[0], planet.Aliens[0].City)
	require.Equal(t, planet.Aliens[0], planet.Cities[0].Visitor)

	// Second alien will result in the destruction of the city
	sim.PlaceAlien(1, &log)

	require.Equal(t, 5, len(log.Events))
	require.Equal(t, "alien 2 descends on A", log.Events[1].String())
	require.Equal(t, "A has been destroyed by alien 2 and alien 1", log.Events[2].String())
	require.Equal(t, "alien 2 dies in a fight", log.Events[3].String())
	require.Equal(t, "alien 1 dies in a fight", log.Events[4].String())
	require.Equal(t, 0, len(planet.Aliens))
	// require.Equal(t, nil, planet.Cities[0].Visitor)
	require.Nil(t, planet.Cities[0].Visitor)

	// Third alien dies immediately in the destroyed city
	sim.PlaceAlien(2, &log)

	require.Equal(t, 7, len(log.Events))
	require.Equal(t, "alien 3 descends on A", log.Events[5].String())
	require.Equal(t, "alien 3 dies from radiation", log.Events[6].String())
	require.Equal(t, 0, len(planet.Aliens))
}
