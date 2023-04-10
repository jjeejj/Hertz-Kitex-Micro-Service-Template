package log

import (
	"os"
	"path"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	kitexzap "github.com/kitex-contrib/obs-opentelemetry/logging/zap"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
)

// InitKLogger to init logrus.
func InitKLogger(logDirPath string, levelStr string) {
	// Customizable output directory.
	// logFilePath := consts.KlogFilePath
	if err := os.MkdirAll(logDirPath, 0o777); err != nil {
		panic(err)
	}

	// Set filename to date.
	logFileName := time.Now().Format("2006-01-02") + ".log"
	fileName := path.Join(logDirPath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			panic(err)
		}
	}
	// 默认的 log 不打印行号.
	logger := kitexzap.NewLogger()
	level := klog.Level(levelStr2iIntMap[levelStr])
	switch {
	case level >= klog.LevelInfo:
		// Provides compression and deletion.
		lumberjackLogger := &lumberjack.Logger{
			Filename:   fileName,
			MaxSize:    20,   // A file can be up to 20M.
			MaxBackups: 5,    // Save up to 5 files at the same time.
			MaxAge:     10,   // A file can exist for a maximum of 10 days.
			Compress:   true, // Compress with gzip.
		}
		logger.SetOutput(lumberjackLogger)
	case level >= klog.LevelTrace:
		logger = kitexzap.NewLogger(
			kitexzap.WithZapOptions(
				zap.AddCaller(),
				zap.AddCallerSkip(3),
				zap.Development(),
			))
		logger.SetOutput(os.Stdout)
	default:
		klog.Fatal("not support log level")
	}
	logger.SetLevel(level)
	klog.SetLogger(logger)
}
