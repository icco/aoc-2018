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

	charCopy := make([]string, len(chars))
	copy(charCopy, chars)

	alphabet := strings.Split("abcdefghijklmnopqrstuvwxyz", "")
	for _, r := range alphabet {
		chars := charCopy
		InfiniStrip(chars, r)
	}
}

func InfiniStrip(chars []string, r string) int {
	oldLen := len(chars)
	for {
		chars = Strip(chars, r)

		if len(chars) == oldLen {
			log.Printf("Completed %s : %d", r, len(chars))
			return len(chars)
		}
		oldLen = len(chars)
	}
}

func Strip(chars []string, r string) []string {
	for i := 0; i < len(chars); i++ {
		if strings.ToLower(chars[i]) == r {
			chars = append(chars[:i], chars[i+1:]...)
			return chars
		}
	}

	for i := 0; i < len(chars)-1; i++ {
		if isPair(chars[i], chars[i+1]) {
			chars = append(chars[:i], chars[i+2:]...)
			return chars
		}
	}

	return chars
}

func isPair(a, b string) bool {
	aLower := strings.ToLower(a)
	bLower := strings.ToLower(b)

	return (a != b) && (a == bLower || aLower == b)
}
