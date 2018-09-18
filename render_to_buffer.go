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

func (this *RenderToBuffer) String(s string) {
	this.stringsOutput = append(this.stringsOutput, s)
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

func (this *RenderToBuffer) ReadIndexedString(i int) string {
	if i < 0 || i >= len(this.stringsOutput) {
		return ""
	}
	return this.stringsOutput[i]
}

func (this *RenderToBuffer) ReadLastString() string {
	return this.ReadIndexedString(len(this.stringsOutput) - 1)
}

func (this *RenderToBuffer) AddInputString(s string) {
	this.stringsInput = append(this.stringsInput, s)
}

func (this *RenderToBuffer) AddIntputInt(i int) {
	this.intsInput = append(this.intsInput, i)
}