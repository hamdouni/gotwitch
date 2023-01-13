package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/scenes"
)

var (
	scene = flag.String("s", "default", "OBS scene name")
	list  = flag.Bool("l", false, "liste available OBS scenes")
	vers  = flag.Bool("v", false, "print version")
)

func main() {

	flag.Parse()

	client, err := goobs.New("localhost:4444", goobs.WithPassword(""))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect()

	if *vers {
		version, _ := client.General.GetVersion()
		fmt.Printf("OBS Studio version: %s\n", version.ObsVersion)
		fmt.Printf("Websocket server version: %s\n", version.ObsWebSocketVersion)
	}

	if *list {
		resp, _ := client.Scenes.GetSceneList()
		for _, v := range resp.Scenes {
			fmt.Printf("%2d %s\n", v.SceneIndex, v.SceneName)
		}
	}

	param := scenes.SetCurrentProgramSceneParams{
		SceneName: *scene,
	}
	client.Scenes.SetCurrentProgramScene(&param)
}
