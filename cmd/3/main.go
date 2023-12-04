package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	//wonGames := 0

	totalPowerSet := 0

	for scanner.Scan() { // internally, it advances token based on sperator
		text := scanner.Text() // token in unicode-char

		// get right of Game #: from the front

		another := strings.Split(text, ": ")

		game := another[1]

		//Grab the game id
		//temporary := another[0]
		//gameID := strings.Split(temporary, "Game ")[1]
		//gameIDInt, _ := strconv.Atoi(gameID)
		red := 0
		blue := 0
		green := 0

		rounds := strings.Split(game, "; ")

		for _, round := range rounds {
			//Get the different colours
			colors := strings.Split(round, ", ")
			for _, color := range colors {
				comp := strings.Split(color, " ")
				color := comp[1]
				temp, _ := strconv.Atoi(comp[0])
				numOfColor := temp

				if color == "red" && numOfColor > red {
					red = numOfColor
				} else if color == "blue" && numOfColor > blue {
					blue = numOfColor
				} else if color == "green" && numOfColor > green {
					green = numOfColor
				}
			}
		}

		//Check

		// if red <= 12 && green <= 13 && blue <= 14 {
		// 	wonGames += gameIDInt
		// 	fmt.Println(gameIDInt)
		// } else {
		// fmt.Println("red: ", red)
		// fmt.Println("blue: ", blue)
		// fmt.Println("green: ", green)
		//}

		totalPowerSet += (red * green * blue)
	}
	fmt.Println(totalPowerSet)

}
