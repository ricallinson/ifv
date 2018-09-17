package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
)

func main() {
	this := readYamlFileToStory("./fixtures/story.yaml")
	rand.Seed(time.Now().UTC().UnixNano())
	this.OutputBegin()
	this.Loop()
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

func (this *Story) GetItems(locId string) []*Item {
	items := []*Item{}
	for item, loc := range this.ItemsLocations {
		if loc == locId {
			items = append(items, this.GetItem(item))
		}
	}
	return items
}

func (this *Story) GetCurrentLocationItems() []*Item {
	return this.GetItems(this.currentLocation.Id)
}

func (this *Story) OutputBegin() {
	Output(this.Title, true)
	Output(this.Scene, false)
	this.currentLocation = this.GetLocation(this.LocationStart)
}

func (this *Story) OutputScene() {
	Output(this.currentLocation.Discribe(), true)
}

func (this *Story) OutputItems() {
	// Output the options
	items := this.GetCurrentLocationItems()
	if len(items) == 0 {
		Output(randStringSelection(this.SearchFailTitle), true)
		return
	}
	Output(randStringSelection(this.SearchingTitle), true)
	for i, item := range items {
		OutputOption(i+1, item.Discribe())
	}
	Output("", true)
	Output(randStringSelection(this.SearchChooseTitle), true)
	// Wait for input
	var choice int
	fmt.Scan(&choice)
	// Parse input.
	if choice > 0 && choice <= len(items) {
		// Pick up item.
		this.ItemsLocations[items[choice-1].Id] = "player"
	}
}

func (this *Story) OutputOptions() {
	// Output the options
	Output(randStringSelection(this.OptionsTitle), true)
	for i, l := range this.currentLocation.Exits {
		OutputOption(i+1, this.GetLocation(l.Id).Discribe())
	}
	OutputOption(len(this.currentLocation.Exits)+1, randStringSelection(this.SearchOptionsTitle))
	Output("", true)
	Output(randStringSelection(this.OptionsChooseTitle), true)
	// Wait for input
	var choice int
	fmt.Scan(&choice)
	// Parse input.
	if choice > 0 && choice <= len(this.currentLocation.Exits) {
		this.currentLocation = this.GetLocation(this.currentLocation.Exits[choice-1].Id)
	}
	if choice == len(this.currentLocation.Exits)+1 {
		this.OutputItems()
	}
}

func (this *Story) Loop() {
	this.OutputScene()
	this.OutputOptions()
	this.Loop()
}

func OutputOption(i int, s string) {
	fmt.Println(i, ") ", s)
}

func Output(s string, newline bool) {
	fmt.Print(s)
	if newline {
		fmt.Print("\n\n")
	} else {
		fmt.Print(" ")
	}
}

func randStringSelection(s []string) string {
	i := rand.Intn(len(s))
	return s[i]
}

// Reads a file into a byte array or exits.
func readFileToByteArray(p string) []byte {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		log.Fatalf("Error reading file: #%v ", err)
	}
	return b
}

func readYamlFileToStory(p string) *Story {
	r := &Story{}
	if p != "" {
		err := yaml.Unmarshal(readFileToByteArray(p), r)
		if err != nil {
			log.Fatalf("YAML Parse Error: %v", err)
		}
	}
	return r
}

func interfaceToYaml(i interface{}) []byte {
	d, err := yaml.Marshal(i)
	if err != nil {
		log.Fatalf("YAML Write Error: %v", err)
	}
	return d
}
