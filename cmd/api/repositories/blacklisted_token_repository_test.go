package repositories_test

import (
	"github.com/stretchr/testify/assert"
	models2 "taska-core-me-go/cmd/api/models"
	"taska-core-me-go/cmd/api/repositories"
	"taska-core-me-go/cmd/api/repositories/models"
	"taska-core-me-go/cmd/api/utils"
	"testing"
	"time"
)

func TestBlacklistedTokenRepository_FindByToken(t *testing.T) {
	assert := assert.New(t)

	conn, err := utils.GetTestConnection()
	assert.NoError(err)

	repository := &repositories.BlacklistedTokenRepository{Conn: conn}
	header := utils.GetTestRequestWithHeaders()
	ctx := utils.CreateRequest(header)

	token := "token-test"
	t.Run("Ok", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})
		utils.CreateTable(ctx, conn, models.BlacklistedTokenDb{})

		blacklistedTokenDb := models.BlacklistedTokenDb{
			ID:        1,
			UserID:    1,
			UserAgent: "test_agent",
			Token:     token,
		}
		repository.Conn.Create(&blacklistedTokenDb)

		user, err := repository.FirstByToken(ctx, token)
		assert.NoError(err)
		assert.Equal(int64(1), user.ID)
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})
	})

	t.Run("Error", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})

		blacklistedTokenDb, err := repository.FirstByToken(ctx, token)
		assert.Error(err)
		assert.Equal(int64(0), blacklistedTokenDb.ID)
		utils.DropTable(ctx, conn, models.UserDb{})
	})

	t.Run("ErrRecordNotFound", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})
		utils.CreateTable(ctx, conn, models.BlacklistedTokenDb{})

		blacklistedTokenDb, err := repository.FirstByToken(ctx, token)
		assert.Error(err)
		assert.Equal(int64(0), blacklistedTokenDb.ID)
		utils.DropTable(ctx, conn, models.UserDb{})
	})

}

func TestBlacklistedTokenRepository_FindByTokenNil(t *testing.T) {
	assert := assert.New(t)

	conn, err := utils.GetTestConnection()
	assert.NoError(err)

	repository := &repositories.BlacklistedTokenRepository{Conn: conn}
	header := utils.GetTestRequestWithHeaders()
	ctx := utils.CreateRequest(header)

	token := "token-test"

	t.Run("Ok", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})
		utils.CreateTable(ctx, conn, models.BlacklistedTokenDb{})

		blacklistedTokenDb := models.BlacklistedTokenDb{
			ID:        1,
			UserID:    1,
			UserAgent: "test_agent",
			Token:     token,
		}
		repository.Conn.Create(&blacklistedTokenDb)

		user, err := repository.FirstByTokenNil(ctx, token)
		assert.NoError(err)
		assert.Equal(int64(1), user.ID)
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})
	})

	t.Run("Error", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})

		blacklistedTokenDb, err := repository.FirstByTokenNil(ctx, token)
		assert.Error(err)
		assert.Equal(int64(0), blacklistedTokenDb.ID)
		utils.DropTable(ctx, conn, models.UserDb{})
	})

	t.Run("ErrRecordNotFound", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})
		utils.CreateTable(ctx, conn, models.BlacklistedTokenDb{})

		blacklistedTokenDb, err := repository.FirstByTokenNil(ctx, token)
		assert.Nil(err)
		assert.Equal(int64(0), blacklistedTokenDb.ID)
		utils.DropTable(ctx, conn, models.UserDb{})
	})

}

func TestBlacklistedTokenRepository_DeleteExpired(t *testing.T) {
	assert := assert.New(t)

	conn, err := utils.GetTestConnection()
	assert.NoError(err)

	repository := &repositories.BlacklistedTokenRepository{Conn: conn}
	header := utils.GetTestRequestWithHeaders()
	ctx := utils.CreateRequest(header)

	token := "token-test"
	before := time.Time{}
	t.Run("Ok", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})
		utils.CreateTable(ctx, conn, models.BlacklistedTokenDb{})

		createBlacklistedTokenDb := models.BlacklistedTokenDb{
			ID:        1,
			UserID:    1,
			UserAgent: "test_agent",
			Token:     token,
		}
		repository.Conn.Create(&createBlacklistedTokenDb)

		rowsAffected, err := repository.DeleteExpired(ctx, before)
		assert.NoError(err)
		assert.Equal(int64(0), rowsAffected)
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})
	})

	t.Run("Error", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})

		rowsAffected, err := repository.DeleteExpired(ctx, before)
		assert.Error(err)
		assert.Equal(int64(0), rowsAffected)
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})
	})

}

func TestBlacklistedTokenRepository_Save(t *testing.T) {
	assert := assert.New(t)

	conn, err := utils.GetTestConnection()
	assert.NoError(err)

	repository := &repositories.BlacklistedTokenRepository{Conn: conn}
	header := utils.GetTestRequestWithHeaders()
	ctx := utils.CreateRequest(header)

	token := "token-test"
	t.Run("Save_Ok", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})
		utils.CreateTable(ctx, conn, models.BlacklistedTokenDb{})

		createBlacklistedTokenDb := models.BlacklistedTokenDb{
			ID:        1,
			UserID:    1,
			UserAgent: "test_agent",
			Token:     token,
		}
		repository.Conn.Create(&createBlacklistedTokenDb)

		requestBlacklistedTokenDb := models2.BlacklistedToken{
			ID:        1,
			UserID:    1,
			UserAgent: "test_agent",
			Token:     token,
		}
		user, err := repository.Save(ctx, requestBlacklistedTokenDb)
		assert.NoError(err)
		assert.Equal(int64(1), user.ID)
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})
	})

	t.Run("Create_Ok", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})
		utils.CreateTable(ctx, conn, models.BlacklistedTokenDb{})

		createBlacklistedTokenDb := models.BlacklistedTokenDb{
			UserID:    1,
			UserAgent: "test_agent",
			Token:     "token-test-1",
			TokenType: "refreshs",
			IPAddress: "127.0.0.1",
			Reason:    "test_reason",
			CreatedAt: time.Now(),
		}
		repository.Conn.Create(&createBlacklistedTokenDb)

		requestBlacklistedTokenDb := models2.BlacklistedToken{
			UserID:    1,
			UserAgent: "test_agent",
			Token:     "token-test-0",
			TokenType: "refreshs",
			IPAddress: "127.0.0.1",
			Reason:    "test_reason",
			CreatedAt: time.Now(),
		}
		user, err := repository.Save(ctx, requestBlacklistedTokenDb)
		assert.NoError(err)
		assert.Equal(int64(2), user.ID)
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})
	})

	t.Run("Error", func(t *testing.T) {
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})

		requestBlacklistedTokenDb := models2.BlacklistedToken{
			UserID:    1,
			UserAgent: "test_agent",
			Token:     token,
		}
		user, err := repository.Save(ctx, requestBlacklistedTokenDb)
		assert.Error(err)
		assert.Equal(int64(0), user.ID)
		utils.DropTable(ctx, conn, models.BlacklistedTokenDb{})
	})

}
