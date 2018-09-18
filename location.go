package main

import ()

type LocationExit struct {
	Id           string   `yaml:"Id"`
	Scene        string   `yaml:"Scene"`
	Descriptions []string `yaml:"Descriptions"`
	Options      []string `yaml:"Options"`
	visited      bool
}

func (this *LocationExit) Discribe() string {
	if this.visited == false {
		this.visited = true
		return this.Scene
	}
	return randStringSelection(this.Descriptions)
}

func (this *LocationExit) DiscribeOptions() string {
	return randStringSelection(this.Options)
}

type Location struct {
	Id           string          `yaml:"Id"`
	Scene        string          `yaml:"Scene"`
	Descriptions []string        `yaml:"Descriptions"`
	Exits        []*LocationExit `yaml:"Exits"`
	visited      bool
}

func (this *Location) Discribe() string {
	if this.visited == false {
		this.visited = true
		return this.Scene
	}
	return randStringSelection(this.Descriptions)
}
