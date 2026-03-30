package gorrito

import (
	"sync"
	"time"
)

type rateLimiter struct {
	perSecond int
	perMinute int

	mu sync.Mutex

	secondWindowStart time.Time
	minuteWindowStart time.Time
	secondCount       int
	minuteCount       int

	nowFn   func() time.Time
	sleepFn func(time.Duration)
}

func newRateLimiter(perSecond, perMinute int) *rateLimiter {
	return &rateLimiter{
		perSecond: perSecond,
		perMinute: perMinute,
		nowFn:     time.Now,
		sleepFn:   time.Sleep,
	}
}

func (r *rateLimiter) wait() {
	if r == nil {
		return
	}
	if r.perSecond <= 0 && r.perMinute <= 0 {
		return
	}

	for {
		now := r.nowFn()
		r.mu.Lock()

		if r.secondWindowStart.IsZero() || now.Sub(r.secondWindowStart) >= time.Second {
			r.secondWindowStart = now
			r.secondCount = 0
		}
		if r.minuteWindowStart.IsZero() || now.Sub(r.minuteWindowStart) >= time.Minute {
			r.minuteWindowStart = now
			r.minuteCount = 0
		}

		secondAllowed := r.perSecond <= 0 || r.secondCount < r.perSecond
		minuteAllowed := r.perMinute <= 0 || r.minuteCount < r.perMinute

		if secondAllowed && minuteAllowed {
			r.secondCount++
			r.minuteCount++
			r.mu.Unlock()
			return
		}

		var waitFor time.Duration
		if !secondAllowed {
			waitSecond := r.secondWindowStart.Add(time.Second).Sub(now)
			if waitSecond > waitFor {
				waitFor = waitSecond
			}
		}
		if !minuteAllowed {
			waitMinute := r.minuteWindowStart.Add(time.Minute).Sub(now)
			if waitMinute > waitFor {
				waitFor = waitMinute
			}
		}
		r.mu.Unlock()

		if waitFor <= 0 {
			waitFor = time.Millisecond
		}
		r.sleepFn(waitFor)
	}
}
