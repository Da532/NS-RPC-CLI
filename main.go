package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/hugolgst/rich-go/client"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Games []struct {
	Title string `json:"title"`
	Img   string `json:"img"`
}

type Game struct {
	Title string `json:"title"`
	Img   string `json:"img"`
}

var gamesList Games

const clientID string = "1114647533562646700"
const gamesURL string = "https://raw.githubusercontent.com/Da532/NS-RPC/master/games.json"

func main() {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nExiting..")
		client.Logout()
		os.Exit(0)
	}()

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Error: Pass in a game title")
		os.Exit(1)
	} else if args[0] == " " || args[0] == "" {
		fmt.Println("Error: Pass in a game title")
		os.Exit(1)
	}
	err := GetGamesData()
	if err != nil {
		panic(err)
	}
	err = client.Login(clientID)
	if err != nil {
		panic(err)
	}
	var gameTitle string
	for _, arg := range args {
		gameTitle += arg + " "
	}
	SetGame(gameTitle)
	for {
	}
}

func GetGamesData() error {
	resp, err := http.Get(gamesURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &gamesList)
	if err != nil {
		return err
	}
	return nil
}

func SetGame(title string) {
	var selectedGame Game
	title = cases.Title(language.English).String(title)
	for _, game := range gamesList {
		if game.Title+" " == title {
			selectedGame = game
			break
		}
	}
	if selectedGame.Title != "" {
		err := client.SetActivity(client.Activity{
			LargeImage: selectedGame.Img,
			Details:    selectedGame.Title,
		})
		if err != nil {
			panic(err)
		}
	} else {
		err := client.SetActivity(client.Activity{
			LargeImage: "home",
			Details:    title,
		})
		if err != nil {
			panic(err)
		}
	}
	fmt.Printf("You are now playing: %s\n", title)
}
