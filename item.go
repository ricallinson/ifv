package main

import ()

type Item struct {
	Id           string   `yaml:"Id"`
	Scene        string   `yaml:"Scene"`
	Descriptions []string `yaml:"Descriptions"`
	Success      []string `yaml:"Success"`
	Failure      []string `yaml:"Failure"`
	visited      bool
}

func (this *Item) Discribe() string {
	if this.visited == false {
		this.visited = true
		return this.Scene
	}
	return randStringSelection(this.Descriptions)
}
