package app

import (
	"unpacking/internal/config"
	"unpacking/internal/reader"
	"unpacking/internal/usecase"
)

type App struct {
	conf  *config.Config
	input reader.Reader
}

func NewApp(input reader.Reader, conf *config.Config) *App {
	return &App{input: input, conf: conf}
}

func (a App) Execute() (string, error) {
	input, err := a.input.ReadAll(a.conf.File.Path + a.conf.File.Name)
	if err != nil {
		return "", err
	}

	return usecase.UnpackString(input)
}
