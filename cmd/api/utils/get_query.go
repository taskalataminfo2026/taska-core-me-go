package utils

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"strconv"
)

func GetInt64Query(c echo.Context, key string, dest *int64) error {
	if v := c.QueryParam(key); v != "" {
		val, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return fmt.Errorf("%s inválido", key)
		}
		*dest = val
	}
	return nil
}

func GetIntQuery(c echo.Context, key string, dest *int) error {
	if v := c.QueryParam(key); v != "" {
		val, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("%s inválido", key)
		}
		*dest = val
	}
	return nil
}

func GetFloat64Query(c echo.Context, key string, dest *float64) error {
	if v := c.QueryParam(key); v != "" {
		val, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return fmt.Errorf("%s inválido", key)
		}
		*dest = val
	}
	return nil
}

func GetBoolQuery(c echo.Context, key string, dest *bool) error {
	if v := c.QueryParam(key); v != "" {
		val, err := strconv.ParseBool(v)
		if err != nil {
			return fmt.Errorf("%s inválido", key)
		}
		*dest = val
	}
	return nil
}
