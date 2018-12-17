package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)
	chars := []string{}
	for scanner.Scan() {
		chars = append(chars, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	for {
		removed := false
		for i := 0; i < len(chars)-1; i++ {
			if isPair(chars[i], chars[i+1]) {
				chars = append(chars[:i], chars[i+2:]...)
				removed = true
				break
			}
		}

		if !removed {
			log.Printf("Complete! %d", len(chars))
			return
		}
	}
}

func isPair(a, b string) bool {
	aLower := strings.ToLower(a)
	bLower := strings.ToLower(b)

	return (a != b) && (a == bLower || aLower == b)
}
