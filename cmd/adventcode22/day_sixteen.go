package main

import (
	"fmt"
	"strings"
)

type ValveNode struct {
	start    *Valve
	end      *Valve
	distance int
}

type Valve struct {
	Name          string
	Rate          int
	connectionIds []string
}

type Path struct {
	TotalFlow       int
	TimeUsed        int
	FlowEveryMinute int
	visited         []string
	opened          map[string]bool
}

func CreateValve() *Valve {
	valve := new(Valve)
	valve.connectionIds = []string{}

	return new(Valve)
}

func CreatePath() *Path {
	p := new(Path)
	p.TotalFlow = 0
	p.TimeUsed = 0
	p.FlowEveryMinute = 0
	p.opened = make(map[string]bool)

	return p
}

func ParseValve(input string) *Valve {
	valve := CreateValve()

	fmt.Sscanf(input, "Valve %s has flow rate=%d; tunnels lead to valves", &valve.Name, &valve.Rate)

	connections := strings.Split(input, ":")[1]
	for _, v := range strings.Split(connections, ",") {
		valve.connectionIds = append(valve.connectionIds, strings.TrimSpace(v))
	}

	return valve
}

func FindClosestValuableNeighbor(start *Valve, parent *Valve, connectionIds []string, distance int, valves map[string]*Valve) []ValveNode {
	nodes := []ValveNode{}
	for _, id := range connectionIds {
		connection := valves[id]
		if connection.Rate > 0 {
			nodes = append(nodes, ValveNode{start, connection, distance})
		} else {
			connectionIds := []string{}
			for _, next := range connection.connectionIds {
				if next != parent.Name {
					connectionIds = append(connectionIds, next)
				}
			}
			nodes = append(nodes, FindClosestValuableNeighbor(start, connection, connectionIds, distance+1, valves)...)
		}
	}
	return nodes
}

func BuildGraph(valves map[string]*Valve) map[string][]ValveNode {
	nodes := make(map[string][]ValveNode)

	for _, valve := range valves {
		nodes[valve.Name] = FindClosestValuableNeighbor(valve, valve, valve.connectionIds, 1, valves)
	}

	return nodes
}

func (p *Path) Copy() *Path {
	path := CreatePath()

	path.opened = p.opened
	path.visited = p.visited
	path.FlowEveryMinute = p.FlowEveryMinute
	path.TotalFlow = p.TotalFlow

	return path
}

func (p *Path) Update(graph map[string][]ValveNode, valves map[string]*Valve) ([]*Path, bool) {
	if p.TimeUsed >= 30 {
		return []*Path{}, false
	}

	p.TotalFlow += p.FlowEveryMinute
	p.TimeUsed++

	currentIndex := p.visited[len(p.visited)-1]
	current := valves[currentIndex]
	_, hasOpened := p.opened[currentIndex]

	hasTimeLeft := p.TimeUsed < 30

	newPaths := []*Path{}
	if current.Rate > 0 && !hasOpened {
		p.FlowEveryMinute += current.Rate
		p.opened[currentIndex] = true

		return newPaths, hasTimeLeft
	}

	for _, connection := range graph[currentIndex] {
		newPath := p.Copy()
		newPath.visited = append(newPath.visited, connection.end.Name)
		newPath.TimeUsed += connection.distance
		newPaths = append(newPaths, newPath)
	}

	return newPaths, hasTimeLeft
}

func DaySixteen() (interface{}, interface{}) {
	openFile := readFileByLines("inputs/d16.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return 0, 0
	}

	valves := make(map[string]*Valve)
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()

		valve := ParseValve(input)
		valves[valve.Name] = valve
	}

	graph := BuildGraph(valves)

	initial := graph["AA"]

	fmt.Println(initial)

	// 30 minutes
	// create a frontier with all neighbors priority is

	openFile.File.Close()
	return 0, 0
}
