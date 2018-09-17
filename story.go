package main

import ()

type Story struct {
	Title              string            `yaml:"Title"`
	LocationStart      string            `yaml:"LocationStart"`
	Locations          []*Location       `yaml:"Locations"`
	Items              []*Item           `yaml:"Items"`
	ItemsLocations     map[string]string `yaml:"ItemsLocations"`
	OptionsTitle       []string          `yaml:"OptionsTitle"`
	OptionsChooseTitle []string          `yaml:"OptionsChooseTitle"`
	SearchOptionsTitle []string          `yaml:"SearchOptionsTitle"`
	SearchingTitle     []string          `yaml:"SearchingTitle"`
	SearchFailTitle    []string          `yaml:"SearchFailTitle"`
	SearchChooseTitle  []string          `yaml:"SearchChooseTitle"`
	LocationMoveTitle  string            `yaml:"LocationMoveTitle"`
	ItemUseTitle       string            `yaml:"ItemUseTitle"`
	ItemPickupTitle    string            `yaml:"ItemPickupTitle"`
	ItemPutdownTitle   string            `yaml:"ItemPutdownTitle"`
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
