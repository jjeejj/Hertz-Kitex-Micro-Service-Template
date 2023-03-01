package main

import (
	"log"
	"os"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"

	"github.com/jjeejj/Hertz-Kitex-Micro-Service-Template/server/cmd/auth/model"
)

// 创建数据表结果
func main() {
	dsn := "root:123456@tcp(localhost:3306)/FreeCar?charset=utf8mb4&parseTime=True&loc=Local"
	// global mode
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL Threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color printing
		},
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
			NameReplacer:  strings.NewReplacer("model_", ""),
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	_ = db.AutoMigrate(
		// 这里新增创建的 model
		// &hello.ModelHello{},
		&model.User{},
	)
	db.Migrator().RenameTable("model_hello", "hello")
}
