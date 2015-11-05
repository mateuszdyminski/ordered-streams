package stream

import "testing"

func TestEmptyStream(t *testing.T) {
	s := NewStream([]int{})
	val, err := s.Read()
	if err != EOS || val != -1 {
		t.Errorf("Expecting End Of Stream error!")
	}
}

func TestStreamWithInts(t *testing.T) {
	s := NewStream([]int{1, 2, 3})

	for i := 0; i < 3; i++ {
		val, err := s.Read()
		if err != nil {
			t.Error("There should be 3 elements in stream!")
		}

		expected := i + 1
		if val != expected {
			t.Errorf("Stream returned wrong value! Expected: %d, Got: %d", expected, val)
		}
	}

	val, err := s.Read()
	if err != EOS || val != -1 {
		t.Errorf("Expecting End Of Stream error!")
	}
}
