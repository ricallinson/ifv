package main

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestStory(t *testing.T) {

	var s *Story

	BeforeEach(func() {
		s = CreateStory("./fixtures/story.yaml")
	})

	AfterEach(func() {

	})

	Describe("CreateStory()", func() {

		It("should return an Story object", func() {
			AssertEqual(reflect.TypeOf(s).String(), "*main.Story")
		})
	})

	Describe("GetStartLocation()", func() {

		It("should return bed", func() {
			AssertEqual(s.GetStartLocation().Id, "bed")
		})

	})

	Describe("IsLocationTheEnd()", func() {

		It("should return true as stairs is the end", func() {
			AssertEqual(s.IsLocationTheEnd("stairs"), true)
		})

		It("should return false as bed is NOT the end", func() {
			AssertEqual(s.IsLocationTheEnd("bed"), false)
		})

	})

	Report(t)
}
