package utils

import (
	"context"
	"fmt"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"taska-core-me-go/cmd/api/config"
)

type table struct {
	Name  string
	Model interface{}
}

type SQLConnection struct {
	db *gorm.DB
}

func (s *SQLConnection) Connect() *gorm.DB {
	return s.db
}

var tables = []table{
	//{Name: "users", Model: models.UserDb{}},
}

func GetTestConnection() (*gorm.DB, error) {
	testDBName := config.Config.DB.Name + "_test"
	connString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Config.DB.Host,
		config.Config.DB.Username,
		config.Config.DB.Password,
		testDBName,
		config.Config.DB.Port)
	conn, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	return conn, nil
}

func CreateTable(ctx context.Context, db *gorm.DB, table interface{}) {
	err := db.AutoMigrate(table)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error creating table %v, %v", table, err), err)
	} else {
		logger.Info(ctx, fmt.Sprintf("%v created", table))
	}
}

func DropTable(ctx context.Context, db *gorm.DB, table interface{}) {
	err := db.Migrator().DropTable(table)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error dropping table %v, %v", table, err), err)
	}
}
