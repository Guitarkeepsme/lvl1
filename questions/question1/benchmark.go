package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkStringBuilder(b *testing.B) {
	var sb strings.Builder
	for n := 0; n < b.N; n++ {
		sb.WriteString("hello")
		sb.WriteString("world")
	}
	fmt.Println(sb.String())
}

func BenchmarkBytesBuffer(b *testing.B) {
	var bb bytes.Buffer
	for n := 0; n < b.N; n++ {
		bb.WriteString("hello")
		bb.WriteString("world")
	}
	fmt.Println(bb.String())
}
