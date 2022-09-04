package timeserver

import "time"

type Server interface {
	GetCurrentTime() (time.Time, error)
}
