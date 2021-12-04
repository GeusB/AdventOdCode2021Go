package day2

import (
	"fmt"
	"strconv"
	"strings"

	"AdventOfCode2021Go/tools"
)

func Part2(filePath string) {
	lines := tools.ReadFile(filePath)
	length := len(lines)
	aim := 0
	position := 0
	depth := 0
	for i := 0; i < length; i++ {
		lineArray := strings.Fields(lines[i])
		direction := lineArray[0]
		value, err := strconv.ParseInt(lineArray[1], 0, 0)
		if err == nil {
			if direction == "forward" {
				delta := int(value)
				position = position + delta
				depth = depth + delta*aim
			} else {
				if direction == "up" {
					delta := int(value)
					aim = aim - delta
				} else {
					if direction == "down" {
						delta := int(value)
						aim = aim + delta
					}
				}
			}
		}
	}
	result := position * depth
	fmt.Println(result)
}

func Part1(filePath string) {
	lines := tools.ReadFile(filePath)
	length := len(lines)
	position := 0
	depth := 0
	for i := 0; i < length; i++ {
		lineArray := strings.Fields(lines[i])
		direction := lineArray[0]
		value, err := strconv.ParseInt(lineArray[1], 0, 0)
		if err == nil {
			if direction == "forward" {
				position = position + int(value)
			}

		} else {
			if direction == "up" {
				depth = depth - int(value)
			} else {
				if direction == "down" {
					depth = depth + int(value)
				}

			}
		}
	}
	result := position * depth
	fmt.Println(result)
}
