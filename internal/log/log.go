package log

import (
	"github.com/myaser/zettel-synch/internal/config"
	"go.uber.org/zap"
)

// NewLogger configures the logger
func NewLogger(e config.Environment) (*zap.Logger, error) {
	var lc zap.Config
	switch e {
	case config.Development:
		lc = zap.NewDevelopmentConfig()
	default:
		lc = zap.NewProductionConfig()
	}
	return lc.Build()
}
