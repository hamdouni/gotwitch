package main

import (
	"flag"
	"fmt"

	"github.com/gempir/go-twitch-irc/v3"
	htgotts "github.com/hegedustibor/htgo-tts"
	handlers "github.com/hegedustibor/htgo-tts/handlers"
)

func main() {

	channel := flag.String("channel", "theoldcoder", "the twitch channel to join")
	lang := flag.String("lang", "fr", "the lang code for the voice (fr, en, ...)")

	flag.Parse()

	speech := htgotts.Speech{Folder: "/tmp/audio", Language: *lang, Handler: &handlers.Native{}}

	// for an anonymous user (no write capabilities)
	client := twitch.NewAnonymousClient()
	// client := twitch.NewClient("yourtwitchusername", "oauth:123123123")

	var sentence string
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		sentence = fmt.Sprintf("Message de %s : %s", message.User.DisplayName, message.Message)
		fmt.Println(sentence)
		speech.Speak(sentence)
	})

	client.Join(*channel)

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
