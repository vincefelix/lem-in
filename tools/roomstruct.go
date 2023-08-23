package lem_in

type Room struct {
	Name  string
	X     int
	Y     int
	Start bool
	End   bool
	Links []*Room
}

type Link struct {
	Room1 string
	Room2 string
}

var numAnts int
var Rooms []Room
var links []Link
