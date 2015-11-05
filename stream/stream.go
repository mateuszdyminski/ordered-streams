package stream

import "errors"

// EOS End Of Stream error.
var EOS = errors.New("Reach end of stream!")

// StreamReader interface with only one method Read which returns next element of the stream.
type StreamReader interface {
	Read() (int, error)
}

// NewStream factory method which returns new instance of Stream.
func NewStream(ints []int) *Stream {
	return &Stream{data: ints}
}

// Stream of ints.
type Stream struct {
	data     []int
	position int
}

// Read returns next int if there is some remaining element in stream. Otherwise EOS error will be thrown.
func (s *Stream) Read() (int, error) {
	defer func() { s.position++ }()

	if len(s.data) <= s.position {
		return -1, EOS
	}

	return s.data[s.position], nil
}

