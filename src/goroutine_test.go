package main

import (
	"testing"
	// "time"
)

func BenchmarkPrint1(b *testing.B) {
	print1()
}

func BenchmarkGoPrint1(b *testing.B) {
	goPrint1()
}

func BenchmarkPrint2(b *testing.B) {
	print2()
}

func BenchmarkGoPrint2(b *testing.B) {
	goPrint2()
	// time.Sleep(1 * time.Microsecond)
}