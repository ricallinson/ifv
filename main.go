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
	story := readYamlFileToStory("./fixtures/story.yaml")
	rand.Seed(time.Now().UTC().UnixNano())
	story.OutputBegin()
	story.Loop()
}

func (this *Story) GetLocation(id string) *Location {
	for _, l := range this.Locations {
		if l.Id == id {
			return l
		}
	}
	return nil
}

func (this *Story) OutputBegin() {
	Output(this.Title, true)
	Output(this.Scene, false)
	this.currentLocation = this.GetLocation(this.LocationStart)
}

func (this *Story) OutputScene() {
	Output(this.currentLocation.Discribe(), true)
}

func (this *Story) OutputOptions() {
	Output(randStringSelection(this.OptionsTitle), true)
	for i, l := range this.currentLocation.Exits {
		OutputOption(i+1, this.GetLocation(l.Id).Discribe())
	}
	Output("", true)
	Output(randStringSelection(this.OptionsChooseTitle), true)
}

func (this *Story) Loop() {
	this.OutputScene()
	this.OutputOptions()
}

func OutputOption(i int, s string) {
	fmt.Print(i, ") ", s)
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
