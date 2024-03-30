# poke-it

POKE-IT is a REPL command line game, it fetches directly from the PokeApi in real time, the users can traverse various worlds , capture pokemons if there powers permit, lookup statistics for captured pokemons, etc. it stores the data in a real time custom cache implemented using Sync mutexes increasing the efficiency of the program. 

## Usage


## Build It Locally

It requires golang to be installed and available on the command line client.

clone the repository and select the folder
```
git clone https://github.com/jaydee029/Poke-it
cd Poke-it
```

Build the executable and run the executable to start the REPL backend
For unix based systems/WSL
```
go build -o pokeit
./pokeit
```

For windows cmd
```
go build
poke-it.exe
```

For the program to be globally available add the executable to your system path, and call it from anywhere in the system.

## Salient Features

- **Custom Cache Implementation:** Caching has been implemented from scratch for this program, Sync mutexes has been used for this purpose , which store the data in a key value pair format in the form of bytes, for a certain time period. If a user enters a command , the program first checks the chache , if the key-value pair exists then data is displayed else, the program fetches it from the api.
- **Platform Agnostic:** The program is platform agnostic, ie it works for most of command line clients , including the ones based on windows/Unix/Mac type operating systems.
The only requirement being that Golang is installed. 
