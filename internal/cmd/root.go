package cmd

import (
	"github.com/j3yzz/telespot/internal/config"
	"github.com/j3yzz/telespot/internal/spotify"
	"github.com/spf13/cobra"
	"log"
	"os"
)

const ExitFailure = 1

func Execute() {
	cfg, _ := config.LoadConfig()

	spotify.NewClient(&cfg)

	root := &cobra.Command{
		Use:     "TeleSpot",
		Short:   "Integrate Telegram Music with Spotify playlist.",
		Version: "0.1.1",
	}

	if err := root.Execute(); err != nil {
		log.Fatalln("failed to execute root command")
		os.Exit(ExitFailure)
	}
}
