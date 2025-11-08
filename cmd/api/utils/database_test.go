package utils_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"taska-core-me-go/cmd/api/config"
)

// estructura auxiliar
type table struct {
	Name  string
	Model interface{}
}

func GetTestConnection(t *testing.T) *gorm.DB {
	t.Helper()

	testDBName := config.Config.DB.Name + "_test"

	connString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.Config.DB.Host,
		config.Config.DB.Username,
		config.Config.DB.Password,
		testDBName,
		config.Config.DB.Port,
	)

	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		t.Fatalf("❌ failed to connect to test database: %v", err)
	}

	return db
}

func CreateTable(ctx context.Context, t *testing.T, db *gorm.DB, table interface{}) {
	err := db.AutoMigrate(table)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error creating table %v, %v", table, err), err)
		t.Errorf("❌ failed to create table: %v", err)
	} else {
		logger.Info(ctx, fmt.Sprintf("✅ Table created: %v", table))
	}
}

func DropTable(ctx context.Context, t *testing.T, db *gorm.DB, table interface{}) {
	err := db.Migrator().DropTable(table)
	if err != nil {
		logger.Error(ctx, fmt.Sprintf("Error dropping table %v, %v", table, err), err)
		t.Errorf("❌ failed to drop table: %v", err)
	} else {
		logger.Info(ctx, fmt.Sprintf("🗑️ Table dropped: %v", table))
	}
}

func TestDatabaseConnection(t *testing.T) {
	db := GetTestConnection(t)

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("❌ error getting SQL DB: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("❌ failed to ping test DB: %v", err)
	}

	t.Log("✅ Successfully connected to test database")
}
