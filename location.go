package main

import ()

type LocationExit struct {
	Id           string   `yaml:"Id"`
	Hidden       bool     `yaml:"Hidden"`
	Scene        string   `yaml:"Scene"`
	Descriptions []string `yaml:"Descriptions"`
	Options      []string `yaml:"Options"`
	visited      bool
}

func (this *LocationExit) Discribe() string {
	if this.visited == false {
		this.visited = true
		if len(this.Scene) == 0 {
			return this.shortDiscribe()
		}
		return this.Scene
	}
	if len(this.Descriptions) == 0 {
		return this.shortDiscribe()
	}
	return randStringSelection(this.Descriptions)
}

func (this *LocationExit) DiscribeOptions() string {
	if len(this.Options) == 0 {
		return "Exit to location " + this.Id + "."
	}
	return randStringSelection(this.Options)
}

func (this *LocationExit) shortDiscribe() string {
	return "Location " + this.Id + "."
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
		if len(this.Scene) == 0 {
			return this.shortDiscribe()
		}
		return this.Scene
	}
	if len(this.Descriptions) == 0 {
		return this.shortDiscribe()
	}
	return randStringSelection(this.Descriptions)
}

func (this *Location) shortDiscribe() string {
	return "Location " + this.Id + "."
}
