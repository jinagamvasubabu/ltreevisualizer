package database

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DbHelper interface {
	CreateDBConn() (*gorm.DB, error)
}

type dbHelper struct {
	PostgresURI string
}

//NewDBHelper Creates a New DB helper instance
func NewDBHelper(postgresURI string) DbHelper {
	return &dbHelper{
		postgresURI,
	}
}

//CreateDBConn Create/Open a new DB connection
func (d *dbHelper) CreateDBConn() (*gorm.DB, error) {
	conn, err := gorm.Open("postgres", d.PostgresURI)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("error while creating DB connection = %s", err.Error()))
	}
	return conn, nil
}
