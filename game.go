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
	Title  string
	Action func()
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
	this.render.Section()
	this.discribeLocation()
	this.discribeLocationItems()
	this.discribeLocationExits()
	this.render.Section()
	this.discribeOptions()
	this.Play()
}

func (this *Game) setUserLocation(id string) {
	this.userloc = this.story.GetLocation(id)
}

func (this *Game) pickupItem(id string) {
	this.itemloc[id] = ""
}

func (this *Game) putdownItem(id string) {
	this.itemloc[id] = this.userloc.Id
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
	this.render.String(this.userloc.Discribe())
}

func (this *Game) discribeOptions() {
	opts := []*MenuOption{}
	opts = append(opts, this.createExitOptions()...)
	opts = append(opts, this.createLocationItemOptions()...)
	opts = append(opts, this.createUserItemOptions()...)
	opts = append(opts, this.createGameOptions()...)
	for i, opt := range opts {
		this.render.StringLine(fmt.Sprintf("%d) %s", i+1, opt.Title))
	}
	choice := this.render.AskForInt() - 1
	if choice >= 0 && choice < len(opts) {
		opts[choice].Action()
	}
}

func (this *Game) createExitOptions() []*MenuOption {
	opts := []*MenuOption{}
	for _, exit := range this.getLocationExits() {
		opts = append(opts, &MenuOption{
			Title: exit.DiscribeOptions(),
			Action: func() {
				this.setUserLocation(exit.Id)
			},
		})
	}
	return opts
}

func (this *Game) createLocationItemOptions() []*MenuOption {
	opts := []*MenuOption{}
	for _, item := range this.getLocationItems() {
		opts = append(opts, &MenuOption{
			Title: item.DiscribeOptionsPickup(),
			Action: func() {
				this.pickupItem(item.Id)
			},
		})
	}
	return opts
}

func (this *Game) createUserItemOptions() []*MenuOption {
	opts := []*MenuOption{}
	for _, item := range this.getUserItems() {
		opts = append(opts, &MenuOption{
			Title: item.DiscribeOptionsUse(),
			Action: func() {

			},
		})
		opts = append(opts, &MenuOption{
			Title: item.DiscribeOptionsPutdown(),
			Action: func() {
				this.putdownItem(item.Id)
			},
		})
	}
	return opts
}

func (this *Game) createGameOptions() []*MenuOption {
	opts := []*MenuOption{}
	opts = append(opts, &MenuOption{
		Title: this.story.DiscribeQuit(),
		Action: func() {
			this.render.Quit()
		},
	})
	return opts
}
