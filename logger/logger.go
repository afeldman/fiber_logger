package logger

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// the backend logger
var Logger *zap.SugaredLogger

// set the loglevel depending on a string
func logLevel() zapcore.LevelEnabler {
	// make upper or lower string
	lower := strings.ToLower(LoggingConfig.Level)

	// enable the level
	var level zapcore.LevelEnabler

	// check for the loglevel
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
	} // default error level

	return level
}

// initialize the logger
func InitLogger(logger *LoggerConfig) {

	// if no logger set,
	if logger == nil {
		DefaultLoggingConfig()
	}

	// get the lumberjeck logger
	writerSyncer := getLogWriter()

	// get the encoder and format
	encoder := getEncoder()

	// create the core for logging
	core := zapcore.NewCore(encoder, writerSyncer, logLevel())

	// create the zap logger
	zaplogger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	// get the sugar information
	Logger = zaplogger.Sugar()

	// set the sync on object destruction
	defer Logger.Sync()
}

// get the text encoder
func getEncoder() zapcore.Encoder {
	// get standard encoder
	encoderConfig := zap.NewProductionEncoderConfig()

	// set the timer
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// get the loglevel
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// the console encoder. so console and file encoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// get lumberjack
func getLogWriter() zapcore.WriteSyncer {
	// create lumberjack object
	lumberJackLogger := &lumberjack.Logger{
		Filename:   LoggingConfig.FileName,
		MaxSize:    LoggingConfig.Size,
		MaxBackups: LoggingConfig.Backups,
		MaxAge:     LoggingConfig.Backups,
		Compress:   LoggingConfig.Compress,
	}
	// add to zap
	return zapcore.AddSync(lumberJackLogger)
}

// write debug information
func Debug(args ...interface{}) {
	Logger.Debug(args...)
}

// write debug information with a formatted string
func Debugf(template string, args ...interface{}) {
	Logger.Debugf(template, args...)
}

// get information
func Info(args ...interface{}) {
	Logger.Info(args...)
}

// get formatted information
func Infof(template string, args ...interface{}) {
	Logger.Infof(template, args...)
}

// warning
func Warn(args ...interface{}) {
	Logger.Warn(args...)
}

// formatted warning
func Warnf(template string, args ...interface{}) {
	Logger.Warnf(template, args...)
}

// get error message
func Error(args ...interface{}) {
	Logger.Error(args...)
}

// format error message
func Errorf(template string, args ...interface{}) {
	Logger.Errorf(template, args...)
}

// never run in panic
func DPanic(args ...interface{}) {
	Logger.DPanic(args...)
}

// but if panic make a formatted string
func DPanicf(template string, args ...interface{}) {
	Logger.DPanicf(template, args...)
}

// Panic information
func Panic(args ...interface{}) {
	Logger.Panic(args...)
}

// formated panic
func Panicf(template string, args ...interface{}) {
	Logger.Panicf(template, args...)
}

// fatal is horrifying
func Fatal(args ...interface{}) {
	Logger.Fatal(args...)
}

// horror with format :D
func Fatalf(template string, args ...interface{}) {
	Logger.Fatalf(template, args...)
}
