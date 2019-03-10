package main

import (
	"fmt"
	"math"
)

func signalProducer(in chan float64) {
	for i := 0; i < 360; i++ {
		in <- math.Sin(degreeToRads(float64(i)))
	}
	close(in)
}

func metricsProducer(in chan float64, out chan string) {
	for v := range in {
		out <- fmt.Sprintf("{'temperature': '%.3f'}", v)
	}
	close(out)
}

func metricsPublisher(metrics chan string) {
	for v := range metrics {
		fmt.Println(v)
	}
}

func degreeToRads(degree float64) float64 {
	return degree * (math.Pi / 180)
}

func main() {
	values := make(chan float64)
	metrics := make(chan string)

	go signalProducer(values)
	go metricsProducer(values, metrics)
	metricsPublisher(metrics)
}
