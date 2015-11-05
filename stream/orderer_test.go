package stream

import (
	"sort"
	"testing"
)

func TestOrderWithNoStreams(t *testing.T) {
	out := Order()

	counter := 0
	for range out {
		counter++
	}

	if counter != 0 {
		t.Error("Channel should be closed without emmitting any vals!")
	}
}

func TestOrderSingleStream(t *testing.T) {
	streamElems := []int{1, 2, 3}
	expectedNoElems := 3
	s1 := NewStream(streamElems)

	out := Order(s1)

	result := make([]int, 0, expectedNoElems)
	for v := range out {
		result = append(result, v)
	}

	if len(result) != expectedNoElems {
		t.Errorf("Order should return proper number of elements! Expected %d. Got: %d", expectedNoElems, len(result))
	}

	for i, v := range streamElems {
		if result[i] != v {
			t.Errorf("Wrong order of elements! Expected %d. Got: %d", v, result[i])
		}
	}
}

func TestOrderManyStreams(t *testing.T) {
	streamElems1 := []int{1, 3, 5, 7}
	streamElems2 := []int{1, 3, 6}
	allSortedElems := append(streamElems1, streamElems2...)
	sort.Ints(allSortedElems)
	expectedNoElems := len(allSortedElems)

	out := Order(NewStream(streamElems1), NewStream(streamElems2))

	result := make([]int, 0, expectedNoElems)
	for v := range out {
		result = append(result, v)
	}

	if len(result) != expectedNoElems {
		t.Errorf("Order should return proper number of elements! Expected %d. Got: %d", expectedNoElems, len(result))
	}

	for i, v := range allSortedElems {
		if result[i] != v {
			t.Errorf("Wrong order of elements! Expected %d. Got: %d", v, result[i])
		}
	}
}
