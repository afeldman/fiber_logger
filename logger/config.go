package logger

type LoggerConfig struct {
	FileName string `mapstructure:"path"`
	Size     int    `mapstructure:"max"`
	Backups  int    `mapstructure:"backups"`
	Age      int    `mapstructure:"age"`
	Compress bool   `mapstructure:"compress"`
	Level    string `mapstructure:"level"`
}

var LoggingConfig *LoggerConfig

func InitLoggerConfig(filepath, loglevel string, sizeoffile, numberofbackups, ageoffile int, compress bool) *LoggerConfig {
	LoggingConfig = &LoggerConfig{
		FileName: filepath,
		Level:    loglevel,
		Size:     sizeoffile,
		Backups:  numberofbackups,
		Age:      ageoffile,
		Compress: compress,
	}

	return LoggingConfig
}

func DefaultLoggingConfig() *LoggerConfig {
	return InitLoggerConfig(
		"./fiber_logger.log",
		"info", // loglevel info
		50,     // filesize is 50mb/file
		5,      // 5 files in total
		14,     // 14 days
		false)  // nocompression
}
