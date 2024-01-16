package main

import (
	"fmt"

	"github.com/andreykaipov/goobs"
  "github.com/andreykaipov/goobs/api/requests/scenes"
)

func main() {
  var err error

	client, err := goobs.New("localhost:4455")
	if err != nil {
		panic(err)
	}
	defer client.Disconnect()

  myscene := "Secrets"
	n, err := client.Scenes.SetCurrentProgramScene(&scenes.SetCurrentProgramSceneParams{
		SceneName: &myscene,
	})
  fmt.Println(n)
  if err != nil {
    panic(err)
  }
}
