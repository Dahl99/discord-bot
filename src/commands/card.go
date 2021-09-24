package commands

import (
	"discordbot/src/consts"
	"discordbot/src/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

//Sub struct in exactResult struct. It's used to store the imageURIs from scryfall api
type imageURI struct {
	Png string `json:"png"`
}

type cardFaces struct {
	Image imageURI `json:"image_uris"`
}

type prices struct {
	Usd     string `json:"usd"`
	UsdFoil string `json:"usd_foil"`
}

//Struct used to store data from second http.Get()
type fuzzyResult struct {
	Name   string       `json:"name"`
	Image  imageURI     `json:"image_uris"`
	Prices prices       `json:"prices"`
	Faces  [2]cardFaces `json:"card_faces"`
}

func PostCard(cmd []string, m *discordgo.MessageCreate) {
	// Getting card if card name was entered
	if len(cmd) > 1 {
		utils.SendChannelMessage(m, getCard(cmd))
	}
}

//getCard() fetches a card based on which card name used in command
func getCard(n []string) string {

	name := utils.ReplaceSpace(n[1:]) // Replacing the spaces with "_" to avoid url problems

	if len(name) < 3 {
		return "Name needs to have 3 or more letters to search"
	}

	URL := consts.ScryfallBaseURL + name // Sets url for exact card get request

	res, err := http.Get(URL) // Fetching exact card
	if err != nil {           // Checking for errors
		log.Println(http.StatusServiceUnavailable)
		return consts.ScryfallNotAvailable
	}

	// Decoding fuzzyresult from get request
	var card fuzzyResult
	err = json.NewDecoder(res.Body).Decode(&card)
	if err != nil {
		log.Println(err)
		return consts.DecodingFailed
	}

	res.Body.Close() // Closing body to prevent resource leak

	if card.Image.Png == "" && card.Faces[0].Image.Png == "" && card.Faces[1].Image.Png == "" {
		return "Unable to find requested card, avoid ambigous searches!"
	}

	//	Making the returned string
	var result string

	if card.Prices.Usd != "" || card.Prices.UsdFoil != "" {
		result += "\nTCGPlayer price:"
	}

	if card.Prices.Usd != "" {
		result += "\n\tUSD = " + card.Prices.Usd
	}

	if card.Prices.UsdFoil != "" {
		result += "\n\tUSD Foil = " + card.Prices.UsdFoil
	}

	if card.Image.Png != "" {
		result += "\n" + card.Image.Png
	} else {
		result += "\n" + card.Faces[0].Image.Png + "\n" + card.Faces[1].Image.Png
	}

	return result	// Returning url to png version of card
}
