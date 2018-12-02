package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func main() {
	frequency := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Splits on lines
		line := scanner.Text()
		i, err := strconv.Atoi(line)

		if err != nil {
			log.Printf("Parsing error: %+v", err)
		}

		//log.Printf("Parsed %+v", i)
		frequency = i + frequency
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	log.Printf("Final Frequency: %d", frequency)
}
