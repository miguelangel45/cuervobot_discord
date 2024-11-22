package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

type Secret struct {
	DiscordToken string `json:"discordToken"`
	AppID        string `json:"appId"`
}

func SetConfig() {
	var secret Secret
	env, err := os.ReadFile("env.json")
	if err != nil {
		fmt.Println(err)
	}

	json.NewDecoder(bytes.NewBuffer(env)).Decode(&secret)
	if err != nil {
		_ = fmt.Errorf("error setting environment: %w", err)
	}

	err = os.Setenv("DISCORD_TOKEN", secret.DiscordToken)
	if err != nil {
		fmt.Errorf("error setting DISCORD_TOKEN: %w", err)
	}
}
