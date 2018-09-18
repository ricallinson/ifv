package main

import ()

type Story struct {
	Title           string            `yaml:"Title"`
	LocationStart   []string          `yaml:"LocationStart"`
	LocationExitEnd []string          `yaml:"LocationExitEnd"`
	Locations       []*Location       `yaml:"Locations"`
	Items           []*Item           `yaml:"Items"`
	ItemsLocations  map[string]string `yaml:"ItemsLocations"`
	QuitTitle       []string          `yaml:"QuitTitle"`
}

func (this *Story) GetStartLocation() *Location {
	return this.GetLocation(randStringSelection(this.LocationStart))
}

func (this *Story) IsLocationExitEnd(id string) bool {
	return containsString([]string{id}, this.LocationExitEnd)
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

func (this *Story) DiscribeQuit() string {
	return randStringSelection(this.QuitTitle)
}
