package main

import "testing"

func BenchmarkProccessData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProccessData()
	}
}

func BenchmarkProccessData2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ProccessData2()
	}
}
