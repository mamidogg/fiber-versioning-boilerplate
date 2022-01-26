package adapters

import (
	"database/sql"
	"fmt"
	"log"

	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ENV struct {
	Host         string
	Port         string
	DatabaseName string
	Username     string
	Password     string
}

// type MysqlAdapter interface {
// 	CloseMysqlAdapterAdapter()
// }

type MysqlAdapter struct {
	Connect *gorm.DB
	sql     *sql.DB
	env     ENV
}

// Close Connection DB
func (db *MysqlAdapter) CloseMysqlAdapter() {
	if err := db.sql.Close(); err != nil {
		fmt.Printf("Error closing db connection %s", err)
	} else {
		fmt.Println("DB connection closed")
	}
}

// NewMysqlServer is start connection database.
func NewMysqlServer(env ENV) *MysqlAdapter {
	// dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", env.Host, env.Port, env.Username, env.DatabaseName, env.Password)
	log.Println("--- DB connection ---", env)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", env.Username, env.Password, env.Host, env.Port, env.DatabaseName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Println(err)
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	sqlDB.SetMaxOpenConns(7)
	sqlDB.SetMaxIdleConns(5)
	log.Println("--- DB connection ---")
	return &MysqlAdapter{
		Connect: db,
		sql:     sqlDB,
		env:     env,
	}
}
