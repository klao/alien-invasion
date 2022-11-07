package invasion

import (
	"io"
	"strings"
)

type City struct {
	Name               string
	Neighbors          []*City
	NeighborDirections map[string]string // Maps the name of the neighbor to the cardinal direction
	Destroyed          bool
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
