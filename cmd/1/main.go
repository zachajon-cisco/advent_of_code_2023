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

	file, err := os.Open("1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	total := 0

	for scanner.Scan() { // internally, it advances token based on sperator
		text := scanner.Text() // token in unicode-char
		totalNumString := ""

		//Need to preprocess

		numbersMap := map[string]string{
			"eightwo":   "82",
			"eighthree": "83",
			"oneight":   "18",
			"nineight":  "98",
			"sevenine":  "79",
			"twone":     "21",
		}

		for i, j := range numbersMap {
			if strings.Contains(text, i) {
				text = strings.Replace(text, i, j, 10)
			}
		}

		othersMap := map[string]string{
			"one":   "1",
			"two":   "2",
			"three": "3",
			"four":  "4",
			"five":  "5",
			"six":   "6",
			"seven": "7",
			"eight": "8",
			"nine":  "9",
		}

		for i, j := range othersMap {
			if strings.Contains(text, i) {
				text = strings.Replace(text, i, j, 10)
			}
		}

		// Need to grab only the
		for i := range text {
			v := text[i]

			if _, err := strconv.Atoi(string(v)); err == nil {
				if len(totalNumString) == 0 {
					totalNumString = string(v)
				} else {
					totalNumString = string(totalNumString[0]) + string(v)
				}

			}
		}
		//special case
		if len(totalNumString) == 1 {
			totalNumString = string(totalNumString[0]) + string(totalNumString[0])
		}
		fmt.Println(totalNumString)
		num, _ := strconv.Atoi(totalNumString)
		total += num
	}

	fmt.Println(total)

}
