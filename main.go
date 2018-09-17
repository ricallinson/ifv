package main

import (
	"bufio"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	this := readYamlFileToStory("./fixtures/story.yaml")
	rand.Seed(time.Now().UTC().UnixNano())
	this.input = bufio.NewReader(os.Stdin)
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
	input, _ := this.input.ReadString('\n')
	choice, _ := strconv.Atoi(strings.TrimSpace(input))
	if choice > 0 && choice <= len(this.currentLocation.Exits) {
		this.currentLocation = this.GetLocation(this.currentLocation.Exits[choice-1].Id)
	}
}

func (this *Story) Loop() {
	this.OutputScene()
	this.OutputOptions()
	this.Loop()
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
