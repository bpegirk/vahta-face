package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var configuration = Configuration{}

type FrameConfig struct {
	HardwareId int `json:"hardware_id"`
}
type SocketConfig struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type DbConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}
type Configuration struct {
	Socket SocketConfig  `json:"socket"`
	Db     DbConfig      `json:"db"`
	Frames []FrameConfig `json:"frames"`
}

func initConfig() {
	filename := "config.json"
	jsonFile, err := os.Open(filename)
	defer jsonFile.Close()
	if err != nil {
		fmt.Printf("failed to open json file: %s, error: %v", filename, err)
		return
	}

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Printf("failed to read json file, error: %v", err)
		return
	}
	err = json.Unmarshal(jsonData, &configuration)
	if err != nil {
		fmt.Println("Error decode config")
		panic(err)
	}
}
