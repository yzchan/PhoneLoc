package ip138

import (
	"math/rand"
	"testing"
	"time"
)

type Target struct {
	Phone int
	City  string
	Sp    string
}

func BenchmarkPhoneLoc(b *testing.B) {
	b.StopTimer()

	rand.Seed(time.Now().UnixNano())
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		_, _, _ = PhoneLoc("1300006")
	}
}
