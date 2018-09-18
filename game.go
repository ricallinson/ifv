package main

import (
	"fmt"
)

type Game struct {
	story       *Story
	render      Renderer
	userloc     *Location
	itemloc     map[string]string
	itemsHidden []string
	exitsHidden []string
}

type MenuOption struct {
	Title  string
	Action func()
}

func CreateGame(p string, r Renderer) *Game {
	this := &Game{
		story:       CreateStory(p),
		render:      r,
		itemloc:     map[string]string{},
		itemsHidden: []string{},
		exitsHidden: []string{},
	}
	this.userloc = this.story.GetStartLocation()
	this.itemloc = this.story.GetItemLocations()
	this.itemsHidden = this.story.GetHiddenItems()
	this.exitsHidden = this.story.GetHiddenExits()
	// Validate game data here.
	return this
}

func (this *Game) StoryToYaml() {
	fmt.Println(string(interfaceToYaml(this.story)))
}

func (this *Game) Play() {
	this.render.Section()
	if this.story.IsLocationTheEnd(this.userloc.Id) {
		this.discribeLocation()
		this.render.Quit()
	}
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
		if loc == "" && !containsString(this.itemsHidden, itemId) {
			items = append(items, this.story.GetItem(itemId))
		}
	}
	return items
}

func (this *Game) getLocationItems() []*Item {
	items := []*Item{}
	for itemId, loc := range this.itemloc {
		if loc == this.userloc.Id && !containsString(this.itemsHidden, itemId) {
			items = append(items, this.story.GetItem(itemId))
		}
	}
	return items
}

func (this *Game) getLocationExits() []*LocationExit {
	exits := []*LocationExit{}
	for _, exit := range this.userloc.Exits {
		if !containsString(this.exitsHidden, exit.Id) {
			exits = append(exits, exit)
		}
	}
	return exits
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

func (this *Game) executeResult(exits []string, items []string) {
	this.exitsHidden = toggleStrings(this.exitsHidden, exits)
	this.itemsHidden = toggleStrings(this.itemsHidden, items)
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
				var s string
				if item.Conditions.Test(this.userloc, this.getLocationItems(), this.getUserItems()) {
					this.executeResult(item.Result.Exits, item.Result.Items)
					s = item.DiscribeSuccess()
				} else {
					s = item.DiscribeFailure()
				}
				this.render.String(s)
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
