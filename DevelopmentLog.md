# Initial thoughts after reading the assignment

The basic problem appears to be quite simple, but with a lot of room for interpreting details and
expressing additional, not strictly required functionality. My first impulse is to lean towards
simple direct solutions, which could be extended if desired. For example:

- Aliens could be implemented as goroutines to emphasize Go's wonderful concurrency capabilities.
  But, the basic problem strongly suggests synchronous behavior, so using goroutines would just
  add unnecessary complexity, without any benefit.
- Cardinal directions being specified from one city to another hints at obvious visualization
  possibility: placing the cities on a map and displaying the invasion on it in "real time" as
  it progresses. This can be tricky depending on the input, but an interesting extension option.
- For the basic problem the directions and their number don't matter, so I'm going to implement it
  as a generic graph, directions being arbitrary labels on the edges.
- Relatedly, the problem doesn't specify explicitly (though strongly suggests, by real-world
  associations) that connections between cities are bidirectional. I'm not going to assume it.
  This will make the code slightly more complicated: when a city is destroyed, we can't look up
  all the cities that have a connection to it. But it make the code more universal and we won't
  need to bake in the east-west, north-south correspondence etc.
- Remark about naming the aliens: Maybe look up the library that generates rememberable names for
  Docker containers. It would be fun to have NaughtyEisteins running around ruining cities. :D

# Basic types

There is a lot of potential for abstraction in this assignment. But abstraction, in particular
using interfaces instead of concrete types can be detrimental in the early stages of development.
So, I'm going to use simple concrete types as long as it makes sense.

# Parsing

To make test driven development easier I decided to split input parsing in two stages:

1. Parse input stream into very simple structured data (basically just a slice of words and
   word pairs).
2. Create the graph structure from the pre-parsed structure.

Both phases should be quite straightforward and easy to test.

# Aliens

Alien behavior is relatively straightforward, but we need to iron out a few unclear parts of the
specification.

The assignment talks about "each iteration", which suggests that the movement of the aliens is
(semi-) synchronous. This is easy to simulate, but we still need to decide if within one round
they move one-by-one or all at the same time. That is, do we need to test for "collisions" after
individual moves or only at the end of a round.

The first option is easier to implement, since in that case only one alien can be in a city at any
time (as soon as another moves in, it's immediately destroyed.) If we only test at the end of a
round, more than two aliens can end up in a city randomly… Thus, I will implement the first
option, and maybe later add a synchronous possibility.

There is also the question of the order the aliens move within one round, it could be either stable
(always the same) or random. It's easy to implement both, so I will do that.

The assignment specifies that the simulation stops after every alien moves at least 10000 times
(or all are dead), but an alien can get stuck in an isolated city (all neighbors get destroyed.)
I will assume that this alien also "moves" in every round, just ends up in the same place where
it started. This makes the end condition very simple: we just simulate for 10000 rounds or until
every alien is dead.

Question: how are aliens placed? Just randomly on any city? What if there is another alien there
already? Or randomly "without replacement"? Then what if more aliens are requested than cities?
My initial decision: place aliens "with replacement", that is randomly in any city. If there is
another alien there then the city is destroyed then and there. If an alien descends on a destroyed
city it dies immediately (radiation is too harsh.) This way we can simulate arbitrary number of
aliens. (Otherwise we are limited to at most `2 * #cities` aliens.)

# Events

During the simulation various events can happen: aliens descend on cities, aliens move around,
aliens fight and destroy the city they are in, etc. The assignment only asks us to print the events
where a city is destroyed, but I'd prefer to handle it in a more general way (for easier testing
and debugging, and just because it's more fun to have more descriptive events.)
