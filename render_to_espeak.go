package main

import (
	"github.com/ricallinson/tts"
	"os"
)

type RenderToEspeak struct {
	tts   *tts.Tts
	input *RenderToConsole
}

func CreateRenderToEspeak() *RenderToEspeak {
	this := &RenderToEspeak{
		tts:   tts.CreateTts(),
		input: CreateRenderToConsole(),
	}
	return this
}

func (this *RenderToEspeak) Section() {
	this.String("\n\n\n")
}

func (this *RenderToEspeak) String(s string) {
	this.tts.Speak(s)
}

func (this *RenderToEspeak) StringLine(s string) {
	this.String(s + "\n\n")
}

func (this *RenderToEspeak) AskForString() string {
	return this.input.AskForString()
}

func (this *RenderToEspeak) AskForInt() int {
	return this.input.AskForInt()
}

func (this *RenderToEspeak) Quit() {
	os.Exit(0)
}
