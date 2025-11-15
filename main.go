package main

import (
	"os"
	"time"

	mocking "github.com/holahoon/learn-go-with-tests/9_mocking"
)

func main() {
	// sleeper := &mocking.DefaultSleeper{}
	duration := 1 * time.Second
	sleep := time.Sleep
	sleeper := &mocking.ConfigurableSleeper{Duration: duration, SleepFn: sleep}
	mocking.Countdown(os.Stdout, sleeper)
}
