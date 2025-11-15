package mocking

import (
	"fmt"
	"io"
	"iter"
	"time"
)

type Sleeper interface {
	Sleep()
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := range countDownFrom(3) {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, "Go!")
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

type ConfigurableSleeper struct {
	Duration time.Duration
	SleepFn  func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.SleepFn(c.Duration)
}

func countDownFrom(from int) iter.Seq[int] {
	return func(yield func(int) bool) {
		for i := from; i > 0; i-- {
			if !yield(i) {
				return
			}
		}
	}
}

// func main() {
// 	// sleeper := &DefaultSleeper{}
// 	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
// 	Countdown(os.Stdout, sleeper)
// }
