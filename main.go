package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var err error
var discord *discordgo.Session

func init() {
	discord, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Fatalf("Invalid bot parameters: %v", err)
	}
}

func main() {
	err := discord.Open()
	if err != nil {
		log.Fatalf("Cannot open the session: %v", err)
	}

	defer discord.Close()

	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Shutting down gracefully")
}
