package ratelimiter

import (
  "errors"
  "fmt"
)

func checkPermitsPerSecondsArg(permitsPerSeconds float64) {
    if permitsPerSeconds <=0 {
      panic(fmt.Sprintf("permitsPerSeconds must be positive: %v",
        permitsPerSeconds))
    }
}

func checkWarmupPeriodArg(warmupPeriod int64) {
    if warmupPeriod < 0 {
      panic(fmt.Sprintf("warmupPeriod must not be negative: %v",
        warmupPeriod))
    }
}

func parsePermits(permits []int) (int, error) {
  if len(permits) > 1 {
    return 0, errors.New("there should be empty or one permits")
  }

  // Default to 1 when permits is empty.
  p := 1
  if len(permits) > 0 {
    p = permits[0]
  }
  return p, nil
}

func min(x, y int64) int64 {
  if x < y {
    return x
  }
  return y
}

func max(x, y int64) int64 {
  if x > y {
    return x
  }
  return y
}
