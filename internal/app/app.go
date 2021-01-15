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
		}
	})

	return app, err
}

// Run the main process of the application
func (app *Application) Run() error {
	// ctx, cancelFunc := context.WithCancel(context.Background())
	_, cancelFunc := context.WithCancel(context.Background())

	death := death.NewDeath(syscall.SIGINT, syscall.SIGTERM)

	// TODO: app logic here

	death.WaitForDeathWithFunc(func() {
		cancelFunc()
		app.logger.Sync()
		// err := app.logger.Sync()
		// if err != nil {
		// 	log.Println(err)
		// 	os.Exit(1)
		// }
	})
	return nil
}
