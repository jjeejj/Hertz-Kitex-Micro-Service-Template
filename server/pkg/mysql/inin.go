// Package mysql  init db connection
package mysql

import (
	"fmt"
	"strings"
	"time"

	"github.com/cloudwego/kitex/pkg/klog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/pkg/consts"
)

func InitDB(c Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(consts.MySqlDSN, c.User, c.Password, c.Host, c.Port, c.Name)
	newLogger := logger.New(
		logrus.NewWriter(), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL Threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // Disable color printing
		},
	)

	// global mode
	var err error
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NameReplacer:  strings.NewReplacer("model_", ""),
		},
		Logger: newLogger,
	})
	if err != nil {
		klog.Fatalf("init mysql failed: %s", err.Error())
		return nil, err
	}
	if err := db.Use(tracing.NewPlugin()); err != nil {
		klog.Fatalf("use tracing plugin failed: %s", err.Error())
		return nil, err
	}
	return db, nil
}
