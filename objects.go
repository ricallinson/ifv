package main

import (
	"bufio"
)

type LocationExit struct {
	Id           string   `yaml:"Id"`
	Scene        string   `yaml:"Scene"`
	Descriptions []string `yaml:"Descriptions"`
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

type Item struct {
	Id           string   `yaml:"Id"`
	Scene        string   `yaml:"Scene"`
	Descriptions []string `yaml:"Descriptions"`
	Success      []string `yaml:"Success"`
	Failure      []string `yaml:"Failure"`
}

type Story struct {
	Title              string            `yaml:"Title"`
	Scene              string            `yaml:"Scene"`
	LocationStart      string            `yaml:"LocationStart"`
	Locations          []*Location       `yaml:"Locations"`
	Items              []*Item           `yaml:"Items"`
	ItemsLocations     map[string]string `yaml:"ItemsLocations"`
	OptionsTitle       []string          `yaml:"OptionsTitle"`
	OptionsChooseTitle []string          `yaml:"OptionsChooseTitle"`
	currentLocation    *Location
	input              *bufio.Reader
}
