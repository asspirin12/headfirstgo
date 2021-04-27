package datafile

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetFloats(filName string) ([]float64, error) {

	var numbers []float64

	file, err := os.Open(filName)
	if err != nil {
		log.Fatal(err)
	}
	
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)

		if err != nil {
			return numbers, err
		}
		
		numbers = append(numbers, number)
	}

	err = file.Close()
	if err != nil {
		return numbers, err
	}

	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
		
	return numbers, nil
}

func GetFloatsFromCL() {
	arguments := os.Args[1:]

	var sum float64 = 0

	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		sum += number
	}

	sampleCount := float64(len(arguments))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}