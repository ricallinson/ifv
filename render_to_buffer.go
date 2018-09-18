package main

import ()

type RenderToBuffer struct {
	stringsOutput []string
	stringsInput  []string
	stringsIndex  int
	intsInput     []int
	intsIndex     int
}

func CreateRenderToBuffer() *RenderToBuffer {
	return &RenderToBuffer{
		stringsOutput: []string{},
		stringsInput:  []string{},
		intsInput:     []int{},
	}
}

func (this *RenderToBuffer) Section() {
	this.stringsOutput = append(this.stringsOutput, "\n\n")
}

func (this *RenderToBuffer) String(s string) {
	this.stringsOutput = append(this.stringsOutput, s)
}

func (this *RenderToBuffer) StringLine(s string) {
	this.stringsOutput = append(this.stringsOutput, s+"\n")
}

func (this *RenderToBuffer) AskForString() string {
	if this.stringsIndex >= len(this.stringsOutput) {
		return ""
	}
	s := this.stringsInput[this.stringsIndex]
	this.stringsIndex++
	return s
}

func (this *RenderToBuffer) AskForInt() int {
	if this.intsIndex >= len(this.intsInput) {
		return -1
	}
	i := this.intsInput[this.intsIndex]
	this.intsIndex++
	return i
}

func (this *RenderToBuffer) Quit() {
	this.stringsOutput = append(this.stringsOutput, "quit")
}

func (this *RenderToBuffer) ReadIndexedString(i int) string {
	i = len(this.stringsOutput) - i
	if i <= 0 || i >= len(this.stringsOutput) {
		return ""
	}
	return this.stringsOutput[i]
}

func (this *RenderToBuffer) ReadLastStrings(n int) []string {
	s := []string{}
	for ; n > 0; n-- {
		s = append(s, this.ReadIndexedString(n))
	}
	return s
}

func (this *RenderToBuffer) ReadLastString() string {
	return this.ReadIndexedString(1)
}

func (this *RenderToBuffer) AddInputString(s string) {
	this.stringsInput = append(this.stringsInput, s)
}

func (this *RenderToBuffer) AddIntputInt(i int) {
	this.intsInput = append(this.intsInput, i)
}
