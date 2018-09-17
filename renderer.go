package main

type Renderer interface {
	String(s string)
	AskForString() string
	AskForInt() int
}
