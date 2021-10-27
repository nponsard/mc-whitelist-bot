package config

import (
	"bufio"
	"encoding/json"
	"os"
	"path"
)

type Rcon struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}
type Discord struct {
	Token    string   `json:"token"`
	Channels []string `json:"channels"`
}

type Config struct {
	Rcons   []Rcon  `json:"rcons"`
	Discord Discord `json:"discord"`
}

var config Config

func LoadConfig(filePath string) (c *Config, err error) {

	// ensure that the folder exists

	_, err = os.Stat(path.Dir(filePath))
	if os.IsNotExist(err) {
		err = os.MkdirAll(path.Dir(filePath), 0770)
		if err != nil {
			return &config, err
		}
	} else if err != nil {

		// file system error, exit

		return &config, err
	}

	// Open file, create if needed

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0660)

	defer file.Close()

	// read file content

	content := ""
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content = content + scanner.Text()
		err = scanner.Err()
		if err != nil {
			return &config, err
		}
	}

	// insert default config if empty

	if len(content) == 0 {
		configJson, err := json.MarshalIndent(config, "", "\t")
		if err != nil {
			return &config, err
		}
		file.Write(configJson)
	} else {
		err = json.Unmarshal([]byte(content), &config)
	}

	return &config, err
}

func GetConfig() *Config {
	return &config
}
