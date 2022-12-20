package main

import (
	"fmt"
	"strconv"
)

type Node struct {
	value    int
	next     *Node
	previous *Node
}

func CreateNode(value int, next *Node, previous *Node) *Node {
	node := new(Node)
	node.value = value

	node.next = next
	node.previous = previous

	return node
}

// [] - [] - [] - []
func FindNode(steps int, node *Node) *Node {
	if steps == 0 {
		return nil
	}
	current := node
	if steps < 0 {
		for i := steps; i <= 0; i++ {
			current = current.previous
		}
	} else {
		for i := 0; i <= steps; i++ {
			current = current.next
		}
	}

	return current
}

func DayTwenty() (interface{}, interface{}) {
	openFile := readFileByLines("inputs/d20.txt")
	if openFile == nil {
		fmt.Println("could not open puzzle input!")
		return 0, 0
	}

	nodesByPosition := []*Node{}
	var previous *Node = nil
	var start *Node = nil
	for openFile.Scanner.Scan() {
		input := openFile.Scanner.Text()

		value, _ := strconv.Atoi(input)

		node := CreateNode(value, nil, previous)

		if previous != nil {
			previous.next = node
		}

		if value == 0 {
			start = node
		}

		nodesByPosition = append(nodesByPosition, node)
		previous = node
	}

	nodesByPosition[0].previous = nodesByPosition[len(nodesByPosition)-1]
	nodesByPosition[len(nodesByPosition)-1].next = nodesByPosition[0]

	for _, node := range nodesByPosition {
		nodeBehind := FindNode(node.value, node)
		if nodeBehind == nil {
			continue
		}

		nodeAhead := nodeBehind.next
		// remove gap generated by moving current node
		node.previous.next = node.next
		node.next.previous = node.previous

		// insert node in new position
		nodeBehind.next = node
		nodeAhead.previous = node
		node.next = nodeAhead
		node.previous = nodeBehind
	}

	var first *Node = nil
	var second *Node = nil
	var third *Node = nil
	current := start
	for i := 1; i <= 3000; i++ {
		current = current.next
		if i == 1000 {
			first = current
		}
		if i == 2000 {
			second = current
		}
		if i == 3000 {
			third = current
		}
	}

	partA := first.value + second.value + third.value

	openFile.File.Close()
	return partA, 0
}
