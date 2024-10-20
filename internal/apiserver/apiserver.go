package apiserver

import (
	"fmt"
	"io"
	"log"
	"obstore/internal/model"
	"obstore/internal/store/sqlstore"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Start(config *Config) error {
	gin.DisableConsoleColor()
	f, _ := os.Create("log\\gin.log")
	//	ct := time.Now()
	//	f, _ := os.Create("log\\log-" + ct.Format(time.DateTime) + ".log")
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

	dbMigrate(db)

	fmt.Println("Start OrdersBuid server")

	return srv.router.Run(config.BindAddr)
}

func newDB(databaseURL string) (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, err
	}

	return db, err
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.User{}, &model.Team{}, &model.TeamComposition{},
		&model.Product{}, &model.Order{}, &model.OrderDetails{},
		&model.AssemblyOrder{}, &model.AssemblyOrderDetails{})
}
