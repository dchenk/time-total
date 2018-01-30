// Run this program with one argument: a quoted string containing the work times of each day (month/day at
// beginning of each line).
package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const timeFmt = "15:04"

var outPath = os.Getenv("GOPATH") + "/src/github.com/dchenk/time-total/"

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Supply a string containing the times as an argument to the program!")
		return
	}

	t := os.Args[1]

	f, err := os.Create(outPath + "hours-" + time.Now().Format("02-Jan-06_15-04") + ".csv")
	if err != nil {
		fmt.Printf("could not create file: %s", err)
		return
	}
	defer f.Close()

	enc := csv.NewWriter(f)
	enc.Write([]string{"month/day", "total hrs"})
	defer func() {
		e := enc.Error()
		if e != nil {
			fmt.Printf("error on flushing: %s", e)
		}
	}()
	defer enc.Flush()

	days := strings.Split(t, "\n")

	for _, day := range days {

		ranges := strings.Split(day, " ")

		monthDay := ranges[0]
		ranges = ranges[1:]
		var totHrs float64

		for j := range ranges {
			times := strings.Split(ranges[j], "-")
			if len(times) != 2 {
				panic("len(times) != 2")
			}
			start, err := time.Parse(timeFmt, times[0])
			if err != nil {
				fmt.Printf("could not parse start hr: %s", err)
				return
			}
			end, err := time.Parse(timeFmt, times[1])
			if err != nil {
				fmt.Printf("could not parse start hr: %s", err)
				return
			}
			if end.Before(start) {
				end = end.Add(time.Hour * 12) // set to military time
			}
			totHrs += end.Sub(start).Hours()
		}
		space := ""
		if len(monthDay) == 3 {
			space = " "
		}
		fmt.Printf("(%s)%s %.2f hrs total %s\n", monthDay, space, totHrs, ranges)
		enc.Write([]string{monthDay, strconv.FormatFloat(totHrs, 'f', 2, 64)})

	}

}
