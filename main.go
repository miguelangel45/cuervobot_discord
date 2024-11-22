package main

import (
	"DiscordProject/config"
	"DiscordProject/discord"
	"os"
	"os/signal"
)

func main() {
	config.SetConfig()
	discord.InitDiscord()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
