package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseGame(s string) int {
	sp := strings.Split(s, ":")

	// gameID, err := strconv.Atoi(strings.Trim(sp[0], "Game "))
	// if err != nil {
	// 	panic("failed to find game id")
	// }

	sets := strings.Split(sp[1], ";")

	red, green, blue := 0, 0, 0

	for _, set := range sets {
		setTrimed := strings.TrimSpace(set)

		cubes := strings.Split(setTrimed, ", ")
		for _, cube := range cubes {
			num, err := strconv.Atoi(strings.Fields(cube)[0])
			if err != nil {
				panic("failed to find cube number")
			}
			switch cube[len(cube)-3:] {
			case "red":
				if num > red {
					red = num
				}
			case "een":
				if num > green {
					green = num
				}
			case "lue":
				if num > blue {
					blue = num
				}
			}
		}
	}

	return red * green * blue
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("failed to open input.txt")
		os.Exit(1)
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += parseGame(scanner.Text())
	}

	fmt.Println("result=", sum)
}
