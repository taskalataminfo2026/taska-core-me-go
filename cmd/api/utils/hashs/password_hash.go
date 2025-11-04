package hashs

import (
	"context"
	"taska-core-me-go/cmd/api/constants"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/response_capture"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func GeneratePasswordHash(ctx context.Context, password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", response_capture.NewErrorME(ctx, http.StatusInternalServerError, err, constants.ErrProcessingPassword)
	}
	return string(hash), nil
}

func CheckPassword(password, hashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}
