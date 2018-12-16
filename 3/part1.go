package main

import (
	"bufio"
	"log"
	"math"
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
	maxHeight := 0
	maxWidth := 0
	cuts := []*Cut{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		c := Parse(line)
		cuts = append(cuts, c)

		maxHeight = int(math.Max(float64(maxHeight), float64(c.TopOffset+c.Height)))
		maxWidth = int(math.Max(float64(maxWidth), float64(c.LeftOffset+c.Width)))
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	// x, y
	veryLarge := make([][]int, maxWidth)
	for i := 0; i < maxWidth; i++ {
		veryLarge[i] = make([]int, maxHeight)
	}

	for _, c := range cuts {
		for x := 0; x < c.Width; x++ {
			for y := 0; y < c.Height; y++ {
				xo := x + c.LeftOffset
				yo := y + c.TopOffset
				veryLarge[xo][yo] += 1
			}
		}
	}

	cnt := 0
	for _, r := range veryLarge {
		for _, c := range r {
			if c > 1 {
				cnt++
			}
		}
	}

	log.Printf("Overlapping squares: %d", cnt)
}
