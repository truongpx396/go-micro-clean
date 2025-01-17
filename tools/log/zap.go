package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

// NewLogger initializes a new zap.Logger with log rotation settings
func NewLogger(processID string, rotationSize int, rotationCount int) {
	// Configure lumberjack to handle log rotation by size and age
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename: fmt.Sprintf("./logs/%s.log", processID), // Log file path based on processID
		MaxAge:   rotationCount,                           // Number of days to retain old log files
		MaxSize:  rotationSize,                            // Rotate log when it reaches rotationSize MB
	})

	// Set up the core logging configuration
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()), // Use JSON format for log entries
		w,                 // Set log writer with rotation settings
		zapcore.InfoLevel, // Set minimum log level to Info
	)

	// Return the logger with caller information enabled
	Logger = zap.New(core, zap.AddCaller())
}

func Error(msg string, err error, fields ...interface{}) {
	Logger.Error(msg, zap.Error(err), zap.Any(" ", fields))
}

func Warn(msg string, fields ...interface{}) {
	Logger.Warn(msg, zap.Any(" ", fields))
}

func Info(msg string, fields ...interface{}) {
	Logger.Info(msg, zap.Any(" ", fields))
}
