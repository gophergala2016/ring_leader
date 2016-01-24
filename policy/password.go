package policy

import (
	"time"
)

type Password struct {
	Expiration time.Duration
}

func (p Password) IsExpired(lastChanged time.Time) bool {
	return time.Since(lastChanged) > p.Expiration
}
