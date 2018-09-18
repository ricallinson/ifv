package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"math/rand"
)

func randStringSelection(s []string) string {
	i := rand.Intn(len(s))
	return s[i]
}

func containsString(a []string, b []string) bool {
	for _, av := range a {
		for _, bv := range b {
			if av == bv {
				return true
			}
		}
	}
	return false
}

func containsItem(a []*Item, b []string) bool {
	for _, av := range a {
		for _, bv := range b {
			if av.Id == bv {
				return true
			}
		}
	}
	return false
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
