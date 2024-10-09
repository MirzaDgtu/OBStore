package apiserver

import (
	"fmt"
	"io"
	"obstore/internal/store/sqlstore"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Start(config *Config) error {
	gin.DisableConsoleColor()
	f, _ := os.Create("log\\gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	db, err := newDB(config.DatabaseURL)
	if err != nil {
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil
	}
	defer sqlDB.Close()

	sqlDB.SetMaxIdleConns(20)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(24 * time.Hour)

	store := sqlstore.New(db)
	srv := newServer(store)
	fmt.Println("Start OrdersBuid server")

	return srv.router.Run(config.BindAddr)
}

func newDB(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err
}
