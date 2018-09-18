package main

import (
	"fmt"
)

type RenderToConsole struct {
}

func (this *RenderToConsole) String(s string) {
	fmt.Print(s, " ")
}

func (this *RenderToConsole) AskForString() string {
	// > Bob
	var s string
	fmt.Scan(&s)
	return s
}

func (this *RenderToConsole) AskForInt() int {
	// > 28
	var i int
	fmt.Scan(&i)
	return i
}
