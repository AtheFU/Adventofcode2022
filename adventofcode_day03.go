package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Read file line by line
	readFile, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
	}
	inventory := bufio.NewScanner(readFile)
	inventory.Split(bufio.ScanLines)

	j := 0
	total := 0

	for inventory.Scan() {

		rucksack := inventory.Text()
		rucksackCount := len(rucksack)
		compartmentCount := len(rucksack) / 2

		priority := findItem(rucksack)
		total += priority

		fmt.Printf("%d (%d): %s | %s > %d %d \n", j, rucksackCount, rucksack[:compartmentCount], rucksack[compartmentCount:rucksackCount], priority, total)

		j++
	}
	readFile.Close()

	fmt.Printf("%d Insgesamt", total)
}

func findItem(rucksack string) int {

	priority := 0
	rucksackCount := len(rucksack)
	compartmentCount := rucksackCount / 2

	for i := 0; i < compartmentCount; i++ {
		item := rucksack[i:(i + 1)]

		if strings.Contains(rucksack[compartmentCount:rucksackCount], item) {
			chars := []rune(item)
			itemPrio := int(chars[0])

			if itemPrio > 97 {
				priority = (itemPrio - 96)
			} else {
				priority = (itemPrio - 38)
			}

			return priority
		}
	}

	return 0

}
