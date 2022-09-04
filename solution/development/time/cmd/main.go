package main

import (
	"time/internal/app"
	"time/internal/timeserver"
	"time/internal/visualization"
)

func main() {
	timeApp := app.NewApp(timeserver.NTP{}, visualization.ConsoleTime{})
	timeApp.Launch()
}
