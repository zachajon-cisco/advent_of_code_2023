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

type AllMaps struct {
	seedToSoil          ConversionMap
	soilToFertilizer    ConversionMap
	fertilizerToWater   ConversionMap
	waterToLight        ConversionMap
	lightToTemp         ConversionMap
	tempToHumidity      ConversionMap
	humditityToLocation ConversionMap
}

func (m *AllMaps) FindLocation(value int) int {
	//fmt.Println("Seed: ", value)
	a := m.seedToSoil.hello[value]
	//fmt.Println("Soil: ", a)
	b := m.soilToFertilizer.hello[a]
	if b == 0 {
		b = a
	}
	//fmt.Println("Fertilizer: ", b)
	c := m.fertilizerToWater.hello[b]
	if c == 0 {
		c = b
	}
	//fmt.Println("Water: ", c)
	d := m.waterToLight.hello[c]
	if d == 0 {
		d = c
	}
	//fmt.Println("Light: ", d)
	e := m.lightToTemp.hello[d]
	if e == 0 {
		e = d
	}
	//fmt.Println("Temp: ", e)
	f := m.tempToHumidity.hello[e]
	if f == 0 {
		f = e
	}
	//fmt.Println("Humidity: ", f)
	g := m.humditityToLocation.hello[f]
	if g == f {
		g = f
	}
	fmt.Println("Location: ", g)

	return g

}

type ConversionMap struct {
	hello [5000000000]int
}

func (c *ConversionMap) CheckValue(value int) int {
	if c.hello[value] != 0 {
		return c.hello[value]
	} else {
		return value
	}
}

func (c *ConversionMap) AddNewRange(s []int) {
	destRange := s[0]
	srcRange := s[1]
	rangeLength := s[2]

	for i := 0; i < rangeLength; i++ {
		c.hello[srcRange+i] = destRange + i
	}
}

func main() {

	file, err := os.Open("5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	seeds := []int{}
	onSeedToSoilMap := false
	onSoilToFertilizerMap := false
	onFertilizerToWaterMap := false
	onWaterToLightMap := false
	onLightToTemperatureMap := false
	onTemperatureToHumidityMap := false
	onHumidityToLocationMap := false

	seedToSoil := [][]int{}
	soilToFertilizer := [][]int{}
	fertilizerToWater := [][]int{}
	waterToLight := [][]int{}
	lightToTemp := [][]int{}
	tempToHumidity := [][]int{}
	humditityToLocation := [][]int{}

	for scanner.Scan() { // internally, it advances token based on sperator

		text := scanner.Text()

		// First need to graph seeds
		if strings.Contains(text, "seeds:") {
			//seeds: 79 14 55 13
			//Grab just numbers
			parts := strings.Split(text, ":")
			seeds = convertStringToIntArray(parts[1])
		} else if strings.Contains(text, "seed-to-soil map:") {
			onSeedToSoilMap = true
			onSoilToFertilizerMap = false
			onFertilizerToWaterMap = false
			onWaterToLightMap = false
			onLightToTemperatureMap = false
			onTemperatureToHumidityMap = false
			onHumidityToLocationMap = false
		} else if strings.Contains(text, "soil-to-fertilizer map:") {
			onSeedToSoilMap = false
			onSoilToFertilizerMap = true
			onFertilizerToWaterMap = false
			onWaterToLightMap = false
			onLightToTemperatureMap = false
			onTemperatureToHumidityMap = false
			onHumidityToLocationMap = false
		} else if strings.Contains(text, "fertilizer-to-water map:") {
			onSeedToSoilMap = false
			onSoilToFertilizerMap = false
			onFertilizerToWaterMap = true
			onWaterToLightMap = false
			onLightToTemperatureMap = false
			onTemperatureToHumidityMap = false
			onHumidityToLocationMap = false
		} else if strings.Contains(text, "water-to-light map:") {
			onSeedToSoilMap = false
			onSoilToFertilizerMap = false
			onFertilizerToWaterMap = false
			onWaterToLightMap = true
			onLightToTemperatureMap = false
			onTemperatureToHumidityMap = false
			onHumidityToLocationMap = false
		} else if strings.Contains(text, "light-to-temperature map:") {
			onSeedToSoilMap = false
			onSoilToFertilizerMap = false
			onFertilizerToWaterMap = false
			onWaterToLightMap = false
			onLightToTemperatureMap = true
			onTemperatureToHumidityMap = false
			onHumidityToLocationMap = false
		} else if strings.Contains(text, "temperature-to-humidity map:") {
			onSeedToSoilMap = false
			onSoilToFertilizerMap = false
			onFertilizerToWaterMap = false
			onWaterToLightMap = false
			onLightToTemperatureMap = false
			onTemperatureToHumidityMap = true
			onHumidityToLocationMap = false
		} else if strings.Contains(text, "humidity-to-location map:") {
			onSeedToSoilMap = false
			onSoilToFertilizerMap = false
			onFertilizerToWaterMap = false
			onWaterToLightMap = false
			onLightToTemperatureMap = false
			onTemperatureToHumidityMap = false
			onHumidityToLocationMap = true
		} else if text != "" {
			//Numbers
			if onSeedToSoilMap {
				seedToSoil = append(seedToSoil, convertStringToIntArray(text))
			} else if onSoilToFertilizerMap {
				soilToFertilizer = append(soilToFertilizer, convertStringToIntArray(text))
			} else if onFertilizerToWaterMap {
				fertilizerToWater = append(fertilizerToWater, convertStringToIntArray(text))
			} else if onWaterToLightMap {
				waterToLight = append(waterToLight, convertStringToIntArray(text))
			} else if onLightToTemperatureMap {
				lightToTemp = append(lightToTemp, convertStringToIntArray(text))
			} else if onTemperatureToHumidityMap {
				tempToHumidity = append(tempToHumidity, convertStringToIntArray(text))
			} else if onHumidityToLocationMap {
				humditityToLocation = append(humditityToLocation, convertStringToIntArray(text))
			}

		}

	}
	fmt.Println(seeds)
	fmt.Println("-------------------------")
	fmt.Println(onSeedToSoilMap)
	fmt.Println(onSoilToFertilizerMap)
	fmt.Println(onFertilizerToWaterMap)
	fmt.Println(onWaterToLightMap)
	fmt.Println(onLightToTemperatureMap)
	fmt.Println(onTemperatureToHumidityMap)
	fmt.Println(onHumidityToLocationMap)
	fmt.Println("------------------------")
	fmt.Println(soilToFertilizer)
	fmt.Println(fertilizerToWater)
	fmt.Println(waterToLight)
	fmt.Println(lightToTemp)
	fmt.Println(tempToHumidity)
	fmt.Println(humditityToLocation)

	//1 ) Find the lowest location given the initial seeds
	lowestLocation := 0

	m := AllMaps{
		seedToSoil: ConversionMap{
			hello: [5000000000]int{},
		},
		soilToFertilizer: ConversionMap{
			hello: [5000000000]int{},
		},
		fertilizerToWater: ConversionMap{
			hello: [5000000000]int{},
		},
		waterToLight: ConversionMap{
			hello: [5000000000]int{},
		},
		lightToTemp: ConversionMap{
			hello: [5000000000]int{},
		},
		tempToHumidity: ConversionMap{
			hello: [5000000000]int{},
		},
		humditityToLocation: ConversionMap{
			hello: [5000000000]int{},
		},
	}

	for _, s := range seedToSoil {
		m.seedToSoil.AddNewRange(s)
	}
	for _, s := range soilToFertilizer {
		m.soilToFertilizer.AddNewRange(s)
	}
	for _, s := range fertilizerToWater {
		m.fertilizerToWater.AddNewRange(s)
	}
	for _, s := range waterToLight {
		m.waterToLight.AddNewRange(s)
	}
	for _, s := range lightToTemp {
		m.lightToTemp.AddNewRange(s)
	}
	for _, s := range tempToHumidity {
		m.tempToHumidity.AddNewRange(s)
	}
	for _, s := range humditityToLocation {
		m.humditityToLocation.AddNewRange(s)
	}

	fmt.Println(lowestLocation)
	for _, k := range seeds {
		fmt.Println("Seed: ", k)
		fmt.Println(m.FindLocation(k))
	}
}

func convertStringToIntArray(input string) []int {
	numbers := []int{}
	for _, token := range strings.Split(strings.Trim(input, " "), " ") {
		number, _ := strconv.Atoi(token)
		if !slices.Contains(numbers, number) {
			numbers = append(numbers, number)
		}
	}

	return numbers
}
