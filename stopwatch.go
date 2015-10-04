package ratelimiter

import (
  "time"
)

// A stopWatch implements StopWatch interface. This exists to enable unit test.
type stopWatch struct {
}

// newStopWatch returns a new StopWatch implementation using functions
// in time package.
func newStopWatch() StopWatch {
  return StopWatch(&stopWatch{})
}

func (r *stopWatch) ReadNanoseconds() int64 {
  return time.Now().UnixNano()
}

func (r *stopWatch) SleepNanosecondsUninterruptibly(nanoSeconds int64) {
  time.Sleep(time.Duration(nanoSeconds))
}
