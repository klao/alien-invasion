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
