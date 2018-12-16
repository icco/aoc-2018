package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"time"
)

type Entry struct {
	Guard  int
	Time   time.Time
	Action string
}

type ByTime []*Entry

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByTime) Less(i, j int) bool { return a[i].Time.Before(a[j].Time) }

const WAKE = "wake"
const SLEEP = "sleep"
const SWITCH = "switch"

func Parse(line string) *Entry {
	// [1518-11-05 00:03] Guard #99 begins shift
	// [1518-11-05 00:45] falls asleep
	// [1518-11-05 00:55] wakes up
	var re = regexp.MustCompile(`\[(\d{4}-\d{2}-\d{2} \d{2}:\d{2})\] (.+)`)
	parts := re.FindStringSubmatch(line)
	e := &Entry{}
	t, err := time.Parse("2006-01-02 15:04", parts[1])
	if err != nil {
		log.Panicf("Time parse: %+v", err)
	}
	e.Time = t

	switch parts[2] {
	case "wakes up":
		e.Action = WAKE
	case "falls asleep":
		e.Action = SLEEP
	default:
		p := regexp.MustCompile(`Guard #(\d+) begins shift`).FindStringSubmatch(parts[2])
		i, _ := strconv.Atoi(p[1])
		e.Guard = i
		e.Action = SWITCH
	}

	return e
}

func FillGuard(entries []*Entry) {
	currentGuard := 0
	for _, e := range entries {
		switch e.Action {
		case WAKE:
			e.Guard = currentGuard
		case SLEEP:
			e.Guard = currentGuard
		case SWITCH:
			currentGuard = e.Guard
		}
	}
}

func main() {
	actions := []*Entry{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		actions = append(actions, Parse(line))
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}

	sort.Sort(ByTime(actions))
	FillGuard(actions)

	var lastSleep *Entry
	guardSleep := map[int][60]int{}
	guardSleepTotal := map[int]int{}
	for _, e := range actions {
		log.Printf("%+v", e)
		if e.Action == SLEEP {
			lastSleep = e
		}

		if e.Action == WAKE {
			for i := lastSleep.Time.Minute(); i < e.Time.Minute(); i++ {
				minutes := guardSleep[e.Guard]
				minutes[i] += 1
				guardSleep[e.Guard] = minutes
				guardSleepTotal[e.Guard] += 1
			}
		}
	}

	max := 0
	highestGuard := 0
	for g, t := range guardSleep {
		log.Printf("%4d\t(%d): \t %+v", g, guardSleepTotal[g], t)
		if guardSleepTotal[g] > max {
			max = guardSleepTotal[g]
			highestGuard = g
		}
	}

	max = 0
	highestMin := 0
	for m, c := range guardSleep[highestGuard] {

		if c > max {
			highestMin = m
			max = c
		}
	}

	log.Printf("%d * %d = %d", highestMin, highestGuard, highestMin*highestGuard)
}
