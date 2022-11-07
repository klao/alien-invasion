package invasion

import "strings"

type Event interface {
	String() string
}

type EventLogger interface {
	LogEvent(event Event)
}

type StdoutEventLogger struct{}

func (l *StdoutEventLogger) LogEvent(event Event) {
	println(event.String())
}

var EventPrinter = &StdoutEventLogger{}

type EventCollector struct {
	Events []Event
}

func (c *EventCollector) LogEvent(event Event) {
	c.Events = append(c.Events, event)
}

////////////////////////////////////////////////////////////////////////////////////////////////////
// Various events

type GenericEvent struct {
	Description string
}

func (e *GenericEvent) String() string {
	return e.Description
}

type AlienDescentEvent struct {
	Alien string
	City  string
}

func (e *AlienDescentEvent) String() string {
	return e.Alien + " descends on " + e.City
}

type AlienMoveEvent struct {
	Alien string
	From  string
	To    string
}

func (e *AlienMoveEvent) String() string {
	return e.Alien + " moves from " + e.From + " to " + e.To
}

type CityDestroyedEvent struct {
	City   string
	Aliens []string
}

func (e *CityDestroyedEvent) String() string {
	var buf strings.Builder
	buf.WriteString(e.City)
	buf.WriteString(" has been destroyed by ")
	for i, alien := range e.Aliens {
		if i > 0 {
			if i == len(e.Aliens)-1 {
				// Last alien
				if len(e.Aliens) > 2 {
					// Let's make it generic and allow more than two aliens. Use the Oxford comma
					// for clarity in that case.
					buf.WriteString(", and ")
				} else {
					buf.WriteString(" and ")
				}
			} else {
				buf.WriteString(", ")
			}
		}
		buf.WriteString(alien)
	}
	return buf.String()
}

type AlienDeathReason int

const (
	AlienDeathReasonFight     AlienDeathReason = iota
	AlienDeathReasonRadiation AlienDeathReason = iota
)

type AlienDiesEvent struct {
	Alien  string
	Reason AlienDeathReason
}

func (e *AlienDiesEvent) String() string {
	switch e.Reason {
	case AlienDeathReasonFight:
		return e.Alien + " dies in a fight"
	case AlienDeathReasonRadiation:
		return e.Alien + " dies from radiation"
	default:
		panic("Unknown alien death reason")
	}
}
