package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	lines, err := datafile.GetStrings("votes.txt")
	if err != nil {
		log.Fatal(err)
	}

	counts := make(map[string]int)

	for _, line := range lines {
		counts[line]++
	}

	for name, votes := range counts {
		fmt.Printf("%s : %d\n", name, votes)
	}
}
