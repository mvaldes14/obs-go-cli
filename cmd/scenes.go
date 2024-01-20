package cmd

import (
	"fmt"
	"github.com/andreykaipov/goobs"
	"github.com/andreykaipov/goobs/api/requests/scenes"
	"github.com/spf13/cobra"
)

func ChangeScene(sceneName string) {
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

var scenesCmd = &cobra.Command{
	Use:   "scenes",
	Short: "Set the current scene",
	Long:  `Set the current scene`,
	Run: func(cmd *cobra.Command, args []string) {
		ChangeScene(args[0])
		fmt.Println("Scene changed to " + args[0])
	},
}

func init() {
	rootCmd.AddCommand(scenesCmd)
}
