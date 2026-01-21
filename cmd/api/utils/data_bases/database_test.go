package data_bases_test

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	models2 "taska-core-me-go/cmd/api/controllers/dto"
	"taska-core-me-go/cmd/api/models"
	models3 "taska-core-me-go/cmd/api/repositories/models"
	"taska-core-me-go/cmd/api/utils"
	"taska-core-me-go/cmd/api/utils/data_bases"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringFormatting(t *testing.T) {
	assert := assert.New(t)

	host := "localhost"
	user := "user"
	password := "pass"
	dbname := "test"
	port := "5432"

	expected := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbname, port)

	got := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, user, password, dbname, port)

	assert.Equal(expected, got)
}

func TestCreateTable_Ok(t *testing.T) {
	assert := assert.New(t)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx := utils.CreateRequestContext(c)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		"postgres",
		"Pilotof1988*",
		"postgres",
		"5432",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	assert.NoError(err, "should connect to Postgres test database")

	data_bases.CreateTable(ctx, db, &models3.SkillsDb{})

	assert.True(db.Migrator().HasTable(&models3.SkillsDb{}), "table 'skills' should exist")
}

func TestCreateTable_Error(t *testing.T) {
	assert := assert.New(t)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx := utils.CreateRequestContext(c)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		"postgres",
		"Pilotof1988*",
		"postgres",
		"5432",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	assert.NoError(err, "should connect to Postgres test database")

	// Ejecutar función y capturar error
	data_bases.CreateTable(ctx, db, &models.User{})

	// Verificar que la tabla se creó
	assert.True(db.Migrator().HasTable(&models.User{}), "table 'users' should exist")
}

func TestDropTable_Ok(t *testing.T) {
	assert := assert.New(t)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx := utils.CreateRequestContext(c)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		"postgres",
		"Pilotof1988*",
		"postgres",
		"5432",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	assert.NoError(err, "should connect to Postgres test database")

	data_bases.CreateTable(ctx, db, &models3.SkillsDb{})
	assert.True(db.Migrator().HasTable(&models3.SkillsDb{}), "table 'skills' should exist")

	data_bases.DropTable(ctx, db, &models3.SkillsDb{})
	assert.False(db.Migrator().HasTable(&models3.SkillsDb{}), "table 'skills' should not exist")
}

func TestDropTable_ForceError(t *testing.T) {
	assert := assert.New(t)

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctx := utils.CreateRequestContext(c)

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"localhost",
		"postgres",
		"Pilotof1988*",
		"postgres",
		"5432",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	assert.NoError(err)

	sqlDB, _ := db.DB()
	sqlDB.Close()

	data_bases.DropTable(ctx, db, &models2.ListSkillsResponseDto{})
}
