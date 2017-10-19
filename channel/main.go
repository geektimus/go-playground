package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)
	s := rand.NewSource(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		go getRandom(s, c)
	}

	for n := range c {
		fmt.Println(n)
	}
}

func getRandom(s rand.Source, c chan int) {

	r := rand.New(s)
	for {
		c <- r.Intn(100)
	}
}
