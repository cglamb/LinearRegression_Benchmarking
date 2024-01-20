package main

//package executedin terminal via go test -bench=/ command

import (
	"testing"

	"github.com/montanaflynn/stats"
)

func Benchmark_linear_regression(b *testing.B) {

	data1 := []stats.Coordinate{
		{10, 8.04},
		{8, 6.95},
		{13, 7.58},
		{9, 8.81},
		{11, 8.33},
		{14, 9.96},
		{6, 7.24},
		{4, 4.26},
		{12, 10.84},
		{7, 4.82},
		{5, 5.68},
	}

	//my experimental design invovled running 10 times for PY and R to determine runtime
	//i recongize that go benchmarking has built in interations...which are commeneted out at the bottom of this code
	regress(data1)

	//for i := 0; i < b.N; i++ {
	//		regress(data1)
	//}
}
