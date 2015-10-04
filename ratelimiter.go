// go implementation of com.google.common.util.concurrent.RateLimiter
package ratelimiter

import (
  "sync"
)

// A RateLimiter represents a RateLimiter instance.
type RateLimiter struct {
  sync.Mutex

  worker limterWorker

  stopWatch StopWatch
  nextFreeTicketNano float64
}

// SetRate updates the stable rate of this rateLimiter.
func (r *RateLimiter) SetRate(permitsPerSeconds float64) {
  checkPermitsPerSecondsArg(permitsPerSeconds)
  r.Lock()
  defer r.Unlock()
  r.worker.doSetRate(permitsPerSeconds, r.stopWatch.ReadNanoseconds())
}

// GetRate returns the stable rate (as permits per seconds)
// with which this RateLimiter is configured with.
func (r *RateLimiter) GetRate() float64 {
  r.Lock()
  defer r.Unlock()
  return r.worker.doGetRate()
}

// Acquare reserves the given number of permits from this rateLimiter
// for future use, returning the number of nanoseconds until the reservation
// can be consumed.
func (r *RateLimiter) Acquire(permits ...int) (int64, error) {
  permit, err := parsePermits(permits)
  if err != nil {
    return 0, err
  }

  r.Lock()
  defer r.Unlock()
  nowNanoSeconds := r.stopWatch.ReadNanoseconds()
  momentAvailable := r.worker.reserveEarliestAvailable(float64(permit), nowNanoSeconds)
  nanoSecondsToWait := max(momentAvailable - nowNanoSeconds, 0)
  r.stopWatch.SleepNanosecondsUninterruptibly(nanoSecondsToWait)
  return nanoSecondsToWait, nil
}
