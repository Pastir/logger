package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	AccessLog *zap.Logger
	InfoLog   *zap.Logger
	ErrorLog  *zap.Logger
)

type ConfigLog struct {
	//Path to log dir
	LogDir string
	//log file name
	Filename string
	//Max size(Mb) log file
	MaxSize int
	//The number of log files to store
	MaxBackups int
	//storage time(days) of log files
	MaxAge int
	//do need to archive files?
	Compress   bool
	Debug      bool
	Warning    bool
	Info       bool
	Error      bool
	Critical   bool
	OutConsole bool
	OutFile    bool

	TimeKey    string
	LevelKey   string
	MessageKey string
}

func (c *ConfigLog) SetLog() *zap.Logger {
	access := zapcore.AddSync(&lumberjack.Logger{
		Filename:   c.LogDir + c.Filename,
		MaxSize:    c.MaxSize,
		MaxBackups: c.MaxBackups,
		MaxAge:     c.MaxAge, //days
		Compress:   c.Compress})

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			TimeKey:        c.TimeKey,
			LevelKey:       c.LevelKey,
			NameKey:        "",
			CallerKey:      "",
			MessageKey:     c.MessageKey,
			StacktraceKey:  "",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		}),
		access,
		zap.InfoLevel,
	)

	return zap.New(core)
}
