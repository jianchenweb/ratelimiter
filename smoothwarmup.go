package ratelimiter

import (
  "fmt"
  "math"
)

type SmoothWarmup struct {
  RateLimiter

  maxPermits    float64
  halfPermits   float64
  storedPermits float64

  slope         float64
  stableIntervalNano float64
  warmupPeriodNano float64
}

func NewSmoothWarmupRateLimiter(permitsPerSeconds float64,
  warmupPeriod int64, unit TimeUnit, stopWatch ...StopWatch) SmoothWarmup {
  // Check args.
  checkPermitsPerSecondsArg(permitsPerSeconds)
  checkWarmupPeriodArg(warmupPeriod)

  // Sets StopWatch. This is there for unit test purpose.
  var watch StopWatch
  if len(stopWatch) > 0 {
    watch = stopWatch[0]
  } else {
    watch = newStopWatch()
  }

  s := SmoothWarmup{
    RateLimiter: RateLimiter{stopWatch: watch},
    warmupPeriodNano: float64(unit.ToNanos(warmupPeriod))}
  s.worker = limterWorker(&s)
  s.SetRate(permitsPerSeconds)
  return s
}

func (s *SmoothWarmup) doSetRate(permitsPerSeconds float64, nowNano int64) {
  s.stableIntervalNano = float64(SECONDS.ToNanos(1)) / permitsPerSeconds
  s.resync(nowNano)
  oldMaxPermits := s.maxPermits
  s.maxPermits = s.warmupPeriodNano / s.stableIntervalNano
  s.halfPermits = s.maxPermits / 2
  s.slope =  (s.stableIntervalNano * 2) / s.halfPermits
  if oldMaxPermits == 0.0 {
    s.storedPermits = 0
  } else {
    s.storedPermits = s.storedPermits * s.maxPermits / oldMaxPermits
  }
}

func (s *SmoothWarmup) doGetRate() float64 {
  return float64(SECONDS.ToNanos(1)) / s.stableIntervalNano
}

func (s *SmoothWarmup) reserveEarliestAvailable(permits float64, nowNano int64) int64 {
  s.resync(nowNano)
  returnValue := s.nextFreeTicketNano
  storedPermitsToSpend := math.Min(permits, s.storedPermits)
  freshPermits := permits - storedPermitsToSpend

  waitNano := s.storedPermitsToWaitTime(storedPermitsToSpend) +
    (freshPermits * s.stableIntervalNano)

  s.nextFreeTicketNano = s.nextFreeTicketNano + waitNano
  s.storedPermits = s.storedPermits - storedPermitsToSpend
  return int64(returnValue)
}

func (s *SmoothWarmup) storedPermitsToWaitTime(permitsToTake float64) (waitTime float64) {
  availablePermitsAboveHalf := s.storedPermits - s.halfPermits;
  waitTime = 0
  // measuring the integral on the right part of the function (the climbing line)
  if availablePermitsAboveHalf > 0.0 {
    permitsAboveHalfToTake := math.Min(availablePermitsAboveHalf, permitsToTake)
    waitTime = permitsAboveHalfToTake * (s.permitsToTime(availablePermitsAboveHalf) +
      s.permitsToTime(availablePermitsAboveHalf - permitsAboveHalfToTake)) / 2.0
    permitsToTake = permitsToTake - permitsAboveHalfToTake
  }
  // measuring the integral on the left part of the function (the horizontal line)
  waitTime = waitTime + s.stableIntervalNano * permitsToTake
  return
}

func (s *SmoothWarmup) permitsToTime(permits float64) float64 {
  return s.stableIntervalNano + permits * s.slope
}

func (s *SmoothWarmup) resync(nowNano int64) {
  now := float64(nowNano)
  if now > s.nextFreeTicketNano {
    s.storedPermits = math.Min(s.maxPermits,
      s.storedPermits + (now - s.nextFreeTicketNano) / s.stableIntervalNano)
    s.nextFreeTicketNano = now
  }
}

func (s *SmoothWarmup) ToString() string {
  return fmt.Sprintf("%+v", s)
}
