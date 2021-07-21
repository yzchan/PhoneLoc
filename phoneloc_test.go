package phoneloc

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkFind(b *testing.B) {
	b.StopTimer()

	loc, err := NewParser("phone.dat")
	if err != nil {
		panic(err)
	}

	rand.Seed(time.Now().UnixNano())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _ = loc.Find(1300006)
	}
}
