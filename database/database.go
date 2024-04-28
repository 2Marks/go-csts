package database

import (
	"database/sql"
	"log"
)

type Database struct{}

func NewDatabase() *Database {
	return &Database{}
}

func (d *Database) InitDatabase() (*sql.DB, error) {
	var mysqlDatabase MySQLDatabase

	db, err := mysqlDatabase.NewSQLStorage()
	if err != nil {
		log.Println("Could not connect to database", err.Error())
		return nil, err
	}

	err = mysqlDatabase.PingSQLStorage(db)
	if err != nil {
		log.Println("Could not ping database", err.Error())
		return nil, err
	}

	log.Println("DB: connected successfully")

	return db, nil

}
