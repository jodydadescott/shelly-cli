package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// GetDefaultZapLogger returns ZapLogger with defaults. There should be no error.
func GetDefaultZapLogger() *zap.Logger {

	zapConfig := &zap.Config{
		Development: false,
		Level:       zap.NewAtomicLevelAt(defaultLogLevel),
		Sampling: &zap.SamplingConfig{
			Initial:    defaultSamplingInitial,
			Thereafter: defaultSamplingThereafter,
		},
		Encoding:      defaultEncoding,
		EncoderConfig: zap.NewProductionEncoderConfig(),
	}

	zapConfig.OutputPaths = append(zapConfig.OutputPaths, "stderr")
	zapConfig.ErrorOutputPaths = append(zapConfig.ErrorOutputPaths, "stderr")

	logger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}

	return logger
}

// GetDebugZapLogger returns ZapLogger with debug. There should be no error.
func GetDebugZapLogger() *zap.Logger {

	zapConfig := &zap.Config{
		Development: true,
		Level:       zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Sampling: &zap.SamplingConfig{
			Initial:    defaultSamplingInitial,
			Thereafter: defaultSamplingThereafter,
		},
		Encoding:      defaultEncoding,
		EncoderConfig: zap.NewProductionEncoderConfig(),
	}

	zapConfig.OutputPaths = append(zapConfig.OutputPaths, "stderr")
	zapConfig.ErrorOutputPaths = append(zapConfig.ErrorOutputPaths, "stderr")

	logger, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}

	return logger
}
