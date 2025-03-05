package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gotwitch/audio"

	"github.com/gempir/go-twitch-irc/v3"
	htgotts "github.com/hegedustibor/htgo-tts"
	handlers "github.com/hegedustibor/htgo-tts/handlers"
)

var (
	channel = flag.String("channel", "codecadim", "twitch channel to join")
	lang    = flag.String("lang", "fr", "lang code for the voice (fr, en, ...)")
	media   = flag.String("media", "./media", "mp3 folder with files matching commands")
	speak   = flag.Bool("speak", false, "enable message to speech")
)

func init() {
	flag.Parse()
}

func main() {
	// mediaDir is used by htgotts to store mp3 files produced by text to speech
	mediaDir := os.TempDir()

	speech := htgotts.Speech{Folder: mediaDir, Language: *lang, Handler: &handlers.Native{}}

	client := twitch.NewAnonymousClient()

	var sentence string
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		sentence = fmt.Sprintf("%s: %s", message.User.DisplayName, message.Message)
		log.Println(sentence)
		if message.Message[0] == '#' {
			// commands start with a '#'
			cmd := message.Message[1:]
			switch cmd {
			case "clap":
				audio.PlayClap()
				return
			case "bell":
				audio.PlayBell()
				return
			case "speak":
				*speak = !*speak // turn on/off text to speech
				return
			}
			// must match an audio mp3 in 'audio' folder
			snd := *media + "/" + cmd + ".mp3"
			if _, err := os.Stat(snd); err != nil {
				log.Printf("command not found: %s\n", err)
				return
			}
			// play the audio files in go routines so we do not block the bot
			go func() {
				err := audio.Play(snd)
				if err != nil {
					log.Printf("error from play: %s\n", err)
				}
			}()
			return // so we do nothing more
		}
		// if not a command and speak enabled then say the message
		if *speak {
			err := speech.Speak(sentence)
			if err != nil {
				log.Printf("error from speech: %s", err)
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
