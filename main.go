package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	loggly "github.com/jamespearly/loggly"
)

//JSON Structs
type featuredGames struct {
	GameList []lolGame `json"gameList"`
}

type lolGame struct {
	GameId       int              `json"gameId"`
	GameMode     string           `json"gameMode"`
	GameType     string           `json"gameType"`
	Participants []lolParticipant `json"participants"`
}

type lolParticipant struct {
	TeamId       int    `json"teamId"`
	ChampionId   int    `json"championId"`
	SummonerName string `json"summonerName"`
}

//Main function
func main() {

	//Loggly Testing
	var tag string
	tag = "My-Go-Demo"
	client := loggly.New(tag)

	err := client.EchoSend("error", "This is a debug memssage")

	//Get request for featured games
	resp, err := http.Get("https://na1.api.riotgames.com/lol/spectator/v4/featured-games?api_key=RGAPI-225637a3-5eb0-4b72-b122-c4654faa3103")
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	var y featuredGames
	err = json.Unmarshal(body, &y)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%+v", y)

}
