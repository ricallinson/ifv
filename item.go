package main

import ()

type Result struct {
	Exits []string `yaml:"Exits"`
	Items []string `yaml:"Items"`
}

type Conditions struct {
	Location      []string `yaml:"Location"`
	LocationItems []string `yaml:"LocationItems"`
	UserItems     []string `yaml:"UserItems"`
}

type Item struct {
	Id             string      `yaml:"Id"`
	Hidden         bool        `yaml:"Hidden"`
	Scene          string      `yaml:"Scene"`
	Descriptions   []string    `yaml:"Descriptions"`
	OptionsUse     []string    `yaml:"OptionsUse"`
	OptionsPickup  []string    `yaml:"OptionsPickup"`
	OptionsPutdown []string    `yaml:"OptionsPutdown"`
	Success        []string    `yaml:"Success"`
	Failure        []string    `yaml:"Failure"`
	Conditions     *Conditions `yaml:"Conditions"`
	Result         *Result     `yaml:"Result"`
	visited        bool
}

func (this *Item) Discribe() string {
	if this.visited == false {
		this.visited = true
		return this.Scene
	}
	return randStringSelection(this.Descriptions)
}

func (this *Item) Test(loc *Location, locItems []*Item, userItems []*Item) bool {
	if this.Conditions == nil {
		return false
	}
	test := 0
	if containsString(this.Conditions.Location, loc.Id) {
		test++
	}
	if containsItem(locItems, this.Conditions.LocationItems) {
		test++
	}
	if containsItem(userItems, this.Conditions.UserItems) {
		test++
	}
	if test < 3 {
		return false
	}
	return true
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
