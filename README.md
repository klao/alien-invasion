# Alien Invasion

This repository contains my (Mihaly Barasz) solution for the "Alien Invasion" interview
assignment from Ignite / Gno / NewTendermint.

## Assumptions

- City names do not contain spaces (or any other whitespace) or `=` signs. Any other Unicode
  character is ok.
- Every city will have a line describing it. That is, it's not possible that a city will appear
  as a neighbor of some city, but won't have a line describing it. It's ok to not have any
  outgoing connections from a city.
- Generally, the connections don't have to be bidirectional: city B can be accessible from city A
  without city A being accessible from city B. That is, there can be one-way roads.

## Design decisions

Alien placement and movement:

- Aliens are placed on the cities uniformly at random. If a second alien is placed on a city
  they destroy the city then and there and die in the process. If an alien is placed on a city
  that's already destroyed it dies immediately (from, let's say, radiation.) This way, it's
  possible to request an arbitrary number of aliens on an arbitrary map, the behavior is always
  defined.
- Within a round the aliens move either in the order they were placed or in random order
  (controlled by a command line flag). The fight condition is evaluated after every move (not
  just at the end of a round.) This way it's not possible that more than two aliens end up in
  a city.

## Development process

For my "stream of consciousness" / "brain dump" during development see my
[Development Log](DevelopmentLog.md), if desired. All final architectural decisions will be
documented here as well.
