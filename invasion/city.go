package invasion

import (
	"io"
	"strings"
)

type Connection struct {
	Direction string
	Neighbor  string
}

// A structured representation of one line of the input file.
// This is to facilitate modularity and testing.
type PreParsedCity struct {
	Name      string
	Neighbors []Connection
}

type City struct {
	Name               string
	Neighbors          []*City
	NeighborDirections map[string]string // Maps the name of the neighbor to the cardinal direction
	Destroyed          bool
	// TODO: aliens "visiting" this city go here
}

func (c *City) Format(w io.Writer) {
	// TODO: trim destroyed neighbors
	io.WriteString(w, c.Name)
	for _, neighbor := range c.Neighbors {
		io.WriteString(w, " ")
		io.WriteString(w, c.NeighborDirections[neighbor.Name])
		io.WriteString(w, "=")
		io.WriteString(w, neighbor.Name)
	}
}

func (c *City) ToString() string {
	var buf strings.Builder
	c.Format(&buf)
	return buf.String()
}

// For testing purposes only
func (c *City) CityToPreParsedCity() PreParsedCity {
	var preParsedCity PreParsedCity
	preParsedCity.Name = c.Name
	for _, neighbor := range c.Neighbors {
		preParsedCity.Neighbors = append(preParsedCity.Neighbors, Connection{
			Direction: c.NeighborDirections[neighbor.Name],
			Neighbor:  neighbor.Name,
		})
	}
	return preParsedCity
}
