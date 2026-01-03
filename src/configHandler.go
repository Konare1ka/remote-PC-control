package main

import (
	"encoding/json"
	"log/slog"
	"os"
	"path/filepath"
)

type Config struct {
    Token        string `json:"token"`
	Username		 string `json:"username"`
	AllowPlugList []string `json:"allowPlugins"`
}

var cfg *Config
var logger *slog.Logger

func init() { logger = slog.Default() }

func loadJSON() {

	// get path to config
	execPath, err := os.Executable()
	if err != nil {logger.Error("Failed get a execPath", "error", err.Error()); os.Exit(1)}
	jsonConfig := filepath.Join(filepath.Dir(execPath), "config.json")

	//parse json
	data, err := os.ReadFile(jsonConfig)
	if err != nil {logger.Error("Failed read config.json", "error", err.Error()); os.Exit(1)}
	err = json.Unmarshal(data, &cfg)
	if err != nil {logger.Error("Failed parse JSON", "error", err.Error()); os.Exit(1)}
	
	//json field checking
	if cfg.Token == "" {
		logger.Error("In remote-PC-control/config.json the \"token\" field is empty")
		os.Exit(1)
	} else if cfg.Username == "" {
		logger.Error("In remote-PC-control/config.json the \"username\" field is empty")
		os.Exit(1)
	} else { logger.Info("Successful load config") }
}