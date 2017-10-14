package main

import (
	"fmt"
	"time"
)

var Beginning time.Time = time.Unix(1508434200, 0) //10-19-2017 1:30 PM

func StartCountdown() {
	for true {
		now := time.Now()
		hours := int(Beginning.Sub(now).Truncate(time.Hour).Hours())

		now = now.Add(time.Duration(hours) * time.Hour)
		minutes := int(Beginning.Sub(now).Truncate(time.Minute).Minutes())

		now = now.Add(time.Duration(minutes) * time.Minute)
		seconds := int(Beginning.Sub(now).Truncate(time.Second).Seconds())

		days := int(hours) / 24
		hours = int(hours) - (days * 24)

		fmt.Printf(
			"It will begin in %d %s %d %s %d %s %d %s\n",
			days, DurationTxt("Day", days),
			hours, DurationTxt("Hour", int(hours)),
			minutes, DurationTxt("Minute", int(minutes)),
			seconds, DurationTxt("Second", int(seconds)),
		)

		time.Sleep(time.Second)
	}
}

func DurationTxt(s string, i int) string {
	if i != 1 {
		return s + "s"
	}

	return s
}
