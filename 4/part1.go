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

	for _, e := range actions {
		log.Printf("%+v", e)
	}
}
