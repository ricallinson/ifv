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

type MenuOption struct {
	Title string
	Fn    func()
}

func CreateGame(p string, r Renderer) *Game {
	this := &Game{
		story:   readYamlFileToStory(p),
		render:  r,
		itemloc: map[string]string{},
	}
	this.setUserLocation(this.story.LocationStart)
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

func (this *Game) setUserLocation(id string) {
	this.userloc = this.story.GetLocation(id)
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

func (this *Game) discribeLocationExits() {
	exits := this.getLocationExits()
	for _, exit := range exits {
		this.render.String(exit.Discribe())
	}
}

func (this *Game) discribeLocationItems() {
	items := this.getLocationItems()
	for _, item := range items {
		this.render.String(item.Discribe())
	}
}

func (this *Game) discribeLocation() {
	// You are in a room. On the floor is an old hammer. A cups sits alone.
	this.render.String(this.userloc.Discribe())
	this.discribeLocationItems()
	this.discribeLocationExits()
}

// Creates the main navigation menu.
func (this *Game) menu() {
	exits := this.getLocationExits()
	locItems := this.getLocationItems()
	userItems := this.getUserItems()
	options := []*MenuOption{}
	if len(exits) > 0 {
		options = append(options, &MenuOption{randStringSelection(this.story.LocationMoveTitle), this.move})
	}
	if len(locItems) > 0 {
		options = append(options, &MenuOption{randStringSelection(this.story.ItemPickupTitle), this.pickup})
	}
	if len(userItems) > 0 {
		options = append(options, &MenuOption{randStringSelection(this.story.ItemPutdownTitle), this.putdown})
		options = append(options, &MenuOption{randStringSelection(this.story.ItemUseTitle), this.use})
	}
	this.render.String("\n")
	for i := 0; i < len(options); i++ {
		this.render.String(fmt.Sprintf("\n%d ) %s\n", i+1, options[i].Title))
	}
	this.render.String("\n:")
	choice := this.render.AskForInt()
	if choice >= 1 && choice <= len(options) {
		options[choice-1].Fn()
	}
}

// Creates the move menu.
func (this *Game) move() {
	exits := this.getLocationExits()
	for i, s := range exits {
		this.render.String(fmt.Sprintf("\n%d ) %s\n", i+1, s.Id))
	}
	this.render.String("\n:")
	choice := this.render.AskForInt() - 1
	if choice < len(exits) {
		this.setUserLocation(exits[choice].Id)
	}
}

// Creats the use menu.
func (this *Game) use() {
	this.render.String("\n\n***NOT IMPLEMENTED YET***\n\n")
}

// Creates the pickup menu.
func (this *Game) pickup() {
	items := this.getLocationItems()
	for i, s := range items {
		this.render.String(fmt.Sprintf("\n%d ) %s\n", i+1, s.Id))
	}
	this.render.String("\n:")
	choice := this.render.AskForInt() - 1
	if choice < len(items) {
		this.itemloc[items[choice].Id] = ""
	}
}

// Creates the putdown menu.
func (this *Game) putdown() {
	items := this.getUserItems()
	for i, s := range items {
		this.render.String(fmt.Sprintf("\n%d ) %s\n", i+1, s.Id))
	}
	this.render.String("\n:")
	choice := this.render.AskForInt() - 1
	if choice < len(items) {
		this.itemloc[items[choice].Id] = this.userloc.Id
	}
}
