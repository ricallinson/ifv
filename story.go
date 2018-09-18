package main

import ()

type Story struct {
	Title          string            `yaml:"Title"`
	LocationStart  string            `yaml:"LocationStart"`
	Locations      []*Location       `yaml:"Locations"`
	Items          []*Item           `yaml:"Items"`
	ItemsLocations map[string]string `yaml:"ItemsLocations"`
	QuitTitle      []string          `yaml:"QuitTitle"`
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
