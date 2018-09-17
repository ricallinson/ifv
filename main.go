package main

import ()

func main() {
	CreateGame("./fixtures/story.yaml", &RenderToConsole{}).Play()
}
