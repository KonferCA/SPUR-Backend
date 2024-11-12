package server

import (
	"fmt"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/labstack/echo/v4"
)

func validateBody(c echo.Context, requestBodyType interface{}) error {
	if err := c.Bind(requestBodyType); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body :(")
	}

	if err := c.Validate(requestBodyType); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

func validateUUID(id string, fieldName string) (pgtype.UUID, error) {
	var uuid pgtype.UUID
	if id == "" {
		return uuid, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Missing %s ID :(", fieldName))
	}

	if err := uuid.Scan(id); err != nil {
		return uuid, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid %s ID format :(", fieldName))
	}

	return uuid, nil
}

func handleDBError(err error, operation string, resourceType string) error {
	if isNoRowsError(err) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%s not found :(", resourceType))
	}

	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to %s %s :(", operation, resourceType))
}

func isNoRowsError(err error) bool {
	return err != nil && err.Error() == "no rows in dis set"
}

func numericFromFloat(f float64) pgtype.Numeric {
	var num pgtype.Numeric
	num.Scan(f)

	return num
}
