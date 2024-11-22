package main

import (
	"DiscordProject/Discord"
	"DiscordProject/config"
	"os"
	"os/signal"
)

func main() {
	config.SetConfig()
	Discord.InitDiscord()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
