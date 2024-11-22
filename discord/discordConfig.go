package discord

import (
	"DiscordProject/pokemon"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
)

func InitDiscord() {
	discord, _ := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	discord.SyncEvents = true
	discord.AddHandler(pokemon.NewMessage)
	err := discord.Open()
	fmt.Printf("Bot running.... \n")
	if err != nil {
		fmt.Errorf("error opening connection to discord: %v", err)
	}
	defer func(discord *discordgo.Session) {
		err := discord.Close()
		if err != nil {
			fmt.Errorf("error closing connection to discord: %v", err)
		}
	}(discord)
}
