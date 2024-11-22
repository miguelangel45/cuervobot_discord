package pokemon

import (
	"bytes"
	"discordProject/pokemon/structures"
	"encoding/json"
	"github.com/bwmarrin/discordgo"
	"io"
	"net/http"
	"strings"
)

func NewMessage(discord *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discord.State.User.ID {
		return
	}

	switch {
	case strings.Contains(message.Content, "!pokemon"):
		pokemon := strings.Split(message.Content, " ")
		if len(pokemon) > 1 {
			pokemonImage := getPokemon(pokemon[1])
			discord.ChannelMessageSend(message.ChannelID, pokemonImage)
		} else {
			discord.ChannelMessageSend(message.ChannelID, "Pon el nombre del pokemón! Zoroco...")
		}
	}
}

func getPokemon(pokemon string) string {
	var pokemonStructure structures.Pokemon
	pokemonData, _ := http.Get("https://pokeapi.co/api/v2/pokemon/" + pokemon)
	pokemonJson, _ := io.ReadAll(pokemonData.Body)
	pokemonJsonString := string(pokemonJson)

	if pokemonJsonString == "Not Found" {
		return "No se encontró el pokemon"
	}
	json.NewDecoder(bytes.NewBuffer(pokemonJson)).Decode(&pokemonStructure)
	return pokemonStructure.Sprites.Front_default
}
