package day2

import (
	"fmt"
	"strconv"
	"strings"

	"AdventOfCode2021Go/tools"
)

func Part2() {

}

func Part1(filePath string) {
	lines := tools.ReadFile(filePath)
	length := len(lines)
	position := 0
	depth := 0
	for i := 0; i < length; i++ {
		lineArray := strings.Fields(lines[i])
		direction := lineArray[0]
		if direction == "forward" {
			s, err := strconv.ParseInt(lineArray[1], 0, 0)
			if err == nil {
				position = position + int(s)
			}

		} else {
			if direction == "up" {
				s, err := strconv.ParseInt(lineArray[1], 0, 0)
				if err == nil {
					depth = depth - int(s)
				}

			} else {
				if direction == "down" {
					s, err := strconv.ParseInt(lineArray[1], 0, 0)
					if err == nil {
						depth = depth + int(s)
					}

				}

			}
		}
	}
	result := position * depth
	fmt.Println(result)
}
