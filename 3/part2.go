package main

import (
	"bufio"
	"fmt"
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
	Xs         []int
	Ys         []int
	Overlap    bool
}

var re = regexp.MustCompile(`#(\d+) @ (\d+),(\d+): (\d+)x(\d+)`)

func Parse(line string) *Cut {
	parts := re.FindStringSubmatch(line)

	if len(parts) >= 5 {
		a, _ := strconv.Atoi(parts[1])
		b, _ := strconv.Atoi(parts[2])
		c, _ := strconv.Atoi(parts[3])
		w, _ := strconv.Atoi(parts[4])
		h, _ := strconv.Atoi(parts[5])

		xs := []int{}
		for x := b; x <= b+w; x++ {
			xs = append(xs, x)
		}

		ys := []int{}
		for y := c; y <= c+h; y++ {
			ys = append(ys, y)
		}

		return &Cut{
			ID:         a,
			LeftOffset: b,
			TopOffset:  c,
			Width:      w,
			Height:     h,
			Xs:         xs,
			Ys:         ys,
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

	coords := map[string][]*Cut{}
	for _, c := range cuts {
		for _, x := range c.Xs {
			for _, y := range c.Ys {
				coord := fmt.Sprintf("%d,%d", x, y)
				coords[coord] = append(coords[coord], c)
			}
		}
	}

	for _, cuts := range coords {
		for _, c := range cuts {
			c.Overlap = len(cuts) > 1
		}
	}

	for _, c := range cuts {
		if !c.Overlap {
			log.Printf("-- %+v", c)
		}
	}
}
