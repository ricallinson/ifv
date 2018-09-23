package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var storyYamlPath string
	flag.StringVar(&storyYamlPath, "story", "", "The path to the location of the story file to use for the game.")
	var textInterface bool
	flag.BoolVar(&textInterface, "text", false, "Play the game via the console with text.")
	flag.Parse()
	if storyYamlPath == "" {
		fmt.Println("You must supply a path to the location of the story file to use for the game.")
		os.Exit(1)
	}
	if textInterface {
		CreateGame(storyYamlPath, CreateRenderToConsole()).Play()
	} else {
		CreateGame(storyYamlPath, CreateRenderToEspeak()).Play()
	}
}
