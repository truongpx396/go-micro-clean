package log

import (
	"fmt"
	"log"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

// NewLogger initializes a new zap.Logger with log rotation settings
func NewLogger(processID string, storageLocation string, rotationSize, rotationCount int, isJson, withStack bool) {
	// Configure lumberjack to handle log rotation by size and age
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: fmt.Sprintf("%s/%s.log", storageLocation, processID), // Log file path based on processID
		MaxAge:   rotationCount,                                        // Number of days to retain old log files
		MaxSize:  rotationSize,                                         // Rotate log when it reaches rotationSize MB
	})

	// Set up the core logging configuration
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // Use ISO 8601 format for human-readable timestamps
	encoderConfig.TimeKey = "time"
	encoderConfig.CallerKey = "caller"
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder // Use short caller format

	var encoder zapcore.Encoder
	if isJson {
		encoder = zapcore.NewJSONEncoder(encoderConfig) // Use JSON format for log entries
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig) // Use console (human-readable) format for log entries
	}
	// Set up the core logging configuration
	core := zapcore.NewCore(
		encoder,
		w,                 // Set log writer with rotation settings
		zapcore.InfoLevel, // Set minimum log level to Info
	)

	// Configure logger options
	options := []zap.Option{zap.AddCaller(), zap.AddCallerSkip(1)}
	if withStack {
		options = append(options, zap.AddStacktrace(zapcore.ErrorLevel))
	}

	// Return the logger with caller information and optional stack trace enabled
	Logger = zap.New(core, options...)
}

func Error(msg string, err error, fields ...interface{}) {
	Logger.Error(msg, zap.Error(err), zap.Any(" ", fields))
}

func Fatal(msg string, err error) {
	Logger.Error(msg, zap.Error(err))
	log.Fatalf("Fatal error: %s", err)
}

func Fatal1(msg string, err error, fields ...interface{}) {
	Logger.Error(msg, zap.Error(err), zap.Any(" ", fields))
	log.Fatalf("Fatal error: %s", err)
}

func Warn(msg string, fields ...interface{}) {
	Logger.Warn(msg, zap.Any(" ", fields))
}

func Info(msg string) {
	Logger.Info(msg)
}

func Info1(msg string, fields ...interface{}) {
	Logger.Info(msg, zap.Any(" ", fields))
}
