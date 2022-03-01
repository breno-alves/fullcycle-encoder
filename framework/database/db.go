package database

import (
	"encoder/domain"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	Db            *gorm.DB
	Dsn           string
	DsnTest       string
	DbType        string
	DbTypeTest    string
	Debug         bool
	AutoMigrateDb bool
	Env           string
}

func NewDb() *Database {
	return &Database{}
}

func NewDbTest() *gorm.DB {
	dbInstance := NewDb()
	dbInstance.Env = "test"
	dbInstance.DbTypeTest = "sqlite3"
	dbInstance.DsnTest = ":memory:"
	dbInstance.AutoMigrateDb = true
	dbInstance.Debug = true

	conn, err := dbInstance.Connect()

	if err != nil {
		log.Fatalf("Test Db Error: %v", err)
	}

	return conn
}

func (d *Database) Connect() (*gorm.DB, error) {
	var err error

	if d.Env != "test" {
		d.Db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: d.Dsn,
		}), &gorm.Config{})
	} else {
		d.Db, err = gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	}

	if err != nil {
		return nil, err
	}

	if d.Debug {
		d.Db.Logger.LogMode(logger.Info)
	}

	if d.AutoMigrateDb {
		d.Db.AutoMigrate(&domain.Video{}, &domain.Job{})
	}

	return d.Db, nil
}
