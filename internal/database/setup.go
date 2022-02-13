package database

import (
	"github.com/gabe565/matrimony/internal/database/models"
	"github.com/gabe565/matrimony/internal/datadir"
	flag "github.com/spf13/pflag"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

const DefaultFilename = "database.sqlite3"

var (
	Filepath string
)

func init() {
	flag.StringVar(&Filepath, "db-file", DefaultFilepath(), "SQLite database filename")
}

func DefaultFilepath() string {
	return datadir.Default + string(os.PathSeparator) + DefaultFilename
}

func Setup() (*gorm.DB, error) {
	var err error

	Filepath = datadir.ReplaceIfInDefault(Filepath)
	dsn := Filepath + "?cache=shared"

	l := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: l,
	})
	if err != nil {
		return db, err
	}

	err = db.AutoMigrate(&models.Guest{})
	if err != nil {
		return db, err
	}

	return db, nil
}
