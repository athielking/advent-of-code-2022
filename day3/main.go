package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if(len(line) == 0) {
			continue
		}

		half := len(line)/2
		firstCompartment := line[0:half]
		secondCompartment := line[half:len(line)]

		fmt.Println(firstCompartment, " :: ", secondCompartment)

		set1 := make(map[string]bool)
		set2 := make(map[string]bool)

		var common string

		for i := 0; i<len(firstCompartment); i++ {
			item1 := string(firstCompartment[i])
			item2 := string(secondCompartment[i])

			_, exists := set1[item1]
			if (!exists) {
				set1[item1] = true
			}
			
			_, exists = set2[item2]
			if (!exists) {
				set2[item2] = true			
			}

			_, exists = set1[item2]
			if (exists) {
				common = item2;
				break
			}

			_, exists = set2[item1]
			if (exists) {
				common = item1
				break
			}
		}

		fmt.Println(common, getWeight(common))
		sum += getWeight(common)
	}
	
	fmt.Println(sum)
}

func getWeight(s string) int {
	intVal := int(s[0])
	fmt.Println(intVal)

	if intVal >= 97 && intVal <= 122 {
		 intVal -= 97
		 intVal += 1
		 return intVal
	} else if intVal >= 65 && intVal <= 90 {
		intVal -= 65
		intVal += 27
		return intVal
	}

	return 0
}