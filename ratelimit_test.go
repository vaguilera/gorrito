package gorrito

import (
	"testing"
	"time"
)

func TestRateLimiterIgnoresZeroLimits(t *testing.T) {
	l := newRateLimiter(0, 0)

	now := time.Unix(100, 0)
	l.nowFn = func() time.Time { return now }
	slept := false
	l.sleepFn = func(_ time.Duration) {
		slept = true
	}

	l.wait()
	l.wait()
	l.wait()

	if slept {
		t.Fatal("limiter slept with zero limits")
	}
}

func TestRateLimiterWaitsPerSecond(t *testing.T) {
	l := newRateLimiter(1, 0)

	current := time.Unix(200, 0)
	l.nowFn = func() time.Time { return current }

	var sleeps []time.Duration
	l.sleepFn = func(d time.Duration) {
		sleeps = append(sleeps, d)
		current = current.Add(d)
	}

	l.wait()
	l.wait()

	if len(sleeps) != 1 {
		t.Fatalf("expected 1 sleep, got %d", len(sleeps))
	}
	if sleeps[0] != time.Second {
		t.Fatalf("expected sleep 1s, got %s", sleeps[0])
	}
}

func TestRateLimiterWaitsPerMinute(t *testing.T) {
	l := newRateLimiter(0, 2)

	current := time.Unix(300, 0)
	l.nowFn = func() time.Time { return current }

	var sleeps []time.Duration
	l.sleepFn = func(d time.Duration) {
		sleeps = append(sleeps, d)
		current = current.Add(d)
	}

	l.wait()
	l.wait()
	l.wait()

	if len(sleeps) != 1 {
		t.Fatalf("expected 1 sleep, got %d", len(sleeps))
	}
	if sleeps[0] != time.Minute {
		t.Fatalf("expected sleep 1m, got %s", sleeps[0])
	}
}

func TestRateLimiterAppliesSecondAndMinuteTogether(t *testing.T) {
	l := newRateLimiter(10, 2)

	current := time.Unix(400, 0)
	l.nowFn = func() time.Time { return current }

	var sleeps []time.Duration
	l.sleepFn = func(d time.Duration) {
		sleeps = append(sleeps, d)
		current = current.Add(d)
	}

	l.wait()
	l.wait()
	l.wait()

	if len(sleeps) != 1 {
		t.Fatalf("expected 1 sleep, got %d", len(sleeps))
	}
	if sleeps[0] != time.Minute {
		t.Fatalf("expected minute-based sleep, got %s", sleeps[0])
	}
}
