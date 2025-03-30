package main

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/fatih/color"
	"github.com/itsHardStyl3r/the-swift-codes/cmd/api"
	"github.com/itsHardStyl3r/the-swift-codes/internal/tools"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	color.Green(` ________         ____       _ _____    _____        __      
/_  __/ /  ___   / __/    __(_) _/ /_  / ___/__  ___/ /__ ___
 / / / _ \/ -_) _\ \| |/|/ / / _/ __/ / /__/ _ \/ _  / -_|_-<
/_/ /_//_/\__/ /___/|__,__/_/_/ \__/  \___/\___/\_,_/\__/___/

`)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not load .env file. Make sure that the file is present in project root directory.")
	}
	err = tools.ConnectToDb()
	if err != nil {
		log.Fatalf("Failed to connect to database. %v", err)
	}
	err = tools.SetupDb(true, true)
	if err != nil {
		log.Fatalf("Failed to setup database. %v", err)
	} else {
		log.Info("Database setup complete.")
	}
	tools.LogDatabaseStats()

	log.Info(color.HiCyanString(fmt.Sprintf("Starting HTTP server http://%s...", os.Getenv("httpListenOn"))))
	err = api.Run()
	if err != nil {
		log.Fatalf("Failed to start HTTP server. %v", err)
	}
}
