package main

import (
	"fmt"
	"log"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/scenes"
)

func main() {
	client, err := goobs.New("localhost:4444", goobs.WithPassword(""))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect()

	version, _ := client.General.GetVersion()
	fmt.Printf("OBS Studio version: %s\n", version.ObsVersion)
	fmt.Printf("Websocket server version: %s\n", version.ObsWebSocketVersion)

	resp, _ := client.Scenes.GetSceneList()
	for _, v := range resp.Scenes {
		fmt.Printf("%2d %s\n", v.SceneIndex, v.SceneName)
	}

	param := scenes.SetCurrentProgramSceneParams{
		SceneName: "cam full",
	}
	client.Scenes.SetCurrentProgramScene(&param)

}
