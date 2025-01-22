package main

import (
	"fmt"
	"syscall"
	"time"
)

func main() {
	beforeSleep := time.Now()

	if !sleep() {
		return
	}

	afterSleep := time.Now()

	duration := afterSleep.Sub(beforeSleep)

	fmt.Printf("System went to sleep at: %v\n", beforeSleep)
	fmt.Printf("System woke up at: %v\n", afterSleep)
	fmt.Printf("Elapsed time: %v\n", duration)
}

func sleep() bool {
	powrProf := syscall.NewLazyDLL("powrprof.dll")
	setSuspendState := powrProf.NewProc("SetSuspendState")

	// 0 = Sleep, 0 = No Hibernate, 0 = Allow Wake Events
	ret, _, err := setSuspendState.Call(0, 0, 0)
	if ret == 0 {
		fmt.Printf("Failed to put the system to sleep: %v\n", err)
		return false;
	}

	return true
}

