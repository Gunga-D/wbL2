package visualization

import (
	"fmt"
	"time"
)

type ConsoleTime struct {
}

func (c ConsoleTime) Show(showingTime time.Time) {
	fmt.Println(showingTime)
}
