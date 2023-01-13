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
	channel = flag.String("channel", "theoldcoder", "twitch channel to join")
	lang    = flag.String("lang", "fr", "lang code for the voice (fr, en, ...)")
	media   = flag.String("media", "./media", "mp3 folder with files matching commands")
	speak   = flag.Bool("speak", false, "enable message to speech")
)

func main() {

	flag.Parse()

	// mediaDir is used by htgotts to store mp3 files produced by text to speech
	mediaDir := os.TempDir()

	speech := htgotts.Speech{Folder: mediaDir, Language: *lang, Handler: &handlers.Native{}}

	client := twitch.NewAnonymousClient()

	var sentence string
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if message.Message[0] == '#' {
			// commands start with a '#' and must match an audio mp3 in the 'audio' folder
			cmd := message.Message[1:]
			switch cmd {
			case "clap":
				audio.PlayClap()
				return
			case "bell":
				audio.PlayBell()
				return
			}
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
		audio.PlayBell()
	})

	client.Join(*channel)

	err := client.Connect()
	if err != nil {
		panic(err)
	}
}
