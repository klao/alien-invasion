package invasion

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
