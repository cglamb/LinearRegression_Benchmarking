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

	data2 := []stats.Coordinate{
		{10, 9.14},
		{8, 8.14},
		{13, 8.74},
		{9, 8.77},
		{11, 9.26},
		{14, 8.10},
		{6, 6.13},
		{4, 3.10},
		{12, 9.13},
		{7, 7.26},
		{5, 4.74},
	}

	data3 := []stats.Coordinate{
		{10, 7.46},
		{8, 6.77},
		{13, 12.74},
		{9, 7.11},
		{11, 7.81},
		{14, 8.84},
		{6, 6.08},
		{4, 5.39},
		{12, 8.15},
		{7, 6.42},
		{5, 5.73},
	}

	data4 := []stats.Coordinate{
		{8, 6.58},
		{8, 5.76},
		{8, 7.71},
		{8, 8.84},
		{8, 8.47},
		{8, 7.04},
		{8, 5.25},
		{19, 12.5},
		{8, 5.56},
		{8, 7.91},
		{8, 6.89},
	}

	//my experimental design invovled running 10 times for PY and R to determine runtime
	//i recongize that go benchmarking has built in interations...which are commeneted out at the bottom of this code
	regress(data1)
	regress(data2)
	regress(data3)
	regress(data4)

	// for i := 0; i < b.N; i++ {
	// 	regress(data1)
	// 	regress(data2)
	// 	regress(data3)
	// 	regress(data4)
	// }
}
