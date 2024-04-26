package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	var i1 Item
	fmt.Println(i1)
	fmt.Printf("i1: %#v\n", i1)

	i2 := Item{1, 2}
	fmt.Printf("i2: %#v\n", i2)

	i3 := Item{
		Y: 10,
		// X: 20,
	}
	fmt.Printf("i3: %#v\n", i3)
	fmt.Println(NewItem(10, 20))
	fmt.Println(NewItem(10, -20))

	i3.Move(100, 200)
	fmt.Printf("i3 (move): %#v\n", i3)

	p1 := Player{
		Name: "Parzival",
		Item: Item{500, 300},
	}

	fmt.Printf("p1: %#v\n", p1)
	fmt.Printf("p1.X: %#v\n", p1.X)
	fmt.Printf("p1.Item.X: %#v\n", p1.Item.X)
	p1.Move(400, 600)
	fmt.Printf("p1 (move): %#v\n", p1)

	ms := []mover{
		&i1,
		&p1,
		&i2,
	}
	moveAll(ms, 0, 0)

	for _, m := range ms {
		fmt.Println(m)
	}

	k := Jade
	fmt.Println("k:", k)
	fmt.Println("key:", Key(17))

	// time.Time import json.Marshaler interface
	// json.NewEncoder(os.Stdout).Encode(time.Now())

	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)
	p1.FoundKey(Jade)
	fmt.Println(p1.Keys)

	players := []Player{
		{
			Name: "Jack Nelson",
			Item: Item{2, 3},
		},
		{
			Name: "Giovanni",
			Item: Item{500, 600},
		},
		{
			Name: "Malaquias",
			Item: Item{400, 300},
		},
	}

	fmt.Println(players)

	sortByDistance(players, 500, 600)

	fmt.Println(players)
}

type ByDistance []Player

func (a ByDistance) Len() int      { return len(a) }
func (a ByDistance) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByDistance) Less(i, j int) bool {
	first := a[i].X + a[i].Y
	second := a[j].X + a[j].Y
	return first < second
}

func sortByDistance(players []Player, x, y int) {
	distance := x + y
	sort.Slice(players, func(i, j int) bool {
		first := (players[i].X + players[i].Y) - distance
		second := (players[j].X + players[j].Y) - distance
		if first < 0 {
			first *= -1
		}
		if second < 0 {
			second *= -1
		}
		return first < second
	})
}

// String representation of our type
// Implement fmt.Stringer interface
func (k Key) String() string {
	switch k {
	case Jade:
		return "jade"
	case Copper:
		return "copper"
	case Crystal:
		return "crystal"
	}

	return fmt.Sprintf("<Key %d>", k)
}

/*
	Exercise

- Add "Keys" field to Player which is a slice of Key
- Add a "FoundKey (k Key) error" method to player which will add k to Key if it's not there
  - Err if k is not one of the known keys
*/

func (p *Player) FoundKey(k Key) error {
	// if k != Jade && k != Copper && k != Crystal {
	if k < Jade || k >= invalidKey {
		return fmt.Errorf("Invalid key: %#v", k)
	}
	// if !containsKey(p.Keys, k) {
	if !slices.Contains(p.Keys, k) {
		p.Keys = append(p.Keys, k)
	}
	return nil
}

func containsKey(keys []Key, k Key) bool {
	for _, k2 := range keys {
		if k2 == k {
			return true
		}
	}
	return false
}

// Go's version of "enum"
const (
	Jade Key = iota + 1
	Copper
	Crystal
	invalidKey // internal (not exported)
)

type Key byte

// Rule of thumb: Accept interfaces, return types

func moveAll(ms []mover, x, y int) {
	for _, m := range ms {
		m.Move(x, y)
	}
}

type mover interface {
	Move(x, y int)
	// Move(int, int)
}

type Player struct {
	Name string
	Item // Embed Item
	Keys []Key
}

// i is called "the receiver"
// if you want to mutate, use pointer receiver
func (i *Item) Move(x, y int) {
	i.X = x
	i.Y = y
}

// func NewItem(x, y int) Item
// func NewItem(x, y int) *Item
// func NewItem(x, y int) (Item, error)
func NewItem(x, y int) (*Item, error) {
	if x < 0 || x > maxX || y < 0 || y > maxY {
		return nil, fmt.Errorf("%d/%d out of bounds %d/%d", x, y, maxX, maxY)
	}
	i := Item{
		X: x,
		Y: y,
	}
	// The Go compiler does "escape analysis" and will allocate i on the heap
	return &i, nil
}

// zero vs missing value

const (
	maxX = 1000
	maxY = 600
)

// Item is an item in the game
type Item struct {
	X int
	Y int
}
