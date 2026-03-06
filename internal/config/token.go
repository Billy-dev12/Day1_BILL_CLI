package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	GithubToken string `json:"github_token"`
}

func SaveToken(token string) error {
	conf := Config{
		GithubToken: token,
	}
	data, err := json.MarshalIndent(conf, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile("config.json", data, 0644)
}
