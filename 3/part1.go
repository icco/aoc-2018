package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Cut struct {
	ID         int
	LeftOffset int
	TopOffset  int
	Width      int
	Height     int
}

var re = regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

func Parse(line string) *Cut {
	parts := re.FindStringSubmatch(line)

	if len(parts) >= 5 {
		a, _ := strconv.Atoi(parts[1])
		b, _ := strconv.Atoi(parts[2])
		c, _ := strconv.Atoi(parts[3])
		d, _ := strconv.Atoi(parts[4])
		e, _ := strconv.Atoi(parts[5])

		return &Cut{
			ID:         a,
			LeftOffset: b,
			TopOffset:  c,
			Width:      d,
			Height:     e,
		}
	}

	return nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		c := Parse(line)
		log.Printf("%+v", c)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

}
