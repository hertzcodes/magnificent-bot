package postgres

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBConnOptions struct {
	User     string
	Password string
	Host     string
	Port     uint
	DBName   string
	Schema   string
}

func (o DBConnOptions) PostgresDSN() string {
	return fmt.Sprintf(
		"user=%s password=%s host=%s port=%d dbname=%s search_path=%s sslmode=disable", o.User, o.Password, o.Host, o.Port, o.DBName, o.Schema,
	)
}

func NewPsqlGormConnection(opt DBConnOptions) (*gorm.DB, error) {
	return gorm.Open(postgres.Open(opt.PostgresDSN()), &gorm.Config{
		Logger: logger.Discard,
	})
}
