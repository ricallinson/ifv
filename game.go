package main

import (
	"fmt"
)

type Game struct {
	story   *Story
	render  Renderer
	userloc *Location
	itemloc map[string]string
}

func CreateGame(p string, r Renderer) *Game {
	this := &Game{
		story:   readYamlFileToStory(p),
		render:  r,
		itemloc: map[string]string{},
	}
	this.userloc = this.story.GetLocation(this.story.LocationStart)
	for k, v := range this.story.ItemsLocations {
		this.itemloc[k] = v
	}
	return this
}

func (this *Game) Play() {
	this.discribeLocation()
	this.menu()
	this.Play()
}

func (this *Game) getUserItems() []*Item {
	items := []*Item{}
	for itemId, loc := range this.itemloc {
		if loc == "" {
			items = append(items, this.story.GetItem(itemId))
		}
	}
	return items
}

func (this *Game) getLocationItems() []*Item {
	items := []*Item{}
	for itemId, loc := range this.itemloc {
		if loc == this.userloc.Id {
			items = append(items, this.story.GetItem(itemId))
		}
	}
	return items
}

func (this *Game) getLocationExits() []*LocationExit {
	return this.userloc.Exits
}

func (this *Game) discribeItem(id string) {
	// On the floor is an old hammer.
}

func (this *Game) discribeLocationItems() {
	// On the floor is an old hammer. A cups sits alone.
}

func (this *Game) discribeLocation() {
	// You are in a room. On the floor is an old hammer. A cups sits alone.
	this.render.String(this.userloc.Discribe())
}

func (this *Game) menu() {
	exits := this.getLocationExits()
	locItems := this.getLocationItems()
	userItems := this.getUserItems()
	options := map[string]func(){}
	if len(exits) > 0 {
		options[this.story.LocationMoveTitle] = this.move
	}
	if len(locItems) > 0 {
		options[this.story.ItemPickupTitle] = this.pickup
	}
	if len(userItems) > 0 {
		options[this.story.ItemPutdownTitle] = this.putdown
		options[this.story.ItemUseTitle] = this.use
	}
	i := 1
	this.render.String("\n")
	for s, _ := range options {
		this.render.String(fmt.Sprintf("\n%d ) %s\n", i, s))
		i++
	}
	this.render.String("\n:")
	choice := this.render.AskForInt()
	for _, fn := range options {
		choice--
		if choice == 0 {
			fn()
		}
	}
}

func (this *Game) move() {
	exits := this.getLocationExits()
	for i, s := range exits {
		this.render.String(fmt.Sprintf("\n%d ) %s\n", i+1, s.Id))
	}
	this.render.String("\n:")
	choice := this.render.AskForInt() - 1
	if choice < len(exits) {
		this.userloc = this.story.GetLocation(exits[choice].Id)
	}
}

func (this *Game) use() {
	// 1) Dog
	// 2) Cup
	// 3) Cancel
	this.render.String("\n\nUse")
}

func (this *Game) pickup() {
	// 1) Hammer
	// 2) Cup
	// 3) Cancel
	this.render.String("\n\nPickup")
}

func (this *Game) putdown() {
	// 1) Dog
	// 2) Cup
	// 3) Cancel
	this.render.String("\n\nPutdown")
}
