package daynine

import (
	"strconv"
	"strings"
)

const gap = -1

func DayNinePartOneHandle(file []byte) int {
	chars := strings.Split(string(file), "")
	blocks := make([]int, len(chars))
	for i := range chars {
		v, err := strconv.Atoi(chars[i])
		if err != nil {
			panic(err)
		}

		blocks[i] = v
	}

	fileBlocks := []int{}
	fileID := int(0)
	gapCount := 0

	for i, v := range blocks {
		isEven := i%2 == 0
		if isEven {
			for j := 0; j < v; j++ {
				fileBlocks = append(fileBlocks, fileID)
			}

			fileID++
		} else {
			for j := 0; j < v; j++ {
				fileBlocks = append(fileBlocks, gap)
				gapCount++
			}
		}
	}

	i := len(fileBlocks) - 1
	for !hasFinished(fileBlocks) && i >= 0 {
		if fileBlocks[i] == gap {
			i--
			continue
		}

		gapIdx, hasGap := findNextGapAvailable(fileBlocks)
		if !hasGap {
			break
		}

		fileBlocks[gapIdx], fileBlocks[i] = fileBlocks[i], fileBlocks[gapIdx]
		i--
	}

	sum := 0
	for i, v := range fileBlocks {
		if v == gap {
			continue
		}
		sum += v * i
	}

	return sum
}

func hasFinished(blocks []int) bool {
	foundGap := false
	for _, v := range blocks {
		if v == gap {
			foundGap = true
		} else if foundGap && v != gap {
			return false
		}
	}

	return true
}

func findNextGapAvailable(blocks []int) (int, bool) {
	for i, v := range blocks {
		if v == gap {
			return i, true
		}
	}

	return 0, false
}
