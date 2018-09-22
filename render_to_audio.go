package main

import (
	"github.com/ricallinson/tts"
	"os"
)

type RenderToAudio struct {
	tts   *tts.Tts
	input *RenderToConsole
}

func CreateRenderToAudio() *RenderToAudio {
	this := &RenderToAudio{
		tts:   tts.CreateTts(),
		input: CreateRenderToConsole(),
	}
	return this
}

func (this *RenderToAudio) Section() {
	this.String("\n\n\n")
}

func (this *RenderToAudio) String(s string) {
	this.tts.Speak(s)
}

func (this *RenderToAudio) StringLine(s string) {
	this.String(s + "\n\n")
}

func (this *RenderToAudio) AskForString() string {
	return this.input.AskForString()
}

func (this *RenderToAudio) AskForInt() int {
	return this.input.AskForInt()
}

func (this *RenderToAudio) Quit() {
	os.Exit(0)
}
