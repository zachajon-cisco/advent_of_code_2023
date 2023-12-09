package main

import "fmt"

func main() {
	array := [10][10]int{}

	counter := 0
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			array[x][y] = counter
			counter++
		}
	}

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			fmt.Println(array[y][x], " ")
		}
	}
}
