package main

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

const testString = "test"

// бенчмарк конкатенации строк
func BenchmarkConcat(b *testing.B) { //BenchmarkConcat-16    	  832116	    274271 ns/op
	var str string
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		str += testString
	}
	b.StopTimer()
}

func BenchmarkBuffer(b *testing.B) { //BenchmarkBuffer-16    	260115202	         4.092 ns/op
	var buffer bytes.Buffer

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buffer.WriteString(testString)
	}
	b.StopTimer()
}

func BenchmarkCopy(b *testing.B) { //BenchmarkCopy-16      	676972195	         1.768 ns/op
	bs := make([]byte, b.N)
	bl := 0

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		bl += copy(bs[bl:], testString)
	}
	b.StopTimer()
}

func BenchmarkBuild(b *testing.B) { //BenchmarkBuild-16     	47784919	        23.71 ns/op
	var bs strings.Builder

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		fmt.Fprintf(&bs, testString)
	}
	b.StopTimer()
}
