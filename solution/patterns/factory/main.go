package main

import "fmt"

const (
	lowRespLevel    = 3
	middleRespLevel = 6
	highRespLevel   = 10
)

type Worker interface {
	Work()
}

type Miner struct {
	name                string
	responsibilityLevel int
}

func (m Miner) Work() {
	fmt.Println("Miner is working")
}

type Director struct {
	name                string
	responsibilityLevel int
}

func (d Director) Work() {
	fmt.Println("Director is working")
}

func HireWorker(name string, responsibilityLevel int) Worker {
	if responsibilityLevel >= highRespLevel {
		return &Director{name, responsibilityLevel}
	}
	return &Miner{name, responsibilityLevel}
}

func main() {
	worker := HireWorker("Alex", 12)
	worker.Work()

	anotherWorker := HireWorker("Alex", 3)
	anotherWorker.Work()
}
