package main

import ()

type Story struct {
	Title          string            `yaml:"Title"`
	Scene          string            `yaml:"Scene"`
	LocationStart  []string          `yaml:"LocationStart"`
	LocationEnd    []string          `yaml:"LocationEnd"`
	Locations      []*Location       `yaml:"Locations"`
	Items          []*Item           `yaml:"Items"`
	ItemsLocations map[string]string `yaml:"ItemsLocations"`
	QuitTitle      []string          `yaml:"QuitTitle"`
}

func CreateStory(p string) *Story {
	this := readYamlFileToStory(p)
	return this
}

func (this *Story) GetStartLocation() *Location {
	return this.GetLocation(randStringSelection(this.LocationStart))
}

func (this *Story) IsLocationTheEnd(locId string) bool {
	return containsString(this.LocationEnd, locId)
}

func (this *Story) GetItemLocations() map[string]string {
	locs := map[string]string{}
	for k, v := range this.ItemsLocations {
		locs[k] = v
	}
	return locs
}

func (this *Story) GetHiddenItems() []string {
	hidden := []string{}
	for _, item := range this.Items {
		if item.Hidden {
			hidden = append(hidden, item.Id)
		}
	}
	return hidden
}

func (this *Story) GetHiddenExits() []string {
	hidden := []string{}
	for _, loc := range this.Locations {
		for _, exit := range loc.Exits {
			if exit.Hidden {
				hidden = append(hidden, exit.Id)
			}
		}
	}
	return hidden
}

func (this *Story) GetLocation(id string) *Location {
	for _, l := range this.Locations {
		if l.Id == id {
			return l
		}
	}
	return nil
}

func (this *Story) GetItem(id string) *Item {
	for _, i := range this.Items {
		if i.Id == id {
			return i
		}
	}
	return nil
}

func (this *Story) Discribe() string {
	if len(this.Scene) == 0 {
		return "Story begins."
	}
	return this.Scene
}

func (this *Story) DiscribeQuit() string {
	if len(this.QuitTitle) == 0 {
		return "Quit game."
	}
	return randStringSelection(this.QuitTitle)
}
