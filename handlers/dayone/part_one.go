package dayone

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type BinarySearchTree struct {
	value *int
	left  *BinarySearchTree
	right *BinarySearchTree
}

func (b *BinarySearchTree) Insert(value int) {
	if b.value == nil {
		b.value = &value
		return
	}

	if value < *b.value {
		if b.left == nil {
			b.left = &BinarySearchTree{value: &value}
		} else {
			b.left.Insert(value)
		}
	} else {
		if b.right == nil {
			b.right = &BinarySearchTree{value: &value}
		} else {
			b.right.Insert(value)
		}
	}
}

func (b *BinarySearchTree) GetSortedList() []int {
	list := []int{}

	if b.left != nil {
		list = append(list, b.left.GetSortedList()...)
	}

	if b.value != nil {
		list = append(list, *b.value)
	}

	if b.right != nil {
		list = append(list, b.right.GetSortedList()...)
	}

	return list
}

// DayOneHandlePartOneHandle calculates the total distance between the values of
// the two lists sorted in ascending order
func DayOnePartOneHandle(file []byte) (int, error) {
	lTree := &BinarySearchTree{}
	rTree := &BinarySearchTree{}
	lines := strings.Split(string(file), "\n")
	for i, line := range lines {
		n, err := parseLineAndExtractNumbers(i, line)
		if err != nil {
			return 0, fmt.Errorf("error parsing line %d: %v", i, err)
		}

		if n == nil {
			continue
		}

		l := n[0]
		r := n[1]

		lTree.Insert(l)
		rTree.Insert(r)
	}

	lList := lTree.GetSortedList()
	rList := rTree.GetSortedList()
	totalDistance := 0
	for i := 0; i < len(lList); i++ {
		totalDistance += calcDistance(lList[i], rList[i])
	}

	return totalDistance, nil
}

func calcDistance(a, b int) int {
	return int(math.Abs(float64(a - b)))
}

func parseLineAndExtractNumbers(index int, line string) (*[2]int, error) {
	line = strings.TrimFunc(line, func(r rune) bool {
		return r == '\r' || r == '\n' || r == ' ' || r == '\t'
	})

	if line == "" {
		return nil, nil
	}

	parsedLine := strings.Split(string(line), ",")
	if len(parsedLine) != 2 {
		return nil, fmt.Errorf("Error parsing line %d: %s", index, line)
	}

	l, _ := strconv.Atoi(parsedLine[0])
	r, _ := strconv.Atoi(parsedLine[1])

	return &[2]int{l, r}, nil
}
