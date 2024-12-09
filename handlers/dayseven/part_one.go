package dayseven

import (
	"fmt"
	"strconv"
	"strings"
)

type Line struct {
	Result int64
	Values []int
	Op     []string
}

func (l *Line) String() string {
	str := ""
	for i := range l.Values {
		str += fmt.Sprintf("%d", l.Values[i])
		if i < len(l.Op)-1 {
			str += " " + l.Op[i] + " "
		}
	}
	return str
}

func DaySevenPartOneHandle(file []byte) int64 {
	paths := strings.Split(string(file), "\n")
	data := make([]Line, len(paths))
	for i := range paths {
		parts := strings.Split(paths[i], ": ")
		result, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			panic(err)
		}

		vStr := strings.Split(parts[1], " ")
		values := make([]int, len(vStr))
		for j := range vStr {
			v, err := strconv.Atoi(vStr[j])
			if err != nil {
				panic(err)
			}
			values[j] = v
		}

		data[i] = Line{Result: result, Values: values}
	}

	total := int64(0)
	for _, line := range data {
		combs := generateCombinations(line.Values)
		for _, comb := range combs {
			r := comb.Eval()
			if r == line.Result {
				total += r
				break
			}
		}
	}

	return total
}

const sumOperator = "+"
const multiplyOperator = "*"

type Combination struct {
	Nums       []int
	Operations []string
}

func (c Combination) String() string {
	str := ""
	for i := range c.Nums {
		str += fmt.Sprintf("%d", c.Nums[i])
		if i <= len(c.Operations)-1 {
			str += " " + c.Operations[i] + " "
		}
	}
	return str
}

func (c Combination) Eval() int64 {
	acc := int64(c.Nums[0])
	for i := 1; i < len(c.Nums); i++ {
		op := c.Operations[i-1]
		switch op {
		case sumOperator:
			acc += int64(c.Nums[i])
		case multiplyOperator:
			acc *= int64(c.Nums[i])
		}
	}

	return acc
}

func generateCombinations(nums []int) []Combination {
	operations := []string{sumOperator, multiplyOperator}
	var results []Combination
	if len(nums) == 1 {
		// Base case: only one number left
		results = append(results, Combination{
			Nums:       []int{nums[0]},
			Operations: []string{},
		})
		return results
	}

	// Generate combinations recursively
	for _, op := range operations {
		subResults := generateCombinations(nums[1:])
		for i := range subResults {
			combined := Combination{
				Nums:       append([]int{nums[0]}, subResults[i].Nums...),
				Operations: append([]string{op}, subResults[i].Operations...),
			}

			results = append(results, combined)
		}
	}

	return results
}
