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
	token     string            = os.Getenv("TOKEN")
	channelID string            = os.Getenv("CHANNEL")
	dixionary map[string]string = make(map[string]string)
)

func init() {
	byteFile, _ := ioutil.ReadFile("./dixionary.json")

	rawDixionary := make(map[string]string)

	if err := json.Unmarshal(byteFile, &rawDixionary); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	for englishWord, newWord := range rawDixionary {
		dixionary[strings.ToLower(englishWord)] = newWord
	}
}

func dixionaryGen(message string) string {
	split := strings.Split(message, " ")

	msg := make([]string, 0, len(split))

	for _, word := range split {
		if newWord, ok := dixionary[strings.ToLower(word)]; ok {
			msg = append(msg, newWord)
			continue
		}

		msg = append(msg, word)
	}

	return strings.Join(msg, " ")
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.ChannelID == channelID && len(m.Content) > 0 {
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
