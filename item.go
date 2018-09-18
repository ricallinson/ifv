package main

import ()

type Item struct {
	Id             string   `yaml:"Id"`
	Scene          string   `yaml:"Scene"`
	Descriptions   []string `yaml:"Descriptions"`
	OptionsUse     []string `yaml:"OptionsUse"`
	OptionsPickup  []string `yaml:"OptionsPickup"`
	OptionsPutdown []string `yaml:"OptionsPutdown"`
	Success        []string `yaml:"Success"`
	Failure        []string `yaml:"Failure"`
	visited        bool
}

func (this *Item) Discribe() string {
	if this.visited == false {
		this.visited = true
		return this.Scene
	}
	return randStringSelection(this.Descriptions)
}

func (this *Item) DiscribeOptionsUse() string {
	return randStringSelection(this.OptionsUse)
}

func (this *Item) DiscribeOptionsPickup() string {
	return randStringSelection(this.OptionsPickup)
}

func (this *Item) DiscribeOptionsPutdown() string {
	return randStringSelection(this.OptionsPutdown)
}
