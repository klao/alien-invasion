package invasion_test

import (
	"testing"

	"github.com/klao/alien-invasion/invasion"

	"github.com/stretchr/testify/require"
)

func TestCityDestroyedEvent(t *testing.T) {
	event := invasion.CityDestroyedEvent{
		City:   "Bar",
		Aliens: []string{"X"},
	}
	require.Equal(t, "Bar has been destroyed by X", event.String())

	event.Aliens = []string{"X", "Y"}
	require.Equal(t, "Bar has been destroyed by X and Y", event.String())

	event.Aliens = []string{"X", "Y", "Z"}
	require.Equal(t, "Bar has been destroyed by X, Y, and Z", event.String())
}
