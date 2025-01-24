package main

import (
	"fmt"
	"syscall"
	"time"

	write "github.com/RonDeppe/sleeper/internal/io"
)

var displayFormat string = "15:04:05"

func main() {
	write.CreateDB()

	beforeSleep := time.Now()

	if !sleep() {
		return
	}

	afterSleep := time.Now()

	duration := afterSleep.Sub(beforeSleep)

	compactDuration, displayDuration := formatDuration(duration)

	displayTimes(beforeSleep, afterSleep, displayDuration)

	write.AddRecord(beforeSleep, afterSleep, compactDuration)
}

func sleep() bool {
	powrProf := syscall.NewLazyDLL("powrprof.dll")
	setSuspendState := powrProf.NewProc("SetSuspendState")

	// 0 = Sleep, 0 = No Hibernate, 0 = Allow Wake Events
	ret, _, err := setSuspendState.Call(0, 0, 0)
	if ret == 0 {
		fmt.Printf("Failed to put the system to sleep: %v\n", err)
		return false
	}

	return true
}

func formatDuration(duration time.Duration) (string, string) {
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	compact := fmt.Sprintf("%02d%02d%02d", hours, minutes, seconds)
	display := fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)

	return compact, display
}

func displayTimes(before time.Time, after time.Time, duration string) {
	fmt.Printf("System went to sleep at: %v\n", before.Format(displayFormat))
	fmt.Printf("System woke up at: %v\n", after.Format(displayFormat))
	fmt.Printf("Elapsed time: %s\n", duration)
}
