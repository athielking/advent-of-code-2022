package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(f)

	var maxCalories = [3]int64 {0, 0, 0}
	
	var elfCalories int64 = 0
	var elfIndex int = 1;

	for scanner.Scan() {
		line := scanner.Text()

		if len(line) == 0 {

			for i := 0; i<len(maxCalories); i++ {
				max := maxCalories[i]

				if(elfCalories > max) {
					maxCalories[i] = elfCalories
					elfCalories = max
				}				
			}

			elfCalories = 0
			elfIndex++
		} else {
			i64, _ := strconv.ParseInt(line, 10, 64)
			elfCalories += i64
		}
	}

	var sum int64 = 0
	for i := 0; i<len(maxCalories); i++ {
		sum += maxCalories[i]
	}
	
	fmt.Println(sum)
}