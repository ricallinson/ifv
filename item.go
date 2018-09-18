package main

import (
	"fmt"
)

type Conditions struct {
	Location      []string `yaml:"Location"`
	LocationItems []string `yaml:"LocationItems"`
	UserItems     []string `yaml:"UserItems"`
}

func (this *Conditions) Test(loc *Location, locItems []*Item, userItems []*Item) bool {
	test := 0
	if containsString(this.Location, []string{loc.Id}) {
		test++
	}
	if containsItem(this.LocationItems, locItems) {
		test++
	}
	if containsItem(this.UserItems, userItems) {
		test++
	}
	fmt.Println(test)
	if test < 3 {
		return false
	}
	return true
}

type Item struct {
	Id             string      `yaml:"Id"`
	Scene          string      `yaml:"Scene"`
	Descriptions   []string    `yaml:"Descriptions"`
	OptionsUse     []string    `yaml:"OptionsUse"`
	OptionsPickup  []string    `yaml:"OptionsPickup"`
	OptionsPutdown []string    `yaml:"OptionsPutdown"`
	Success        []string    `yaml:"Success"`
	Failure        []string    `yaml:"Failure"`
	Conditions     *Conditions `yaml:"Conditions"`
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

func (this *Item) DiscribeSuccess() string {
	return randStringSelection(this.Success)
}

func (this *Item) DiscribeFailure() string {
	return randStringSelection(this.Failure)
}
