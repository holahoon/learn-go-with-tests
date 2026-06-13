package concurrency

import (
	"testing"
	"time"
)

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {
	urls := make([]string, 100)

	for i := range urls {
		urls[i] = "a url"
	}

	for range b.N {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}

	// for b.Loop(){
	// 	CheckWebsites(slowStubWebsiteChecker, urls)
	// }
}

// Run go test -bench=.
