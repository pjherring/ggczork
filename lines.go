package main

import (
	"fmt"
	"strings"
)

const Nowhere = -1

var Intro string = `ZORK IV: The Lost Great Gotham Empire
Copyright (c) 1985 GGCCOM, Inc. All Rights Reserved
Revision 94 / Serial number 16241898`

var UnknownResponse = Response{
	Response: "You can't see any such thing.",
}

var EmptyResponse = Response{
	Response: "I beg your pardon?",
}

type Line struct {
	Text      string
	Options   map[string]Response
	Intro     bool
	Countdown bool
}

func (l Line) Say() {
	fmt.Printf("\n%s\n", l.Text)
	if l.Countdown {
		StartCountdown()
	}
}

func (l Line) FindResponse(i string) Response {
	if l.Intro {
		return GotoResponse(1)
	}

	i = strings.TrimSpace(strings.ToLower(i))
	if i == "" {
		return EmptyResponse
	}

	for opt, resp := range l.Options {
		opt = strings.ToLower(opt)
		if len(opt) == 1 && opt == i {
			return resp
		} else if len(opt) > 1 && strings.Contains(i, opt) {
			return resp
		}
	}

	return UnknownResponse
}

type Response struct {
	Goto     int
	Response string
}

func GotoResponse(i int) Response {
	return Response{i, ""}
}

func NowhereResponse(s string) Response {
	return Response{Nowhere, s}
}

var Lines []Line = []Line{
	Line{
		Text: `ZORK IV: The Lost Great Gotham Empire
Copyright (c) 1985 GGCCOM, Inc. All Rights Reserved
Revision 94 / Serial number 16241898`,
		Intro: true,
	},
	Line{
		Text: `West of House
You are standing in an open field west of a white house, with a boarded front door.
There is a small mailbox here.`,
		Options: map[string]Response{
			"Open Mailbox": GotoResponse(2),
			"Walk West":    GotoResponse(3),
			"West":         GotoResponse(3),
			"W":            GotoResponse(3),
			"Walk East":    NowhereResponse("The door is locked, and there is evidently no key."),
			"East":         NowhereResponse("The door is locked, and there is evidently no key."),
			"E":            NowhereResponse("The door is locked, and there is evidently no key."),
			"Walk South":   GotoResponse(5),
			"South":        GotoResponse(5),
			"S":            GotoResponse(5),
			"Walk North":   GotoResponse(6),
			"North":        GotoResponse(6),
			"N":            GotoResponse(6),
		},
	},
	Line{
		Text: `You open the mailbox, revealing a small leaflet.`,
		Options: map[string]Response{
			"Read leaflet": GotoResponse(4),
			"Take leaflet": GotoResponse(4),
			"Walk West":    GotoResponse(3),
			"West":         GotoResponse(3),
			"W":            GotoResponse(3),
			"Walk East":    NowhereResponse("The door is locked, and there is evidently no key."),
			"East":         NowhereResponse("The door is locked, and there is evidently no key."),
			"E":            NowhereResponse("The door is locked, and there is evidently no key."),
			"Walk South":   GotoResponse(5),
			"South":        GotoResponse(5),
			"S":            GotoResponse(5),
			"Walk North":   GotoResponse(6),
			"North":        GotoResponse(6),
			"N":            GotoResponse(6),
		},
	},
	Line{
		Text: `Forest
This is a forest, with trees in all directions around you.`,
		Options: map[string]Response{
			"Walk West":  GotoResponse(3),
			"West":       GotoResponse(3),
			"W":          GotoResponse(3),
			"Walk East":  GotoResponse(1),
			"East":       GotoResponse(1),
			"E":          GotoResponse(1),
			"Walk North": NowhereResponse("You can't go that way."),
			"North":      NowhereResponse("You can't go that way."),
			"N":          NowhereResponse("You can't go that way."),
			"Walk South": NowhereResponse("You can't go that way."),
			"South":      NowhereResponse("You can't go that way."),
			"S":          NowhereResponse("You can't go that way."),
		},
	},
	Line{
		Text: `WELCOME TO THE GREATEST CHALLENGE YOU WILL FACE
This is only the beginning of a game of adventure, danger, and low cunning. In it you will explore some of the most amazing territory ever seen by mortal man.  Hardened adventurers have run screaming from the terrors contained within!

In this challenge, the intrepid explorer delves into the forgotten secrets of a lost labyrinth deep in the canyons of a metropolis, searching for vast treasures long hidden from prying eyes, treasures guarded by fearsome monsters and diabolical traps!

There is no HELP or INFO in this game.`,
		Options: map[string]Response{
			"Walk West":  GotoResponse(3),
			"West":       GotoResponse(3),
			"W":          GotoResponse(3),
			"Walk East":  NowhereResponse("The door is locked, and there is evidently no key."),
			"East":       NowhereResponse("The door is locked, and there is evidently no key."),
			"E":          NowhereResponse("The door is locked, and there is evidently no key."),
			"Walk South": GotoResponse(5),
			"South":      GotoResponse(5),
			"S":          GotoResponse(5),
			"Walk North": GotoResponse(6),
			"North":      GotoResponse(6),
			"N":          GotoResponse(6),
		},
	},
	Line{
		Text: `South of House
You are facing the south side of a white house.  There is no door here, and all the windows are barred.`,
		Options: map[string]Response{
			"Walk West":  GotoResponse(1),
			"West":       GotoResponse(1),
			"W":          GotoResponse(1),
			"Walk East":  GotoResponse(7),
			"East":       GotoResponse(7),
			"E":          GotoResponse(7),
			"Walk North": NowhereResponse("The windows are all barred"),
			"North":      NowhereResponse("The windows are all barred"),
			"N":          NowhereResponse("The windows are all barred"),
			"Walk South": GotoResponse(3),
			"South":      GotoResponse(3),
			"S":          GotoResponse(3),
		},
	},
	Line{
		Text: `North of House
You are facing the north side of a white house.  There is no door here, and all the windows are barred.
			`,
		Options: map[string]Response{
			"Walk West":  GotoResponse(1),
			"West":       GotoResponse(1),
			"W":          GotoResponse(1),
			"Walk East":  GotoResponse(7),
			"East":       GotoResponse(7),
			"E":          GotoResponse(7),
			"Walk South": NowhereResponse("The windows are all barred"),
			"South":      NowhereResponse("The windows are all barred"),
			"S":          NowhereResponse("The windows are all barred"),
			"Walk North": GotoResponse(3),
			"North":      GotoResponse(3),
			"N":          GotoResponse(3),
		},
	},
	Line{
		Text: `Behind House
You are behind the white house.  In one corner of the house there is a small window which is slightly ajar.
			`,
		Options: map[string]Response{
			"open window": GotoResponse(8),
			"Walk West":   NowhereResponse("The window is closed"),
			"West":        NowhereResponse("The window is closed"),
			"W":           NowhereResponse("The window is closed"),
			"Walk East":   GotoResponse(3),
			"East":        GotoResponse(3),
			"E":           GotoResponse(3),
			"Walk South":  GotoResponse(5),
			"South":       GotoResponse(5),
			"S":           GotoResponse(5),
			"Walk North":  GotoResponse(6),
			"North":       GotoResponse(6),
			"N":           GotoResponse(6),
		},
	},
	Line{
		Text: "With Great effort, you open the window far enough to allow entry",
		Options: map[string]Response{
			"Enter":      GotoResponse(9),
			"Walk West":  NowhereResponse("The window is closed"),
			"West":       NowhereResponse("The window is closed"),
			"W":          NowhereResponse("The window is closed"),
			"Walk East":  GotoResponse(3),
			"East":       GotoResponse(3),
			"E":          GotoResponse(3),
			"Walk South": GotoResponse(5),
			"South":      GotoResponse(5),
			"S":          GotoResponse(5),
			"Walk North": GotoResponse(6),
			"North":      GotoResponse(6),
			"N":          GotoResponse(6),
		},
	},
	Line{
		Text: `Kitchen
You are in the kitchen of the white house.  A table seems to have been used recently for the preparation of food.  A passage leads to the west and a dark staircase can be seen leading upward.  To the east is a small window which is open.

Then, a strange figure appears - dressed oddly. Perhaps like an eingineer? In a blue work suit with military patches and a 5-panel hat. "Welcome to the greatest challennge you will ever face," he says. You do not know what to say.`,
		Countdown: true,
	},
}
