package main

import (
	"fmt"
	"os"
)

type RenderToConsole struct {
}

func (this *RenderToConsole) Section() {
	fmt.Print("\n\n")
}

func (this *RenderToConsole) String(s string) {
	fmt.Print(s, " ")
}

func (this *RenderToConsole) StringLine(s string) {
	fmt.Print(s, "\n")
}

func (this *RenderToConsole) AskForString() string {
	var s string
	fmt.Scan(&s)
	return s
}

func (this *RenderToConsole) AskForInt() int {
	var i int
	fmt.Scan(&i)
	return i
}

func (this *RenderToConsole) Quit() {
	fmt.Print("\n\n")
	os.Exit(0)
}
