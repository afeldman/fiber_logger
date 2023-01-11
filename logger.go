package fiber_logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.SugaredLogger

func logLevel() zapcore.LevelEnabler {
	lower := strings.ToLower(noue_config.Config.Logger.Level)
	var level zapcore.LevelEnabler

	switch lower {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "invalide":
		level = zapcore.InvalidLevel
	case "warning":
		level = zapcore.WarnLevel
	case "panic":
		level = zapcore.PanicLevel
	case "fatal":
		level = zapcore.FatalLevel
	default:
		level = zapcore.ErrorLevel
	}

	return level
}

func InitNoueLogger() {
	writerSyncer := getLogWriter()
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writerSyncer, logLevel())
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	NoueLogger = logger.Sugar()
	defer NoueLogger.Sync()
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   noue_config.Config.Logger.FileName,
		MaxSize:    noue_config.Config.Logger.Size,
		MaxBackups: noue_config.Config.Logger.Backups,
		MaxAge:     noue_config.Config.Logger.Backups,
		Compress:   noue_config.Config.Logger.Compress,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func Debug(args ...interface{}) {
	NoueLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	NoueLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	NoueLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	NoueLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	NoueLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	NoueLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	NoueLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	NoueLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	NoueLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	NoueLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	NoueLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	NoueLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	NoueLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	NoueLogger.Fatalf(template, args...)
}
