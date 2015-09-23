// go implementation of java.util.concurrent.TimeUnit
package ratelimiter

// Handy constant for conversion methods.
const (
  sixty = int64(60)
  oneThousand = int64(1000)
  c0 = int64(1)
  c1 = c0 * oneThousand
  c2 = c1 * oneThousand
  c3 = c2 * oneThousand
  c4 = c3 * sixty
  c5 = c4 * sixty
  c6 = c5 * int64(24)
)

// TimeUnit interface.
type TimeUnit interface {
  ToNanos(int64) int64
  ToMicros(int64) int64
  ToMillis(int64) int64
  ToSeconds(int64) int64
  ToMinutes(int64) int64
  ToHours(int64) int64
  ToDays(int64) int64
  //Convert(int64, TimeUnit) int64
}

// Represents nano seconds.
type nanoSeconds int64

func (n nanoSeconds) ToNanos(d int64) int64 {
  return d
}

func (n nanoSeconds) ToMicros(d int64) int64 {
  return d / (c1 / c0)
}

func (n nanoSeconds) ToMillis(d int64) int64 {
  return d / (c2 / c0)
}

func (n nanoSeconds) ToSeconds(d int64) int64 {
  return d / (c3 / c0)
}

func (n nanoSeconds) ToMinutes(d int64) int64 {
  return d / (c4 / c0)
}

func (n nanoSeconds) ToHours(d int64) int64 {
  return d / (c5 / c0)
}

func (n nanoSeconds) ToDays(d int64) int64 {
  return d / (c6 / c0)
}

// Singlton for nano second conversion.
var NANOSECONDS = new(nanoSeconds)

// Represents micro seconds.
type microSeconds int64

func (m microSeconds) ToNanos(d int64) int64 {
  return d * (c1 / c0)
}

func (m microSeconds) ToMicros(d int64) int64 {
  return d
}

func (m microSeconds) ToMillis(d int64) int64 {
  return d / (c2 / c1)
}

func (m microSeconds) ToSeconds(d int64) int64 {
  return d / (c3 / c1)
}

func (m microSeconds) ToMinutes(d int64) int64 {
  return d / (c4 / c1)
}

func (m microSeconds) ToHours(d int64) int64 {
  return d / (c5 / c1)
}

func (m microSeconds) ToDays(d int64) int64 {
  return d / (c6 / c1)
}

// Singlton for micro seconds conversion.
var MICROSECONDS = new(microSeconds)

// Represents milli second.
type milliSeconds int64

func (m milliSeconds) ToNanos(d int64) int64 {
  return d * (c2 / c0)
}

func (m milliSeconds) ToMicros(d int64) int64 {
  return d * (c2 / c1)
}

func (m milliSeconds) ToMillis(d int64) int64 {
  return d
}

func (m milliSeconds) ToSeconds(d int64) int64 {
  return d / (c3 / c2)
}

func (m milliSeconds) ToMinutes(d int64) int64 {
  return d / (c4 / c2)
}

func (m milliSeconds) ToHours(d int64) int64 {
  return d / (c5 / c2)
}

func (m milliSeconds) ToDays(d int64) int64 {
  return d / (c6 / c2)
}

// Singlton for milli seconds conversion.
var MILLISECONDS = new(milliSeconds)

// Represents second.
type seconds int64

func (s seconds) ToNanos(d int64) int64 {
  return d * (c3 / c0)
}

func (s seconds) ToMicros(d int64) int64 {
  return d * (c3 / c1)
}

func (s seconds) ToMillis(d int64) int64 {
  return d * (c3 / c2)
}

func (s seconds) ToSeconds(d int64) int64 {
  return d
}

func (s seconds) ToMinutes(d int64) int64 {
  return d / (c4 / c3)
}

func (s seconds) ToHours(d int64) int64 {
  return d / (c5 / c3)
}

func (s seconds) ToDays(d int64) int64 {
  return d / (c6 / c3)
}

// Singlton for seconds conversion.
var SECONDS = new(seconds)

// Represents minutes.
type minutes int64

func (m minutes) ToNanos(d int64) int64 {
  return d * (c4 / c0)
}

func (m minutes) ToMicros(d int64) int64 {
  return d * (c4 / c1)
}

func (m minutes) ToMillis(d int64) int64 {
  return d * (c4 / c2)
}

func (m minutes) ToSeconds(d int64) int64 {
  return d * (c4 / c3)
}

func (m minutes) ToMinutes(d int64) int64 {
  return d
}

func (m minutes) ToHours(d int64) int64 {
  return d / (c5 / c4)
}

func (m minutes) ToDays(d int64) int64 {
  return d / (c6 / c4)
}

// Singlton for minutes conversion.
var MINUTES = new(minutes)

// Represents hours.
type hours int64

func (h hours) ToNanos(d int64) int64 {
  return d * (c5 / c0)
}

func (h hours) ToMicros(d int64) int64 {
  return d * (c5 / c1)
}

func (h hours) ToMillis(d int64) int64 {
  return d * (c5 / c2)
}

func (h hours) ToSeconds(d int64) int64 {
  return d * (c5 /c3)
}

func (h hours) ToMinutes(d int64) int64 {
  return d * (c5 / c4)
}

func (h hours) ToHours(d int64) int64 {
  return d
}

func (h hours) ToDays(d int64) int64 {
  return d / (c6 / c5)
}

// Singlton for hours conversion.
var HOURS = new(hours)

// Represents days.
type days int64

func (h days) ToNanos(d int64) int64 {
  return d * (c6 / c0)
}

func (h days) ToMicros(d int64) int64 {
  return d * (c6 / c1)
}

func (h days) ToMillis(d int64) int64 {
  return d * (c6 / c2)
}

func (h days) ToSeconds(d int64) int64 {
  return d * (c6 /c3)
}

func (h days) ToMinutes(d int64) int64 {
  return d * (c6 / c4)
}

func (h days) ToHours(d int64) int64 {
  return d * (c6 / c5)
}

func (h days) ToDays(d int64) int64 {
  return d
}

// Singlton for days conversion.
var DAYS = new(days)
