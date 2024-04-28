package database

import (
	"database/sql"
	"log"

	"github.com/2marks/csts/config"
	"github.com/go-sql-driver/mysql"
)

type MySQLDatabase struct{}

func (d *MySQLDatabase) NewSQLStorage() (*sql.DB, error) {
	var config mysql.Config = d.getConfig()
	db, err := sql.Open("mysql", config.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func (d *MySQLDatabase) getConfig() mysql.Config {
	return mysql.Config{
		User:                 config.Envs.DbUser,
		Passwd:               config.Envs.DbPassword,
		Addr:                 config.Envs.DbAddress,
		DBName:               config.Envs.DbName,
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}
}

func (d *MySQLDatabase) PingSQLStorage(db *sql.DB) error {
	if err := db.Ping(); err != nil {
		return err
	}

	return nil
}
