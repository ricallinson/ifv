package main

import ()

type LocationExit struct {
	Id           string   `yaml:"Id"`
	Scene        string   `yaml:"Scene"`
	Descriptions []string `yaml:"Descriptions"`
	visited      bool
}

func (this *LocationExit) Discribe() string {
	if this.visited == false {
		this.visited = true
		return this.Scene
	}
	return randStringSelection(this.Descriptions)
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
