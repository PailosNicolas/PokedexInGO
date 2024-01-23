# PokedexInGO
This is a small project in Go from  [boot.dev](https://www.boot.dev/) where you make a REPL Pokedex (kind of).
While there is a video walkthrough I stayed away from it as much as I could (although I followed it a lot in the cache part), with this small project I have learned about requests, json handling, error handling, and basic cache implementation.

### Commands implemented:
- help
- exit
- explore
- catch
- inspect
- map
- mapb
- pokedex

### Tiny things I added/changed:
- I didn't quite uderstand the explore part since you give it a location name but the location itself does not have pokemons, the area in it does so I request the areas of the given location and return a list of pokemon in those areas (without repeating them).
- Added a really basic shiny calculation (I know it doesn't make sense but I wanted it)
- Added the option to give the catched pokemon a nickname

### Things I want to add:
- Random encounters
- Experience/lvl up/evolve
- Map traversal

Maybe I won't add them here since I think shifts the scope of this little project.


#
If you happen to be reading this please let me know of any mistake, improvement or anything, I will appreciate it a lot.