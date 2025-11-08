package providers

import (
	"context"
	"fmt"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"taska-core-me-go/cmd/api/config"
	"taska-core-me-go/cmd/api/constants"
)

func DatabaseConnectionPostgres() (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	for retry := 0; retry < config.Config.DB.MaxConnectionRetries; retry++ {
		if db, err = GetDBConnectionPostgres(); err != nil {
			continue
		}

		sqlConn, _ := db.DB()
		if err = sqlConn.Ping(); err == nil {
			break
		}
	}

	return db, err
}

func GetDBConnectionPostgres() (*gorm.DB, error) {
	sslMode := config.Config.DB.SSLMode
	if sslMode == "" {
		sslMode = "require"
	}

	connString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=America/Bogota",
		config.Config.DB.Host,
		config.Config.DB.Username,
		config.Config.DB.Password,
		config.Config.DB.Name,
		config.Config.DB.Port,
		sslMode,
	)

	conn, err := gorm.Open(postgres.Open(connString), &gorm.Config{PrepareStmt: true, QueryFields: true})
	if err != nil {
		logger.Error(context.Background(), constants.ErrOpeningDBConnection, err)
		return nil, err
	}

	sqlDB, err := conn.DB()
	if err != nil {
		logger.Error(context.Background(), constants.ErrGettingDatabase, err)
		return nil, err
	}
	sqlDB.SetConnMaxLifetime(config.Config.DB.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(config.Config.DB.ConnMaxIdleTime)
	sqlDB.SetMaxIdleConns(config.Config.DB.MaxIdleConnections)
	sqlDB.SetMaxOpenConns(config.Config.DB.MaxOpenConnections)

	logger.Info(context.Background(), fmt.Sprintf(constants.OpenDBConnectionsMessage, sqlDB.Stats().OpenConnections))

	env := os.Getenv("GO_ENVIRONMENT")
	if env == "test" || env == "" {
		//test_utils.CreateTestDatabase(conn)
	}

	return conn, err
}
