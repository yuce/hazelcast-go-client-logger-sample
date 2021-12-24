package main

import (
	"fmt"

	"github.com/hazelcast/hazelcast-go-client/logger"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogAdaptor adapts zap.SugaredLogger to use as a custom Hazelcst logger.
type ZapLogAdaptor struct {
	lg     *zap.SugaredLogger
	weight logger.Weight
}

// NewZapLogAdaptor creates a new zap log adaptor.
func NewZapLogAdaptor(weight logger.Weight, lg *zap.Logger) *ZapLogAdaptor {
	return &ZapLogAdaptor{
		lg:     lg.Sugar(),
		weight: weight,
	}
}

// Log implements Hazelcast custom logger.
func (z ZapLogAdaptor) Log(wantWeight logger.Weight, f func() string) {
	// Do not bother calling f if the current log level does not permit logging this message.
	if z.weight < wantWeight {
		return
	}
	// Call the appropriate log function.
	switch wantWeight {
	case logger.WeightTrace:
		fallthrough
	case logger.WeightDebug:
		z.lg.Debug(f())
	case logger.WeightInfo:
		z.lg.Info(f())
	case logger.WeightWarn:
		z.lg.Warn(f())
	case logger.WeightError:
		z.lg.Error(f())
	case logger.WeightFatal:
		z.lg.Fatal(f())
	}
}

// MakeZapLogger creates a zap Logger with defaults.
func MakeZapLogger(callerSkip int) *zap.Logger {
	var cfg zap.Config
	cfg.Encoding = "console"
	// Set the logger to the finest setting, since we handle log filtering manually.
	cfg.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	cfg.OutputPaths = []string{"stdout"}
	cfg.ErrorOutputPaths = []string{"stderr"}
	// Use production defaults ...
	cfg.EncoderConfig = zap.NewProductionEncoderConfig()
	// ... with our settings.
	ec := &cfg.EncoderConfig
	ec.EncodeLevel = zapcore.CapitalColorLevelEncoder
	ec.EncodeTime = zapcore.ISO8601TimeEncoder
	// Try commenting out the following line.
	//ec.FunctionKey = "func"
	// Adjust the call stack so the root coller is displayed in the logs.
	lg, err := cfg.Build(zap.AddCallerSkip(callerSkip))
	if err != nil {
		panic(fmt.Errorf("creating ZapLogAdaptor: %w", err))
	}
	return lg
}
