package main

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v3"
	htgotts "github.com/hegedustibor/htgo-tts"
	handlers "github.com/hegedustibor/htgo-tts/handlers"
	voices "github.com/hegedustibor/htgo-tts/voices"
)

func main() {
	speech := htgotts.Speech{Folder: "/tmp/audio", Language: voices.French, Handler: &handlers.Native{}}

	// for an anonymous user (no write capabilities)
	client := twitch.NewAnonymousClient()
	// client := twitch.NewClient("yourtwitchusername", "oauth:123123123")

	var sentence string
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		sentence = fmt.Sprintf("Message de %s : %s", message.User.DisplayName, message.Message)
		fmt.Println(sentence)
		speech.Speak(sentence)
	})

	client.Join("theoldcoder")

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
