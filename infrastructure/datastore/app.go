package datastore

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/smartinsantos/go-auth-api/config"
	"log"
)

type AppDataStore struct {
	User 	UserDataStore
	db		*gorm.DB
}

// AppDataStore constructor
func New() (*AppDataStore, error) {
	env := config.Get()
	DBURL := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		env.DBConfig.User,
		env.DBConfig.Pass,
		env.DBConfig.Host,
		env.DBConfig.Port,
		env.DBConfig.Name)

	db, err := gorm.Open(env.DBConfig.Driver, DBURL)

	if err != nil {
		log.Println("Unable to open database")
		return nil, err
	}

	db.LogMode(true)

	ads := AppDataStore{
		User: UserDataStore{ db: db },
		db: db,
	}
	return &ads, nil
}

// Closes the  database connection
func (ds *AppDataStore) Close() error {
	return ds.db.Close()
}

// Migrates all tables
//func (ds *AppDataStore) AutoMigrate() error {
//	return ds.db.AutoMigrate(&entity.User{}).Error
//}