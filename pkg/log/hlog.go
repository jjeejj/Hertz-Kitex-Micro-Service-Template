package log

import (
	"os"
	"path"
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
	hertzzap "github.com/hertz-contrib/obs-opentelemetry/logging/zap"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitHLogger to init logrus.
// logDirPath 日志目录.
func InitHLogger(logDirPath string, levelStr string) {
	// Customizable output directory.
	// logFilePath := consts.HlogFilePath
	if err := os.MkdirAll(logDirPath, 0o777); err != nil {
		panic(err)
	}

	// Set filename to date
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logDirPath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}
	level := hlog.Level(levelStr2iIntMap[levelStr])
	// 默认的 log 不打印行号.
	logger := hertzzap.NewLogger()
	switch {
	case level >= hlog.LevelInfo:
		// Provides compression and deletion
		lumberjackLogger := &lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    20,   // A file can be up to 20M.
			MaxBackups: 5,    // Save up to 5 files at the same time.
			MaxAge:     10,   // A file can exist for a maximum of 10 days.
			Compress:   true, // Compress with gzip.
		}
		logger.SetOutput(lumberjackLogger)
	case level >= hlog.LevelTrace:
		logger = hertzzap.NewLogger(
			hertzzap.WithZapOptions(
				zap.AddCaller(),
				zap.AddCallerSkip(3),
				zap.Development(),
			))
		logger.SetOutput(os.Stdout)
	default:
		hlog.Fatal("not support log level")
	}
	logger.SetLevel(level)
	hlog.SetLogger(logger)
}
