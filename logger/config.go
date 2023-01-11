package logger

// configure the logging
type LoggerConfig struct {
	FileName string `mapstructure:"path"`     // filepath
	Size     int    `mapstructure:"max"`      // maximum size of file in megabyte
	Backups  int    `mapstructure:"backups"`  // number of backup files
	Age      int    `mapstructure:"age"`      // age of a file in days
	Compress bool   `mapstructure:"compress"` // compress the files
	Level    string `mapstructure:"level"`    // set the log level default is error
}

// global logger configuration
var LoggingConfig *LoggerConfig

// build a logger configuration
func InitLoggerConfig(filepath, loglevel string, sizeoffile, numberofbackups, ageoffile int, compress bool) *LoggerConfig {
	// the logging configuration
	return &LoggerConfig{
		FileName: filepath,
		Level:    loglevel,
		Size:     sizeoffile,
		Backups:  numberofbackups,
		Age:      ageoffile,
		Compress: compress,
	}

}

// build a default configuration
func DefaultLoggingConfig() *LoggerConfig {
	return InitLoggerConfig(
		"./fiber_logger.log", //file name
		"info",               // loglevel info
		50,                   // filesize is 50mb/file
		5,                    // 5 files in total
		14,                   // 14 days
		false)                // nocompression
}
