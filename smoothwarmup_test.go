package ratelimiter_test

import (
  "reflect"
  "testing"

  . "github.com/jianguochen/ratelimiter"
)

func TestToNanos(t *testing.T) {
  stopWatch := &fakeWatch{0, []int64{}}
  limiter := NewSmoothWarmupRateLimiter(2.0, 4000, MILLISECONDS, stopWatch)
  for i := 0; i < 2; i++ {
    limiter.Acquire()
    t.Log(limiter)
  }
  expected := []int64{0, 500000000}
  assertRateEquals(2.0, limiter.GetRate(), t)
  assertEquals(expected, stopWatch, t)
}

func assertEquals(expected []int64, watch *fakeWatch, t *testing.T) {
  if !reflect.DeepEqual(expected, watch.events) {
    t.Errorf("Expected %v, got %v", expected, watch.events)
  }
}

func assertRateEquals(expected float64, actual float64, t *testing.T) {
  if expected != actual {
    t.Errorf("Expected rate %v, got %v", expected, actual)
  }
}
type fakeWatch struct {
  instant int64
  events []int64
}

func (f *fakeWatch) ReadNanoseconds() int64 {
  return f.instant
}

func (f *fakeWatch) SleepNanosecondsUninterruptibly(nanoSeconds int64) {
  f.instant = f.instant + nanoSeconds
  f.events = append(f.events, nanoSeconds)
}
