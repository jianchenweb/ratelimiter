// go implementation of java TimeUnitTest.
package ratelimiter_test

import (
  "testing"

  helper "github.com/jianguochen/ratelimiter"
)

func TestToNanos(t *testing.T) {
  for i :=int64(0); i < 88888; i++ {
    assertEquals(i*1000000000*60*60*24, helper.DAYS.ToNanos(i), t)
    assertEquals(i*1000000000*60*60, helper.HOURS.ToNanos(i), t)
    assertEquals(i*1000000000*60, helper.MINUTES.ToNanos(i), t)
    assertEquals(i*1000000000, helper.SECONDS.ToNanos(i), t)
    assertEquals(i*1000000, helper.MILLISECONDS.ToNanos(i), t)
    assertEquals(i*1000, helper.MICROSECONDS.ToNanos(i), t)
    assertEquals(i, helper.NANOSECONDS.ToNanos(i), t)
  }
}

func TestToMicros(t *testing.T) {
  for i :=int64(0); i < 88888; i++ {
    assertEquals(i*1000000*60*60*24, helper.DAYS.ToMicros(i), t)
    assertEquals(i*1000000*60*60, helper.HOURS.ToMicros(i), t)
    assertEquals(i*1000000*60, helper.MINUTES.ToMicros(i), t)
    assertEquals(i*1000000, helper.SECONDS.ToMicros(i), t)
    assertEquals(i*1000, helper.MILLISECONDS.ToMicros(i), t)
    assertEquals(i, helper.MICROSECONDS.ToMicros(i), t)
    assertEquals(i, helper.NANOSECONDS.ToMicros(i * 1000), t)
  }
}

func TestToMillis(t *testing.T) {
  for i :=int64(0); i < 88888; i++ {
    assertEquals(i*1000*60*60*24, helper.DAYS.ToMillis(i), t)
    assertEquals(i*1000*60*60, helper.HOURS.ToMillis(i), t)
    assertEquals(i*1000*60, helper.MINUTES.ToMillis(i), t)
    assertEquals(i*1000, helper.SECONDS.ToMillis(i), t)
    assertEquals(i, helper.MILLISECONDS.ToMillis(i), t)
    assertEquals(i, helper.MICROSECONDS.ToMillis(i * 1000), t)
    assertEquals(i, helper.NANOSECONDS.ToMillis(i * 1000000), t)
  }
}

func TestToSeconds(t *testing.T) {
  for i :=int64(0); i < 88888; i++ {
    assertEquals(i*60*60*24, helper.DAYS.ToSeconds(i), t)
    assertEquals(i*60*60, helper.HOURS.ToSeconds(i), t)
    assertEquals(i*60, helper.MINUTES.ToSeconds(i), t)
    assertEquals(i, helper.SECONDS.ToSeconds(i), t)
    assertEquals(i, helper.MILLISECONDS.ToSeconds(i * 1000), t)
    assertEquals(i, helper.MICROSECONDS.ToSeconds(i * 1000000), t)
    assertEquals(i, helper.NANOSECONDS.ToSeconds(i * 1000000000), t)
  }
}

func TestToMinutes(t *testing.T) {
  for i :=int64(0); i < 88888; i++ {
    assertEquals(i*60*24, helper.DAYS.ToMinutes(i), t)
    assertEquals(i*60, helper.HOURS.ToMinutes(i), t)
    assertEquals(i, helper.MINUTES.ToMinutes(i), t)
    assertEquals(i, helper.SECONDS.ToMinutes(i*60), t)
    assertEquals(i, helper.MILLISECONDS.ToMinutes(i * 1000 * 60), t)
    assertEquals(i, helper.MICROSECONDS.ToMinutes(i * 1000000 * 60), t)
    assertEquals(i, helper.NANOSECONDS.ToMinutes(i * 1000000000 * 60), t)
  }
}

func TestToHours(t *testing.T) {
  for i :=int64(0); i < 88888; i++ {
    assertEquals(i*24, helper.DAYS.ToHours(i), t)
    assertEquals(i, helper.HOURS.ToHours(i), t)
    assertEquals(i, helper.MINUTES.ToHours(i * 60), t)
    assertEquals(i, helper.SECONDS.ToHours(i * 60 * 60), t)
    assertEquals(i,
      helper.MILLISECONDS.ToHours(i * 1000 * 60 * 60), t)
    assertEquals(i,
      helper.MICROSECONDS.ToHours(i * 1000000 * 60 * 60), t)
    assertEquals(i,
      helper.NANOSECONDS.ToHours(i * 1000000000 * 60 * 60), t)
  }
}

func TestToDays(t *testing.T) {
  for i :=int64(0); i < 88888; i++ {
    assertEquals(i, helper.DAYS.ToDays(i), t)
    assertEquals(i, helper.HOURS.ToDays(i * 24), t)
    assertEquals(i, helper.MINUTES.ToDays(i * 24 * 60), t)
    assertEquals(i, helper.SECONDS.ToDays(i * 24 * 60 * 60), t)
    assertEquals(i,
      helper.MILLISECONDS.ToDays(i * 1000 * 24 * 60 * 60), t)
    assertEquals(i,
      helper.MICROSECONDS.ToDays(i * 1000000 * 24 * 60 * 60), t)
    assertEquals(i,
      helper.NANOSECONDS.ToDays(i * 1000000000 * 24 * 60 * 60), t)
  }
}

func assertEquals(expected int64, actual int64, t *testing.T) {
  if expected != actual {
    t.Errorf("Expected %v, got %v", expected, actual)
  }
}
