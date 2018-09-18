package main

import (
	. "github.com/ricallinson/simplebdd"
	"testing"
)

func TestUtils(t *testing.T) {

	Describe("toggleString()", func() {

		It("should add string to end", func() {
			s := []string{"foo", "bar", "baz"}
			s = toggleString(s, "bob")
			AssertEqual(len(s), 4)
			AssertEqual(s[0], "foo")
			AssertEqual(s[1], "bar")
			AssertEqual(s[2], "baz")
			AssertEqual(s[3], "bob")
		})

		It("should remove string first", func() {
			s := []string{"foo", "bar", "baz"}
			s = toggleString(s, "foo")
			AssertEqual(len(s), 2)
			AssertEqual(s[0], "bar")
			AssertEqual(s[1], "baz")
		})

		It("should remove string middle", func() {
			s := []string{"foo", "bar", "baz"}
			s = toggleString(s, "bar")
			AssertEqual(len(s), 2)
			AssertEqual(s[0], "foo")
			AssertEqual(s[1], "baz")
		})

		It("should remove string last", func() {
			s := []string{"foo", "bar", "baz"}
			s = toggleString(s, "baz")
			AssertEqual(len(s), 2)
			AssertEqual(s[0], "foo")
			AssertEqual(s[1], "bar")
		})

		It("should add string to empty slice", func() {
			s := []string{}
			s = toggleString(s, "foo")
			AssertEqual(len(s), 1)
			AssertEqual(s[0], "foo")
		})
	})

	Report(t)
}
