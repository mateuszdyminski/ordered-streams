package main

import (
	"github.com/mateuszdyminski/ordered-streams/stream"
	"fmt"
)

func main() {
	stream0 := stream.NewStream([]int{1, 3, 8, 10})
	stream1 := stream.NewStream([]int{1, 4, 5, 6, 12})
	stream2 := stream.NewStream([]int{3, 7, 8, 9})

	for val := range stream.Order(stream0, stream1, stream2) {
		fmt.Printf("%d, ", val)
	}
}
