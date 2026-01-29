package orm

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/glebarez/sqlite"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const (
	DIALECT_MYSQL  = "mysql"
	DIALECT_SQLITE = "sqlite"
)

type Config struct {
	Dialect      string
	SQLLog       bool
	MaxIdleConns int
	MaxOpenConns int
	Prefix       string
	Mysql        MysqlConfig
	Sqlite       SqliteConfig
}

type MysqlConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	Database string
	Charset  string
}

type SqliteConfig struct {
	DBPath string
}

func New(c *Config) *gorm.DB {
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             0, // 慢 SQL 阈值
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true, // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,
		},
	)

	var db *gorm.DB
	var err error

	gormConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{TablePrefix: c.Prefix, SingularTable: true},
		Logger:         logger,
	}

	if c.Dialect == "mysql" {
		conf := c.Mysql
		url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database, conf.Charset)
		fmt.Println(url)
		db, err = gorm.Open(mysql.Open(url), &gormConfig)
		if err != nil {
			panic(err)
		}
	} else if c.Dialect == "sqlite" {
		conf := c.Sqlite
		fmt.Println(conf.DBPath)

		// Ensure the directory exists before opening the DB
		if err := os.MkdirAll(filepath.Dir(conf.DBPath), os.ModePerm); err != nil {
			panic(fmt.Errorf("failed to create database directory: %w", err))
		}
		db, err = gorm.Open(sqlite.Open(conf.DBPath), &gormConfig)
		if err != nil {
			panic(err)
		}

		db.Exec("PRAGMA foreign_keys = ON;")
		db.Exec("PRAGMA journal_mode = WAL;")
		db.Exec("PRAGMA busy_timeout = 5000;")
	}

	if c.SQLLog {
		// db.LogMode(true)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)

	return db
}
