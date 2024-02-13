package cmd

import (
	"fmt"

	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/scenes"
	"github.com/spf13/cobra"
)

var sceneList = []string{"Starting", "Secrets", "Coding", "Ending","Gaming", "Talk"}

func changeScene(sceneName string) {
	var err error
	client, err := goobs.New("localhost:4455")
	if err != nil {
		panic(err)
	}
	defer client.Disconnect()
	n, err := client.Scenes.SetCurrentProgramScene(&scenes.SetCurrentProgramSceneParams{
		SceneName: &sceneName,
	})
	fmt.Println(n)
	if err != nil {
		panic(err)
	}
}

func manipulateStream(action string) {
  var err error
  client, err := goobs.New("localhost:4455")
  if err != nil {
    panic(err)
  }
  defer client.Disconnect()
  if action == "start" {
    _, err = client.Stream.StartStream()
  } else if action == "stop" {
    _, err = client.Stream.StopStream()
  }
  if err != nil {
    panic(err)
  }
}

func isStringNotInArray(target string, array []string) bool {
	for _, element := range array {
		if element == target {
			return false
		}
	}
	return true
}

var ScenesCmd = &cobra.Command{
	Use:   "scenes",
	Short: "Set the current scene",
	Long:  `Set the current scene`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a scene name")
			return
		}
		if isStringNotInArray(args[0], sceneList) {
			fmt.Println("Please provide a valid scene name")
			return
		} else {
			changeScene(args[0])
			fmt.Println("Scene changed to " + args[0])
		}
	},
}

var StreamCmd = &cobra.Command{
  Use: "stream",
  Short: "Starts or stops the stream",
  Long: `Starts or stops the stream`,
  Run : func(cmd *cobra.Command, args []string) {
    if len(args) < 1 {
      fmt.Println("Please provide a valid action")
      return
    }
    manipulateStream(args[0])
  },
}


func init() {
	rootCmd.AddCommand(ScenesCmd)
	rootCmd.AddCommand(StreamCmd)
}
