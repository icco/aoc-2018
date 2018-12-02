package main

import (
	"bufio"
	"log"
	"os"
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

	for _, a := range boxList {
		for _, b := range boxList {
			c, str := Diff(a, b)
			if c == 1 {
				log.Printf("Diff(%s, %s): %s", a, b, str)
			}
		}
	}
}

// Diff takes two strings and returns the count of letters that differ and
// string they have in common.
func Diff(a, b string) (int, string) {
	ret := ""
	diffCnt := 0
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] == b[i] {
			ret += string(a[i])
		} else {
			diffCnt += 1
		}
	}

	return diffCnt, ret
}
