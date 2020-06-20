package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

var (
	token     string = os.Getenv("TOKEN")
	channelID string = os.Getenv("CHANNEL")
	dixionary map[string]string
)

func init() {
	byteFile, _ := ioutil.ReadFile("./dixionary.json")

	if err := json.Unmarshal(byteFile, &dixionary); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

func dixionaryGen(message string) string {
	var msg []string

	split := strings.Split(message, " ")
	for _, word := range split {
		newWord := word
		for englishWord, scamWord := range dixionary {
			if strings.ToLower(englishWord) == strings.ToLower(word) {
				newWord = scamWord
				break
			}
		}

		msg = append(msg, newWord)
	}

	return strings.Join(msg, " ")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.ChannelID == channelID {
		s.ChannelMessageSend(channelID, dixionaryGen(m.Content))
	}
}

func main() {
	dg, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Printf("Error creating discord session: %v \n", err)
		return
	}

	dg.AddHandler(messageCreate)

	if err = dg.Open(); err != nil {
		fmt.Printf("Error opening connection: %v \n", err)
		return
	}

	fmt.Println("Sadly dixionary bot is running again...")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	dg.Close()
}
