package ratelimiter

type limterWorker interface {
  doSetRate(permitsPerSecond float64, nowNanoSeconds int64)
  doGetRate() float64
  reserveEarliestAvailable(permits float64, nowNanoSeconds int64) int64
}

type StopWatch interface {
  ReadNanoseconds() int64
  SleepNanosecondsUninterruptibly(nanoSeconds int64)
}
