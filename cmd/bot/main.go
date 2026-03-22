package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"ludo-bot/internal/bot"
	"ludo-bot/internal/config"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	discordBot, err := bot.New(cfg, "configs/games")
	if err != nil {
		log.Fatal("Failed to create bot:", err)
	}

	if err := discordBot.Start(); err != nil {
		log.Fatal("Failed to start bot:", err)
	}

	log.Println("Ludo bot is running. Press Ctrl+C to exit.")

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop

	log.Println("Shutting down bot...")
	discordBot.Stop()
}
