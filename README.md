# NS-RPC-CLI
 A CLI reproduction of [NS-RPC](https://github.com/da532/ns-rpc)

## Introduction
This project is the result of frustrations I've experienced trying to wrangle WebViews on various Linux distros and needing a quick-and-dirty alternative to display my Nintendo Switch games on discord via Linux.
Eventually, the main application will feature full Linux compatibility but until this will suffice.

## Requirements
All that is required is at least [Golang](https://go.dev/) 1.21.2.

## Usage
1. First build the program by executing `go build` in the project's directory via your terminal.
2. Run the program, passing your game title as arguments e.g. `./NS-RPC-CLI splatoon 3` (A full games list can be found on the main repo).
3. Once gamer-time has concluded, exit gracefully by pressing `CTRL+C` in the terminal.

## Requests
If you have any game requests, open them as issues on the [main repo](https://github.com/da532/ns-rpc) as they share the same game list.
Otherwise, you're welcome to make a pull request yourself if you would like to see a feature added to this project directly as for me this is a hot-glue solution to real Linux support.
