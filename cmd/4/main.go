package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("4.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	totalPoints := 0
	totalCopies := [193]int{}
	for i, _ := range totalCopies {
		totalCopies[i] = 1
	}
	currentCard := 0

	for scanner.Scan() { // internally, it advances token based on sperator

		text := scanner.Text()

		// Let's grab the input

		a := strings.Split(text, ":")
		//cardNum := a[0]

		b := strings.Split(a[1], "|")
		hand := convertStringToIntArray(b[1])
		winning := convertStringToIntArray(b[0])

		//fmt.Println(hand)
		//fmt.Println(winning)

		//fmt.Println(hand)
		//fmt.Println(winning)

		totalPoints = getNumOfWinningNumbers(hand, winning)

		for i := 0; i < totalPoints; i++ {
			if currentCard+i < 200 {
				totalCopies[currentCard+i+1] += 1 * totalCopies[currentCard]
			}
		}
		fmt.Println(totalCopies)
		currentCard++

	}

	//Total up array
	totalCards := 0
	for _, i := range totalCopies {
		totalCards += i
	}

	fmt.Println(totalCards)
	fmt.Println(len(totalCopies))
}

func getNumOfWinningNumbers(hand []int, winning []int) int {
	points := 0
	for _, num := range hand {
		//Go through each number in the hand
		if slices.Contains(winning, num) {
			if points == 0 {
				points++
			} else {
				points++
			}
		}
	}
	return points
}

func convertStringToIntArray(input string) []int {
	numbers := []int{}
	for _, token := range strings.Split(strings.Trim(input, " "), " ") {
		number, _ := strconv.Atoi(token)
		if !slices.Contains(numbers, number) && number != 0 {
			numbers = append(numbers, number)
		}
	}

	return numbers
}
