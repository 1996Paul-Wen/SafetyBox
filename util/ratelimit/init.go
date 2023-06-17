package ratelimit

import (
	"github.com/1996Paul-Wen/watchdog"
)

var defaultLimit = 200.00
var defaultBurst = 200.00
var limiters = make(map[string]*watchdog.Limiter)

// SetDefaultLimit 更新defaultLimit
func SetDefaultLimit(limit float64) {
	defaultLimit = limit
}

// SetDefaultBurst 更新defaultBurst
func SetDefaultBurst(burst float64) {
	defaultBurst = burst
}

// InsertLimiter 插入一个Limiter
func InsertLimiter(key string, singleUserLimit float64, burst float64) {
	limiter := watchdog.NewLimiter(singleUserLimit, burst)
	limiters[key] = limiter
}

// LoadLimiter 载入Limiter
func LoadLimiter(key string) *watchdog.Limiter {
	if _, ok := limiters[key]; !ok {
		InsertLimiter(key, defaultLimit, defaultBurst)
	}
	return limiters[key]
}

// RefreshLimiters 更新limiters
func RefreshLimiters(limit float64, burst float64) {
	// refresh default conf
	if limit > 0 {
		SetDefaultLimit(limit)
	}
	if burst > 0 {
		SetDefaultBurst(burst)
	}

	// refresh existed limiters
	for _, l := range limiters {
		if limit > 0 {
			l.SetLimit(limit)
		}
		if burst > 0 {
			l.SetBurst(burst)
		}
	}
}
