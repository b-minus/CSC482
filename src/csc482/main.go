package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"

	cron "github.com/robfig/cron"

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

func getRequest() {
	//Loggly Testing
	var tag string
	tag = "cron-test"
	client := loggly.New(tag)

	err := client.EchoSend("error", "This is a debug memssage")

	//Get request for featured games
	var apiKey string
	var url string
	apiKey = "RGAPI-f02697b0-1f2c-41ec-88d2-2fbd8c9fbb7d"
	url = "https://na1.api.riotgames.com/lol/spectator/v4/featured-games?api_key=" + apiKey
	resp, err := http.Get(url)
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

//Main function
func main() {
	//Run code once every 15 minutes
	c := cron.New()
	c.AddFunc("@every 15m", func() { getRequest() })
	c.Start()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}
