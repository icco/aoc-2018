package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	frequency := 0
	adjustments := []int{}
	frequencies := map[int]int{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Printf("Parsing error: %+v", err)
		}

		adjustments = append(adjustments, i)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	i := 0
	for {
		frequency = frequency + adjustments[i]
		frequencies[frequency] += 1
		i = (i + 1) % len(adjustments)

		if frequencies[frequency] >= 2 {
			log.Printf("Seen twice: %d", frequency)
			os.Exit(0)
		}
	}
}
