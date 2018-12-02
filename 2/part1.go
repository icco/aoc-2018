package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	boxList := []string{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		boxList = append(boxList, line)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	twoTimes := 0
	threeTimes := 0

	for _, b := range boxList {
		parts := strings.Split(b, "")
		counts := map[string]int{}
		for _, let := range parts {
			counts[let] += 1
		}

		cnts := map[int]int{}
		for _, cnt := range counts {
			cnts[cnt] += 1
		}

		if cnts[2] >= 1 {
			twoTimes += 1
		}

		if cnts[3] >= 1 {
			threeTimes += 1
		}

		log.Printf("%s: 2x: %d 3x: %d", b, twoTimes, threeTimes)
	}

	log.Printf("Final math: %d", twoTimes*threeTimes)
}
