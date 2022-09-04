package timeserver

import (
	"time"

	"github.com/beevik/ntp"
)

type NTP struct {
}

func (s NTP) GetCurrentTime() (time.Time, error) {
	return ntp.Time("0.beevik-ntp.pool.ntp.org")
}
