package main

import (
	"encoding/json"
	"sort"

	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
)

// First letter needs to be in capital
// https://stackoverflow.com/questions/26327391/golang-json-marshalstruct-returns
type Player struct {
	Name  string `json:"name" binding:"required"`
	Score int    `json:"score" binding:"required"`
}

type Players [5]Player

var players Players

var m *martini.Martini

type ByScore []Player

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].Score > a[j].Score }

// top_5 is a slice

func main() {

	players = Players{
		Player{Name: "God", Score: 1}}
	Init()
}

func getData() string {

	b, _ := json.Marshal(players)

	return string(b)
}

// Handling the sorting issues
func addData(player Player) {

	// temp is a slice
	// http://blog.denevell.org/golang-slices.html
	var slice []Player
	// players[0:len(players)] convert array to slice
	// because append is only using slices
	slice = append(players[0:len(players)], player)

	sort.Sort(ByScore(slice))

	players = [5]Player{slice[0], slice[1], slice[2], slice[3], slice[4]}

}

func Init() {
	m = martini.New()

	// Setup middleware
	m.Use(martini.Recovery())
	m.Use(martini.Logger())

	// Setup routes
	r := martini.NewRouter()
	r.Get("/", getData)

	// Using https://github.com/martini-contrib/binding
	// to create automaticly the Player struct based on JSON
	// and give the struct to addData
	// http://progadventure.blogspot.fr/2014/03/learning-go-with-martini-working-with.html
	r.Post("/", binding.Json(Player{}), addData)

	// Add the router action
	m.Action(r.Handle)
	m.Run()
}
