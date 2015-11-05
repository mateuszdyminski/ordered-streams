package stream

import (
	"github.com/Sirupsen/logrus"
	"sort"
)

// Order returns channel of ordered elements taken from all streams. Streams are passed as arguments to the Order method.
func Order(streams ...StreamReader) chan int {
	// create a map of stream head values
	heads := make(map[int]head)

	// fulfil heads
	for i, s := range streams {
		val, err := s.Read()
		if err != nil {
			logrus.Debugf("Stream %d is empty!", i)
		} else {
			heads[i] = head{val: val, streamID: i}
		}
	}

	// create result chan
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			// check if there is any element in heads map
			if len(heads) == 0 {
				return
			}

			// build array of heads
			headsArr := make(byHeadVal, 0, len(heads))
			for k := range heads {
				headsArr = append(headsArr, heads[k])
			}

			// sort
			sort.Sort(headsArr)

			// get index of stream with lowest value
			i := headsArr[0].streamID

			// get lowest value and get next item from stream
			logrus.Debugf("Sending value %d taken from stream %d", heads[i].val, i)
			out <- heads[i].val
			val, err := streams[i].Read()
			if err != nil {
				logrus.Debugf("Stream %d is empty!", i)
				delete(heads, i)

				if len(heads) == 0 {
					logrus.Debugf("All values printed!")
					return
				}
			} else {
				h := heads[i]
				h.val = val
				heads[i] = h
			}
		}
	}()

	return out
}

// Head - helper struct which holds info about the current head of the stream and its index.
type head struct {
	val      int
	streamID int
}

// ByHeadVal - helper struct which holds encapsulates slice of Head elements. Needed only for sorting purpose.
type byHeadVal []head

func (a byHeadVal) Len() int           { return len(a) }
func (a byHeadVal) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byHeadVal) Less(i, j int) bool { return a[i].val < a[j].val }
