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

func toggleString(a []string, toggle string) []string {
	for i, v := range a {
		if v == toggle {
			return append(a[:i], a[i+1:]...)
		}
	}
	return append(a, toggle)
}

func toggleStrings(a []string, toggle []string) []string {
	for _, v := range toggle {
		a = toggleString(a, v)
	}
	return a
}

func containsString(a []string, str string) bool {
	for _, v := range a {
		if v == str {
			return true
		}
	}
	return false
}

func containsStrings(a []string, str []string) bool {
	for _, v := range a {
		if containsString(str, v) {
			return true
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
