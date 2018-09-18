package main

import (
	. "github.com/ricallinson/simplebdd"
	"reflect"
	"testing"
	//"fmt" // fmt.Println(b.ReadLastStrings(6))
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

		It("should quit game", func() {
			b.AddIntputInt(0)
			g.Play()
			AssertEqual(b.ReadLastString(), "quit")
		})

		It("should play game to end", func() {
			b.AddIntputInt(2) // Pickup hammer.
			b.AddIntputInt(1) // Move to kitchen.
			b.AddIntputInt(3) // Use hammer.
			b.AddIntputInt(2) // Move to stairs.
			g.Play()
			AssertEqual(b.ReadLastString(), "quit")
		})

		It("should move to kitchen", func() {
			b.AddIntputInt(1) // Move to kitchen.
			b.AddIntputInt(0) // Quit.
			g.Play()
			AssertEqual(g.userloc.Id, "kitchen")
		})

		It("should move to kitchen then back to bed", func() {
			b.AddIntputInt(1) // Move to kitchen.
			b.AddIntputInt(1) // Move to bed.
			b.AddIntputInt(0) // Quit.
			g.Play()
			AssertEqual(g.userloc.Id, "bed")
		})

		It("should pickup the hammer", func() {
			b.AddIntputInt(2) // Pickup hammer.
			b.AddIntputInt(0) // Quit.
			g.Play()
			AssertEqual(g.itemloc["hammer"], "")
		})

		It("should move hammer to the kitchen", func() {
			b.AddIntputInt(2) // Pickup hammer.
			b.AddIntputInt(1) // Move to kitchen.
			b.AddIntputInt(4) // Put down hammer.
			b.AddIntputInt(0) // Quit.
			g.Play()
			AssertEqual(g.itemloc["hammer"], "kitchen")
		})

		It("should move cup to bed and use hammer in the kitchen", func() {
			b.AddIntputInt(1) // Move to kitchen.
			b.AddIntputInt(2) // Pickup cup.
			b.AddIntputInt(1) // Move to bed.
			b.AddIntputInt(4) // Put down cup.
			b.AddIntputInt(2) // Pickup hammer.
			b.AddIntputInt(1) // Move to kitchen.
			b.AddIntputInt(2) // Use hammer.
			b.AddIntputInt(0) // Quit.
			g.Play()
			AssertEqual(containsString(g.exitsHidden, "stairs"), true)
		})
	})

	Report(t)
}
