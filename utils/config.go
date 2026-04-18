package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
)

const ConfigFile string = "/gitstats.json"
const ConfigJson string = "/.gitstatsConfig"

type Config struct {
	Email string   `json:"email"`
	Repos []string `json:"repos"`
}

func LoadConfig() (*Config, error) {
	file := openFile(GetDotFilePath())

	defer file.Close()

	configBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var config Config
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &config, nil
}

func SaveConfig(config *Config) error {
	configBytes, err := json.Marshal(config)

	if err != nil {
		return err
	}

	err = os.WriteFile(GetConfigFilePath(), configBytes, 0644)

	if err != nil {
		return err
	}

	return nil
}

func getFilePathFromHome(file string) string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dotFile := usr.HomeDir + file
	return dotFile
}

func GetConfigFilePath() string {
	return getFilePathFromHome(ConfigJson)
}

func GetDotFilePath() string {
	return getFilePathFromHome(ConfigFile)
}
