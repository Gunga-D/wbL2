package visualization

import "time"

type TimeShower interface {
	Show(time.Time)
}
