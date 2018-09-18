package main

type Renderer interface {
	Section()
	String(s string)
	StringLine(s string)
	AskForString() string
	AskForInt() int
	Quit()
}
