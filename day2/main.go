package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var shapeWeight map[string]int 
	shapeWeight = make(map[string]int)	
	shapeWeight["X"] = 0
	shapeWeight["Y"] = 3
	shapeWeight["Z"] = 6

	shapeWeight["A"] = 1
	shapeWeight["B"] = 2
	shapeWeight["C"] = 3

	fmt.Println(shapeWeight)
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if(len(line) == 0) {
			continue
		}

		chars := strings.Split(line, " ")
		elfWeight := shapeWeight[chars[0]]
		roundWeight := shapeWeight[chars[1]]
		myWeight := 0

		if(roundWeight == 3) {
			myWeight = elfWeight
		} else if ( roundWeight == 0 ) {
			if (elfWeight == 1) {
				myWeight = 3
			} else if(elfWeight == 2) {
				myWeight = 1
			} else if ( elfWeight == 3) {
				myWeight = 2
			}
		} else {
			if (elfWeight == 1) {
				myWeight = 2
			} else if(elfWeight == 2) {
				myWeight = 3
			} else if ( elfWeight == 3) {
				myWeight = 1
			}
		}

		fmt.Println(chars, myWeight, roundWeight)

		sum += roundWeight + myWeight
	}
	
	fmt.Println(sum)
}