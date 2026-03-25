package mysqlrepo

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // .
	"github.com/jmoiron/sqlx"
	"github.com/kelseyhightower/envconfig"
)

type DBConfig struct {
	Driver      string `json:"driver" envconfig:"DB_DRIVER" default:"mysql"`
	Host        string `json:"host" envconfig:"DB_HOST"`
	Port        string `json:"port" envconfig:"DB_PORT"`
	User        string `json:"user" envconfig:"DB_USER"`
	Pwd         string `json:"password" envconfig:"DB_PWD"`
	Database    string `json:"database" envconfig:"DB_DATABASE"`
	Suffix      string `json:"suffix" envconfig:"DB_SUFFIX"`
	MaxIdleConn int    `json:"max_idle_conn" envconfig:"DB_MAXIDLE" default:"3"`
	MaxOpenConn int    `json:"max_open_conn" envconfig:"DB_MAXOPEN" default:"5"`
}

func (dc *DBConfig) DSNString() string {
	return fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?%v",
		dc.User,
		dc.Pwd,
		dc.Host,
		dc.Port,
		dc.Database,
		dc.Suffix,
	)
}

func DefaultDBConfig() DBConfig {
	config := DBConfig{}
	envconfig.Process("", &config)
	return config
}

func NewSqlxDBConn(c DBConfig) (*sqlx.DB, error) {
	return NewSqlxDBWithConnStr(c.Driver, c.DSNString(), c.MaxIdleConn, c.MaxOpenConn)
}

func NewSqlxDBWithConnStr(driver, connStr string, maxIdle, maxOpen int) (*sqlx.DB, error) {
	dbConn, err := sqlx.Open(driver, connStr)
	if err != nil {
		return nil, err
	}

	dbConn.SetMaxIdleConns(maxIdle)
	dbConn.SetMaxOpenConns(maxOpen)

	return dbConn, nil
}

func NewSqlDBconn(driver, dsn string) (*sql.DB, error) {
	return sql.Open(driver, dsn)
}

func DefaultSqlxDBConn() (*sqlx.DB, error) {
	conf := DefaultDBConfig()
	return NewSqlxDBConn(conf)
}
