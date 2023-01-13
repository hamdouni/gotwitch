package audio

import (
	"embed"
	"io"
	"os"
	"time"

	"github.com/hajimehoshi/go-mp3"
	"github.com/hajimehoshi/oto/v2"
)

func Play(filename string) error {
	// Read the mp3 file as a stream
	sndFile, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer sndFile.Close()
	return p(sndFile)
}

//go:embed assets/bell.mp3
var bellFile embed.FS

//go:embed assets/clap.mp3
var clapFile embed.FS

func PlayBell() error {
	bellReader, _ := bellFile.Open("assets/bell.mp3")
	return p(bellReader)
}

func PlayClap() error {
	clapReader, _ := clapFile.Open("assets/clap.mp3")
	return p(clapReader)
}

func p(sndFile io.Reader) error {
	// Decode file
	decodedMp3, err := mp3.NewDecoder(sndFile)
	if err != nil {
		return err
	}

	// Prepare an Oto context (this will use your default audio device) that will
	// play all our sounds. Its configuration can't be changed later.

	// Usually 44100 or 48000. Other values might cause distortions in Oto
	samplingRate := 48000

	// Number of channels (aka locations) to play sounds from. Either 1 or 2.
	// 1 is mono sound, and 2 is stereo (most speakers are stereo).
	numOfChannels := 2

	// Bytes used by a channel to represent one sample. Either 1 or 2 (usually 2).
	audioBitDepth := 2

	// Remember that you should **not** create more than one context
	otoCtx, readyChan, err := oto.NewContext(samplingRate, numOfChannels, audioBitDepth)
	if err != nil {
		return err
	}
	// It might take a bit for the hardware audio devices to be ready, so we wait on the channel.
	<-readyChan

	// Create a new 'player' that will handle our sound. Paused by default.
	player := otoCtx.NewPlayer(decodedMp3)

	// Play starts playing the sound and returns without waiting for it (Play() is async).
	player.Play()

	// We can wait for the sound to finish playing using something like this
	for player.IsPlaying() {
		time.Sleep(time.Second)
	}

	// If you don't want the player/sound anymore simply close
	err = player.Close()
	if err != nil {
		return err
	}
	return nil
}
