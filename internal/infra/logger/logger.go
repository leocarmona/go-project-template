package logger

import (
	"context"
	"github.com/google/uuid"
	"github.com/leocarmona/go-project-template/internal/infra/logger/attributes"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sync"
)

type Option struct {
	ServiceName    string
	ServiceVersion string
	Environment    string
	LogLevel       string
}

var (
	once    sync.Once
	skipper = zap.AddCallerSkip(1)
	option  *Option
)

func Init(opt *Option) {
	once.Do(func() {
		option = opt
		zapLogger, err := newZap()

		if err != nil {
			panic(err)
		}

		zap.ReplaceGlobals(zapLogger)
	})
}

func Sync() {
	_ = zap.L().Sync()
}

func Debug(ctx context.Context, message string, attr attributes.Attributes) {
	zap.L().WithOptions(skipper).Debug(message, makeAttributes(ctx, attr)...)
}

func Info(ctx context.Context, message string, attr attributes.Attributes) {
	zap.L().WithOptions(skipper).Info(message, makeAttributes(ctx, attr)...)
}

func Warn(ctx context.Context, message string, attr attributes.Attributes) {
	zap.L().WithOptions(skipper).Warn(message, makeAttributes(ctx, attr)...)
}

func Fatal(ctx context.Context, message string, attr attributes.Attributes) {
	zap.L().WithOptions(skipper).Fatal(message, makeAttributes(ctx, attr)...)
}

func Error(ctx context.Context, message string, attr attributes.Attributes) {
	zap.L().WithOptions(skipper).Error(message, makeAttributes(ctx, attr)...)
}

func newZap(options ...zap.Option) (*zap.Logger, error) {
	return newZapConfig().Build(options...)
}

func newZapConfig() zap.Config {
	logLevel := zap.NewAtomicLevel()
	err := logLevel.UnmarshalText([]byte(option.LogLevel))

	if err != nil {
		panic(err)
	}

	return zap.Config{
		Level:       logLevel,
		Development: false,
		Sampling: &zap.SamplingConfig{
			Initial:    100,
			Thereafter: 100,
		},
		Encoding:         "json",
		EncoderConfig:    newZapEncoderConfig(),
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
	}
}

func newZapEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "Timestamp",
		LevelKey:       "SeverityText",
		NameKey:        "Logger",
		CallerKey:      "Caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "Body",
		StacktraceKey:  "Stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.EpochNanosTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func makeAttributes(ctx context.Context, attr attributes.Attributes) []zapcore.Field {
	if attr == nil {
		attr = attributes.New()
	}

	if attr["cid"] == nil {
		cid := ctx.Value("cid")

		if cid == nil || cid == "" {
			cid = uuid.New().String()
		}

		attr["cid"] = cid
	}

	return []zapcore.Field{
		zap.Any("Attributes", attr),
		zap.Any("Resource", map[string]interface{}{
			"service.name":        option.ServiceName,
			"service.version":     option.ServiceVersion,
			"service.environment": option.Environment,
		}),
	}
}
