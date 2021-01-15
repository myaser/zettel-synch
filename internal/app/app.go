package app

import (
	"context"
	"sync"
	"syscall"

	"github.com/myaser/zettel-synch/internal/config"
	"github.com/myaser/zettel-synch/internal/log"
	"go.uber.org/zap"
	"gopkg.in/vrecan/death.v3"
)

// Application is the main object
type Application struct {
	config *config.Config
	logger *zap.Logger
	closer *death.Death
}

var app *Application
var once sync.Once

// GetApplication initializes the app
func GetApplication() (*Application, error) {
	var err error
	once.Do(func() {
		c, e := config.Load()
		if e != nil {
			err = e
			return
		}
		logger, e := log.NewLogger(c.Environment)
		if e != nil {
			err = e
			return
		}
		app = &Application{
			config: c,
			logger: logger,
			closer: death.NewDeath(syscall.SIGINT, syscall.SIGTERM),
		}
	})

	return app, err
}

// func (app *Application) Close() error{
// 	cancelContext()
// 	if e := app.logger.Sync(); e != nil {
// 		err = e
// 	}
// 	return nil
// }

// Run the main process of the application
func (app *Application) Run() (err error) {
	// ctx, cancelFunc := context.WithCancel(context.Background())
	_, cancelContext := context.WithCancel(context.Background())

	// TODO: app logic here

	app.closer.WaitForDeathWithFunc(func() {
		cancelContext()
		if e := app.logger.Sync(); e != nil {
			err = e
		}
	})
	return err
}
