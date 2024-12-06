package dayfive

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	value  int
	before NodeMap
	after  NodeMap
}

func newNode(value int) *Node {
	return &Node{value, nil, nil}
}

func (n *Node) InsertBefore(node *Node) {
	if n.before == nil {
		n.before = newNodeMap(node)
		return
	}

	n.before[node.value] = node
}

func (n *Node) InsertAfter(node *Node) {
	if n.after == nil {
		n.after = newNodeMap(node)
		return
	}

	n.after[node.value] = node
}

func (n *Node) isAfterOf(value int) bool {
	if n.before == nil {
		return false
	}
	_, ok := n.before[value]
	return ok
}

func (n *Node) isBeforeOf(value int) bool {
	if n.after == nil {
		return false
	}
	_, ok := n.after[value]
	return ok
}

type NodeMap map[int]*Node

func newNodeMap(node *Node) NodeMap {
	nm := NodeMap{}
	nm[node.value] = node
	return nm
}

func (nm NodeMap) Has(value int) bool {
	_, ok := nm[value]
	return ok
}

func DayFivePartOneHandle(file []byte) int {
	lists := strings.Split(string(file), "\n\n")

	pageOrderingRules := strings.Split(lists[0], "\n")
	nodeMap, err := createNodeMapFromList(pageOrderingRules)
	if err != nil {
		panic(err)
	}

	pageUpdates := strings.Split(lists[1], "\n")
	sumOfMiddleValuesForCorrectUpdates := 0
	for _, u := range pageUpdates {
		values := strings.Split(u, ",")
		correct, err := isUpdateSequenceCorrect(*nodeMap, values)
		if err != nil {
			panic(err)
		}

		if !correct {
			continue
		}

		i := int(len(values) / 2)
		middle := values[i]
		middleValue, _ := strconv.Atoi(middle)
		sumOfMiddleValuesForCorrectUpdates += middleValue
	}

	return sumOfMiddleValuesForCorrectUpdates
}

func isUpdateSequenceCorrect(nodeMap NodeMap, values []string) (bool, error) {
	for i, value := range values {
		valueInt, _ := strconv.Atoi(value)
		node, ok := nodeMap[valueInt]
		if !ok {
			return false, fmt.Errorf("value not found in nodemap: %d, values %v", valueInt, values)
		}

		for j := i - 1; j >= 0; j-- {
			valueBefore, _ := strconv.Atoi(values[j])
			if !node.isAfterOf(valueBefore) {
				return false, nil
			}
		}

		for j := i + 1; j < len(values); j++ {
			valueAfter, _ := strconv.Atoi(values[j])
			if !node.isBeforeOf(valueAfter) {
				return false, nil
			}
		}
	}

	return true, nil
}

func createNodeMapFromList(list []string) (*NodeMap, error) {
	nodeMap := NodeMap{}
	for _, value := range list {
		values := strings.Split(value, "|")
		if len(values) != 2 {
			return nil, errors.New("invalid input")
		}

		value1, _ := strconv.Atoi(values[0])
		value2, _ := strconv.Atoi(values[1])

		n1 := newNode(value1)
		if nodeMap.Has(value1) {
			n1 = nodeMap[value1]
		} else {
			nodeMap[value1] = n1
		}

		n2 := newNode(value2)
		if nodeMap.Has(value2) {
			n2 = nodeMap[value2]
		} else {
			nodeMap[value2] = n2
		}

		n1.InsertAfter(n2)
		n2.InsertBefore(n1)
	}

	return &nodeMap, nil
}
