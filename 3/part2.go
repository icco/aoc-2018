package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"regexp"
	"sort"
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
		for x := b; x <= w; x++ {
			xs = append(xs, x)
		}

		ys := []int{}
		for y := c; y <= h; y++ {
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

func MakeOverlap(cuts []*Cut, maxWidth, maxHeight int) [][]int {
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

	return veryLarge
}

func FindCut(cuts []*Cut, x, y int) *Cut {
	validX := []*Cut{}

	for _, c := range cuts {
		i := sort.Search(len(c.Xs), func(i int) bool { return c.Xs[i] == x })
		if i < len(c.Xs) {
			// x is present
			validX = append(validX, c)
		}
	}

	for _, c := range validX {
		i := sort.Search(len(c.Ys), func(i int) bool { return c.Ys[i] == y })
		if i < len(c.Ys) {
			// y is present
			return c
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

	overlaps := MakeOverlap(cuts, maxWidth, maxHeight)

	for x, row := range overlaps {
		for y, c := range row {
			if c < 3 {
				cut := FindCut(cuts, x, y)
				if cut != nil {
					log.Printf("c == %d: \t %+v", c, cut)
				}
			}
		}
	}
}
