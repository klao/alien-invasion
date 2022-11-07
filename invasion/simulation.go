package invasion

import "math/rand"

type Simulation struct {
	*Planet
}

func (s *Simulation) RemoveAlien(alien *Alien) {
	delete(s.Aliens, alien.Id)
}

func (s *Simulation) MoveAlien(alien *Alien, log EventLogger) {
	oldCity := alien.City
	if oldCity == nil {
		panic("Alien is not in a city")
	}

	newCity := oldCity.RandomNeighbor()
	if newCity == nil {
		log.LogEvent(&AlienTrappedEvent{alien.Name(), oldCity.Name})
		// If it's trapped, it won't ever move again
		s.RemoveAlien(alien)
		return
	}

	log.LogEvent(&AlienMoveEvent{alien.Name(), oldCity.Name, newCity.Name})
	s.moveAlienToCity(alien, newCity, log)
}

func (s *Simulation) moveAlienToCity(alien *Alien, newCity *City, log EventLogger) {
	alien.City = newCity

	other := newCity.Visitor
	if other != nil {
		log.LogEvent(&CityDestroyedEvent{newCity.Name, []string{alien.Name(), other.Name()}})
		newCity.Destroyed = true
		newCity.Visitor = nil
		log.LogEvent(&AlienDiesEvent{alien.Name(), AlienDeathReasonFight})
		log.LogEvent(&AlienDiesEvent{other.Name(), AlienDeathReasonFight})
		s.RemoveAlien(alien)
		s.RemoveAlien(other)
		return
	}

	newCity.Visitor = alien
}

func (s *Simulation) PlaceAlien(id int, log EventLogger) {
	alien := &Alien{id, nil}
	s.Aliens[id] = alien

	i := rand.Intn(len(s.Cities))
	city := s.Cities[i]

	log.LogEvent(&AlienDescentEvent{alien.Name(), city.Name})

	if city.Destroyed {
		log.LogEvent(&AlienDiesEvent{alien.Name(), AlienDeathReasonRadiation})
		// Record that the alien died in that city for the sake of completeness
		alien.City = city
		s.RemoveAlien(alien)
		return
	}

	s.moveAlienToCity(alien, city, log)
}

func (s *Simulation) PlaceAliens(n int, log EventLogger) {
	for i := 0; i < n; i++ {
		s.PlaceAlien(i, log)
	}
}
