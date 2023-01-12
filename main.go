package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gempir/go-twitch-irc/v3"
	htgotts "github.com/hegedustibor/htgo-tts"
	handlers "github.com/hegedustibor/htgo-tts/handlers"
	"gotwitch/audio"
)

var (
	channel = flag.String("channel", "theoldcoder", "the twitch channel to join")
	lang    = flag.String("lang", "fr", "the lang code for the voice (fr, en, ...)")
	media   = flag.String("media", "./media", "folder with mp3 files matching commands")
	speak   = flag.Bool("speak", false, "speak the messages")
)

func main() {

	flag.Parse()

	mediaDir := os.TempDir()

	speech := htgotts.Speech{Folder: mediaDir, Language: *lang, Handler: &handlers.Native{}}

	client := twitch.NewAnonymousClient()

	var sentence string
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if message.Message[0] == '#' {
			// commands start with a '#' and must match an audio mp3 in the 'audio' folder
			cmd := message.Message[1:]
			snd := *media + "/" + cmd + ".mp3"
			if _, err := os.Stat(snd); err != nil {
				fmt.Printf("command not found: %s\n", err)
				return
			}
			// play the audio files in go routines so we do not block the bot
			go func() {
				err := audio.Play(snd)
				if err != nil {
					fmt.Printf("error from play: %s\n", err)
				}
			}()
			return // so we do nothing more
		}
		// if not a command and speak enabled then say the message
		if *speak {
			sentence = fmt.Sprintf("%s %s", message.User.DisplayName, message.Message)
			fmt.Println(sentence)
			err := speech.Speak(sentence)
			if err != nil {
				fmt.Printf("error from speech: %s", err)
			}
			return
		}
		// default is to ring the bell
		audio.Play(*media + "/bell.mp3")
	})

	client.Join(*channel)

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
