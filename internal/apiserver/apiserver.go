package apiserver

import (
	"fmt"
	"obstore/internal/store/sqlstore"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Start(config *Config) error {
	db, err := newDb(config.BindAddr)
	if err != nil {
		return err
	}
	defer db.Close()

	store := sqlstore.New(db)
	srv := newServer(store)
	fmt.Println("Start OrdersBuid server")

}

func newDB(databaseURL string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(databaseURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, err

}
