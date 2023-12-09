package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func main() {

	file, err := os.Open("3.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	da := [140][]rune{}
	counter := 0

	//Need to read in file and turn into a double array

	for scanner.Scan() { // internally, it advances token based on sperator

		text := scanner.Text() // token in unicode-char

		da[counter] = []rune(text)
		counter++
	}

	da_bool := [140][140]bool{}
	da_part_num := [140][140]int{}

	for x := 0; x < 140; x++ {
		for y := 0; y < 140; y++ {
			da_bool[x][y] = false
		}
	}

	for x := 0; x < 140; x++ {
		for y := 0; y < 140; y++ {
			//(x,y)
			//Need to check all around
			//Good to go! Do the checks
			if isSymbol(x-1, y-1, da) {
				da_bool[x][y] = true
			}
			if isSymbol(x-1, y, da) {
				da_bool[x][y] = true
			}
			if isSymbol(x-1, y+1, da) {
				da_bool[x][y] = true
			}
			if isSymbol(x+1, y-1, da) {
				da_bool[x][y] = true
			}
			if isSymbol(x+1, y, da) {
				da_bool[x][y] = true
			}
			if isSymbol(x+1, y+1, da) {
				da_bool[x][y] = true
			}
			if isSymbol(x, y+1, da) {
				da_bool[x][y] = true
			}
			if isSymbol(x, y-1, da) {
				da_bool[x][y] = true
			}
		}
	}

	//now need to go around and figure out the numbers

	total := 0
	for x := 0; x < 140; x++ {
		numString := ""
		leftNum := false
		for y := 0; y < 140; y++ {
			temp := string(da[x][y])
			fmt.Println(temp)
			//fmt.Println("temp: ", temp)
			//fmt.Println(numString)
			//fmt.Println("isBool: ", da_bool[x][y])
			//fmt.Println("leftNum: ", leftNum)
			if _, err := strconv.Atoi(temp); err == nil {
				//TODO: May need to test if this actually makes sense
				// Number
				numString += temp
				if da_bool[x][y] {
					leftNum = true
				}

				//Edge case where number is at the end
				if y == 139 && leftNum {
					num, _ := strconv.Atoi(numString)
					total += num
					lengthToDelete := len(numString)
					for i := 0; i < lengthToDelete; i++ {
						//fmt.Println(x, ",", y)
						//Then update da_part_num
						da_part_num[x][y-i-1] = num
					}
					numString = ""
					//fmt.Println(num)
					leftNum = false
				}
			} else {
				// Either period or symbol
				if leftNum {
					// Check if numbers should be added
					//fmt.Println(numString)
					num, _ := strconv.Atoi(numString)
					total += num
					//fmt.Println("Total: ", total)
					lengthToDelete := len(numString)
					for i := 0; i < lengthToDelete; i++ {
						//fmt.Println(x, ",", y)
						//Then update da_part_num
						da_part_num[x][y-i-1] = num
					}
					numString = ""
					//fmt.Println(num)
					leftNum = false
				} else {
					// Reset number otherwise
					numString = ""
				}
			}
		}
	}

	fmt.Println(total)

	//Need to go back through the array and pull

	total = 0
	for x := 0; x < 140; x++ {
		for y := 0; y < 140; y++ {
			partNumsFound := []int{}
			if isGear(x, y, da) {
				fmt.Println("I'm a gear")
				for i := -1; i <= 1; i++ {
					for j := -1; j <= 1; j++ {
						fmt.Println("(", x+i, ",", y+j, ") = ", da_part_num[x+i][y+j])
					}
				}
				//This is a gear and need to check around
				if isPartNum(x-1, y-1, da_part_num) {
					if !slices.Contains(partNumsFound, da_part_num[x-1][y-1]) {
						partNumsFound = append(partNumsFound, da_part_num[x-1][y-1])
					}
				}

				if isPartNum(x-1, y, da_part_num) {
					if !slices.Contains(partNumsFound, da_part_num[x-1][y]) {
						partNumsFound = append(partNumsFound, da_part_num[x-1][y])
					}
				}
				if isPartNum(x-1, y+1, da_part_num) {
					if !slices.Contains(partNumsFound, da_part_num[x-1][y+1]) {
						partNumsFound = append(partNumsFound, da_part_num[x-1][y+1])
					}
				}
				if isPartNum(x+1, y-1, da_part_num) {
					if !slices.Contains(partNumsFound, da_part_num[x+1][y-1]) {
						partNumsFound = append(partNumsFound, da_part_num[x+1][y-1])
					}
				}
				if isPartNum(x+1, y, da_part_num) {
					if !slices.Contains(partNumsFound, da_part_num[x+1][y]) {
						partNumsFound = append(partNumsFound, da_part_num[x+1][y])
					}
				}
				if isPartNum(x+1, y+1, da_part_num) {
					if !slices.Contains(partNumsFound, da_part_num[x+1][y+1]) {
						partNumsFound = append(partNumsFound, da_part_num[x+1][y+1])
					}
				}
				if isPartNum(x, y+1, da_part_num) {
					if !slices.Contains(partNumsFound, da_part_num[x][y+1]) {
						partNumsFound = append(partNumsFound, da_part_num[x][y+1])
					}
				}
				if isPartNum(x, y-1, da_part_num) {
					if !slices.Contains(partNumsFound, da_part_num[x][y-1]) {
						partNumsFound = append(partNumsFound, da_part_num[x][y-1])
					}
				}

				fmt.Println(partNumsFound)

				if len(partNumsFound) == 2 {
					total += partNumsFound[0] * partNumsFound[1]
				}
			}
		}
	}

	fmt.Println(total)
}

func isSymbol(x int, y int, da [140][]rune) bool {
	if x < 0 || y < 0 || x >= 140 || y >= 140 {
		return false
	} else {
		b := string(da[x][y])
		nonSymbols := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "."}

		if slices.Contains(nonSymbols, b) {
			return false
		} else {
			return true
		}
	}
}

func isGear(x int, y int, da [140][]rune) bool {
	if x < 0 || y < 0 || x >= 140 || y >= 140 {
		return false
	} else {
		b := string(da[x][y])
		nonSymbols := []string{"*"}

		if slices.Contains(nonSymbols, b) {
			return true
		} else {
			return false
		}
	}
}

func isNumber(x int, y int, da [140][]rune) bool {
	if x < 0 || y < 0 || x >= 140 || y >= 140 {
		return false
	} else {
		b := string(da[x][y])
		nonSymbols := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

		if slices.Contains(nonSymbols, b) {
			return true
		} else {
			return false
		}
	}
}

func isPartNum(x int, y int, da_part_num [140][140]int) bool {
	if x < 0 || y < 0 || x >= 140 || y >= 140 {
		return false
	} else {
		if da_part_num[x][y] != 0 {
			return true
		} else {
			return false
		}
	}
}
