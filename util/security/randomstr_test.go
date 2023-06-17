package security

import (
	"fmt"
	"testing"
)

func TestRandStr(t *testing.T) {
	fmt.Println(RandStr(10))
}

func BenchmarkRandStr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandStr(10)
	}
}
