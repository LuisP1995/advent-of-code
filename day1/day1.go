package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// LoadFile should load data from a file
func LoadFile(fileName string, container map[int]interface{}) error {
	var void interface{}
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Could not open file")
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		data := scanner.Text()
		num, _ := strconv.Atoi(data)
		container[num] = void
	}

	return nil
}

// Calc2020 Add two numbers together from the map together
func Calc2020(firstVal int) int {
	target := 2020
	return (target - firstVal)
}

// Problem1 part 1 of 2 for advent of code
func Problem1(dataSet map[int]interface{}) int {
	code := 0
	for key := range dataSet {
		result := Calc2020(key)
		if _, ok := dataSet[result]; ok {
			fmt.Println("Value found")
			code = (key * result)
			return code
		}

	}
	return -1
}

func main() {
	file := "numbers"

	dataSet := make(map[int]interface{})

	LoadFile(file, dataSet)

	code := Problem1(dataSet)

	fmt.Println(code)
}
