package app

import (
	"fmt"
	"time/internal/timeserver"
	"time/internal/visualization"
)

type App struct {
	server    timeserver.Server
	shower    visualization.TimeShower
	isWorking bool
}

func NewApp(server timeserver.Server, shower visualization.TimeShower) *App {
	return &App{server: server, shower: shower, isWorking: true}
}

func (a *App) Launch() {
	a.isWorking = true
	for a.isWorking {
		currentTime, err := a.server.GetCurrentTime()
		if err != nil {
			fmt.Println("Error:", err.Error())
		}

		a.shower.Show(currentTime)
	}
}

func (a *App) Shutdown() {

}
