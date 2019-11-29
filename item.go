package main

import ()

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
	if len(this.OptionsUse) == 0 {
		return "Use item " + this.Id + "."
	}
	return randStringSelection(this.OptionsUse)
}

func (this *Item) DiscribeOptionsPickup() string {
	if len(this.OptionsPickup) == 0 {
		return "Pickup item " + this.Id + "."
	}
	return randStringSelection(this.OptionsPickup)
}

func (this *Item) DiscribeOptionsPutdown() string {
	if len(this.OptionsPutdown) == 0 {
		return "Put down item " + this.Id + "."
	}
	return randStringSelection(this.OptionsPutdown)
}

func (this *Item) DiscribeSuccess() string {
	if len(this.Success) == 0 {
		return "Success using item " + this.Id + "."
	}
	return randStringSelection(this.Success)
}

func (this *Item) DiscribeFailure() string {
	if len(this.Failure) == 0 {
		return "Failure using item " + this.Id + "."
	}
	return randStringSelection(this.Failure)
}

func (this *Item) shortDiscribe() string {
	return "Item " + this.Id + "."
}
