package data_bases

import (
	"context"
	"fmt"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

type SQLConnection struct {
	db *gorm.DB
}

func (s *SQLConnection) Connect() *gorm.DB {
	return s.db
}
func GetTestConnection() (*gorm.DB, error) {

	dialector := sqlite.Dialector{
		DriverName: "sqlite",
		DSN:        ":memory:",
	}

	db, err := gorm.Open(dialector, &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to test database (sqlite): %w", err)
	}

	return db, nil
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

func MustMigrate(db *gorm.DB, models ...any) {
	if err := db.AutoMigrate(models...); err != nil {
		panic(err)
	}
}
