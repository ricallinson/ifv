package main

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
)

func TestGame(t *testing.T) {

	var g *Game
	var b *RenderToBuffer

	BeforeEach(func() {
		b = CreateRenderToBuffer()
		g = CreateGame("./fixtures/story.yaml", b)
	})

	AfterEach(func() {

	})

	Describe("Game()", func() {

		It("should return an Game object", func() {
			AssertEqual(reflect.TypeOf(g).String(), "*main.Game")
		})

		It("should output location text", func() {
			g.Play()
			AssertEqual(b.ReadIndexedString(0), "foo")
		})

		It("should output location and menu", func() {
			b.AddIntputInt(1)
			g.Play()
			AssertEqual(b.ReadLastString(), "3 ) Use\n")
		})
	})

	Report(t)
}
