package utils

import (
	"context"
	"crypto/rand"
	"github.com/taskalataminfo2026/tool-kit-lib-go/pkg/response_capture"
	"math/big"
	"net/http"
	"taska-core-me-go/cmd/api/constants"
)

func GenerateUniqueCode(ctx context.Context) (string, error) {
	var (
		code string
		err  error
	)

	if code, err = GenerateRandomCode(); err != nil {
		return code, response_capture.NewErrorME(ctx, http.StatusInternalServerError, err, constants.ErrGeneratingAccessToken)
	}

	return code, nil
}

func GenerateRandomCode() (string, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(900000))
	if err != nil {
		return "", err
	}
	code := n.Int64() + 100000
	return ConvertInt64ToString(code), nil
}
