package stringconcat

import (
	"strings"
	"testing"
)

func BenchmarkConcatPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		s := ""
		for j := 0; j < 100; j++ {
			s += "a"
		}
		_ = s
	}
}

func BenchmarkConcatBuilderPrealloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var builder strings.Builder
		builder.Grow(100)

		for j := 0; j < 100; j++ {
			builder.WriteString("a")
		}
		_ = builder.String()
	}
}
