package main

import "fmt"

type State interface {
	Execute()
}

type Starting struct {
}

func (s Starting) Execute() {
	fmt.Println("Execute starting...")
}

type Processing struct {
}

func (p Processing) Execute() {
	fmt.Println("Execute processing...")
}

type Ending struct {
}

func (e Ending) Execute() {
	fmt.Println("Execute ending...")
}

type Game struct {
	workingOrder []State
	progress     int
}

func NewGame(workingOrder []State) Game {
	return Game{
		workingOrder: workingOrder,
		progress:     0,
	}
}

func (g *Game) Continue() {
	g.workingOrder[g.progress].Execute()

	g.progress++
	if g.progress > len(g.workingOrder) {
		g.progress = 0
	}
}

func main() {
	game := NewGame([]State{Starting{}, Processing{}, Ending{}})
	game.Continue()
	game.Continue()
	game.Continue()
}
