package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/klao/alien-invasion/invasion"
)

const numberOfRounds = 10000

const usageString = "Usage: simulate [flags] <map file> <number of aliens>\n\n"

var alienOrderRandom = flag.Bool("order_random", false, "Order of alien moves in each round is random")
var logCitiesOnly = flag.Bool("log_cities_only", false, "Only log city destruction events")
var friendlyAliens = flag.Bool("friendly_aliens", false, "Give aliens names to seem more friendly")

func init() {
	flag.Usage = func() {
		fmt.Print(usageString)
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 2 {
		flag.Usage()
		return
	}

	alienNumber, err := strconv.Atoi(flag.Arg(1))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Can't parse number of aliens: %v\n\n", err)
		flag.Usage()
		os.Exit(1)
	}

	if alienNumber <= 0 {
		fmt.Fprintf(os.Stderr, "Number of aliens must be positive\n\n")
		flag.Usage()
		os.Exit(1)
	}

	mapFile, err := os.Open(flag.Arg(0))
	if err != nil {
		panic(err)
	}

	planet, err := invasion.ParsePlanet(mapFile)
	if err != nil {
		panic(err)
	}

	simulation := invasion.Simulation{
		Planet:           planet,
		AlienOrderRandom: *alienOrderRandom,
		FriendlyAliens:   *friendlyAliens,
	}

	var eventLog invasion.EventLogger
	if *logCitiesOnly {
		eventLog = &invasion.OfficialLogger{}
	} else {
		eventLog = invasion.EventPrinter
	}

	// Seed the randomness
	rand.Seed(time.Now().UnixNano())

	simulation.PlaceAliens(alienNumber, eventLog)
	simulation.Run(numberOfRounds, eventLog)

	fmt.Println("\n--- Map of the scorched Earth --")
	planet.Format(os.Stdout)
}
