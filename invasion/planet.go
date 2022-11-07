package invasion

import (
	"bufio"
	"io"
	"strings"
)

// Structured representation of the input file.
type PreParsedPlanet = []PreParsedCity

type Planet struct {
	Cities []*City
	Aliens map[int]*Alien
}

func (p *Planet) Format(w io.Writer) {
	for _, city := range p.Cities {
		city.Format(w)
		io.WriteString(w, "\n")
	}
}

func (p *Planet) ToString() string {
	var buf strings.Builder
	p.Format(&buf)
	return buf.String()
}

func PlanetFromPreParsedPlanet(preParsedPlanet PreParsedPlanet) *Planet {
	planet := &Planet{
		Cities: make([]*City, len(preParsedPlanet)),
	}

	// First pass: create all cities
	nameToCity := make(map[string]*City, len(preParsedPlanet))
	for i, preParsedCity := range preParsedPlanet {
		planet.Cities[i] = &City{Name: preParsedCity.Name}
		nameToCity[preParsedCity.Name] = planet.Cities[i]
	}

	// Second pass: create connections between cities
	for i, preParsedCity := range preParsedPlanet {
		city := planet.Cities[i]
		city.Neighbors = make([]*City, len(preParsedCity.Neighbors))
		city.NeighborDirections = make(map[string]string, len(preParsedCity.Neighbors))
		for j, neighbor := range preParsedCity.Neighbors {
			city.Neighbors[j] = nameToCity[neighbor.Neighbor]
			city.NeighborDirections[neighbor.Neighbor] = neighbor.Direction
		}
	}

	return planet
}

// For testing purposes only
func (p *Planet) PlanetToPreParsedPlanet() PreParsedPlanet {
	var preParsedPlanet PreParsedPlanet
	preParsedPlanet = make([]PreParsedCity, 0, len(p.Cities))
	for _, city := range p.Cities {
		preParsedPlanet = append(preParsedPlanet, city.CityToPreParsedCity())
	}
	return preParsedPlanet
}

func PreParsePlanet(r io.Reader) (PreParsedPlanet, error) {
	var preParsedPlanet PreParsedPlanet

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		preParsedCity, err := PreParseCity(scanner.Text())
		if err != nil {
			return nil, err
		}
		preParsedPlanet = append(preParsedPlanet, preParsedCity)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return preParsedPlanet, nil
}

func ParsePlanet(r io.Reader) (*Planet, error) {
	preParsedPlanet, err := PreParsePlanet(r)
	if err != nil {
		return nil, err
	}
	return PlanetFromPreParsedPlanet(preParsedPlanet), nil
}
