package mysqlrepo

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"runtime"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v3"
	"github.com/golang-migrate/migrate/v3/database"
	"github.com/golang-migrate/migrate/v3/database/mysql"

	_ "github.com/golang-migrate/migrate/v3/source/file" // .
)

var (
	MultiStatement = "multiStatements=true"
)

func NewMigrateDBConnWithStr(driver, connStr string) (*sql.DB, error) {
	return NewSqlDBconn(driver, fmt.Sprintf("%s&%s", connStr, MultiStatement))
}
func NewMigrateDBConn(c DBConfig) (*sql.DB, error) {
	return NewSqlDBconn(c.Driver, fmt.Sprintf("%s&%s", c.DSNString(), MultiStatement))
}

func NewMigrateDriver(conn *sql.DB) (database.Driver, error) {
	return mysql.WithInstance(conn, &mysql.Config{})
}

// DefaultMigrate migrates sql scripts in two places
// windows
// file:///" + pwd[3:] + "\\..\\..\\db\\migration
// else
// file://%v/db/migration
// all the dependencies are inited with default values from environment variables
func DefaultMigrate() {
	pwd, _ := os.Getwd()
	var sourcename string
	if runtime.GOOS == "windows" {
		sourcename = "file:///" + pwd[3:] + "\\..\\..\\migration"
	} else {
		sourcename = fmt.Sprintf("file://%v/migration", pwd)
	}

	err := MigrateDB(DefaultDBConfig(), sourcename)
	if err != nil {
		log.Fatalf("migration failed: %s\n", err.Error())
	}

	log.Println("migration completed")
}

func MigrateDBWithDriver(driver database.Driver, sourcePath string) error {
	m, err := migrate.NewWithDatabaseInstance(sourcePath, "mysql", driver)
	if err != nil {
		return err
	}

	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func MigrateDBWithConn(conn *sql.DB, sourcePath string) error {
	dd, err := NewMigrateDriver(conn)
	if err != nil {
		return err
	}

	defer dd.Close()

	return MigrateDBWithDriver(dd, sourcePath)
}

func MigrateDB(c DBConfig, sourcePath string) error {
	conn, err := NewMigrateDBConn(c)
	if err != nil {
		return err
	}

	defer conn.Close()

	return MigrateDBWithConn(conn, sourcePath)
}
