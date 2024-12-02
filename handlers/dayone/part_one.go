package dayone

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Node struct {
	value  int
	parent *Node
	left   *Node
	right  *Node
}

func newNode(value int) *Node {
	return &Node{value, nil, nil, nil}
}

func (n *Node) Insert(node *Node) {
	if node.value < n.value {
		if n.left == nil {
			n.left = node
			node.parent = n
		} else {
			n.left.Insert(node)
		}
	} else {
		if n.right == nil {
			n.right = node
			node.parent = n
		} else {
			n.right.Insert(node)
		}
	}
}

func (n *Node) String(level int) string {
	str := fmt.Sprintf("%d:", n.value)
	if n.left != nil {
		tab := strings.Repeat("\t", level+1)
		str = fmt.Sprintf("%s\n%sL:%s", str, tab, n.left.String(level+1))
	}

	if n.right != nil {
		tab := strings.Repeat("\t", level+1)
		str = fmt.Sprintf("%s\n%sR:%s", str, tab, n.right.String(level+1))
	}

	return str
}

type BinarySearchTree struct {
	node     *Node
	smallest *Node
}

// Insert inserts values smaller than the current node at left and bigger or equal at right
// while keeping track of the reference for the smallest value in the tree
func (b *BinarySearchTree) Insert(value int) {
	if b.node == nil {
		b.node = newNode(value)
		b.smallest = b.node
		return
	}

	node := newNode(value)
	b.node.Insert(node)

	if value < b.smallest.value {
		b.smallest = node
	}
}

// Pop returns the current smallest value and removes it
func (b *BinarySearchTree) Pop() *Node {
	smallest := b.smallest

	// since the smallest is always the leftmost node we only need to
	// swap with the right node or the parent node

	if b.smallest.right != nil {
		parent := b.smallest.parent
		n := b.smallest.right

		// link the parent of the previous smallest to the rightmost node
		n.parent = parent

		// try to find the leftmost node of the current node
		for {
			if n.left == nil {
				break
			}

			// if some node on the left was found, the parent is the current node
			parent = n
			n = n.left
		}

		b.smallest = n

		// this is needed when the smallest is the rightmost node, so we need to link the parent of the
		// previous smallest to the current smallest node
		if parent != nil {
			parent.left = b.smallest
		}

		// if we don't have a right node, we can assume that the parent is the new smallest node
	} else if b.smallest.parent != nil {
		b.smallest = b.smallest.parent
		b.smallest.left = nil
	}

	return smallest
}

func (b *BinarySearchTree) String() string {
	return b.node.String(0)
}

// DayOneHandlePartOneHandle calculates the total distance between the values of
// the two lists sorted in ascending order
func DayOnePartOneHandle(file []byte) (int, error) {
	lTree := &BinarySearchTree{}
	rTree := &BinarySearchTree{}
	lines := strings.Split(string(file), "\n")
	count := 0
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
		count++
	}

	totalDistance := 0
	for i := 0; i < count; i++ {
		l := lTree.Pop()
		r := rTree.Pop()

		if l == nil || r == nil {
			break
		}

		d := calcDistance(l.value, r.value)
		totalDistance += d
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
