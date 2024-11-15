package server

import (
	"fmt"
	"net/http"
	"time"

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
	if err == nil {
		return nil
	}

	if isNoRowsError(err) {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%s not found :(", resourceType))
	}

	fmt.Printf("Database error during %s %s: %v\n", operation, resourceType, err)

	return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Failed to %s %s :(", operation, resourceType))
}

func isNoRowsError(err error) bool {
	if err == nil {
		return false
	}
	errMsg := err.Error()
	return errMsg == "no rows in result set" ||
		errMsg == "no rows in dis set" ||
		errMsg == "scanning empty row"
}

func numericFromFloat(f float64) pgtype.Numeric {
	var num pgtype.Numeric
	num.Scan(f)
	return num
}

func validateNumeric(value string) (pgtype.Numeric, error) {
	var num pgtype.Numeric
	err := num.Scan(value)
	if err != nil {
		return num, echo.NewHTTPError(http.StatusBadRequest, "Invalid numeric value :(")
	}

	return num, nil
}

func validateTimestamp(timeStr string, fieldName string) (pgtype.Timestamp, error) {
	var ts pgtype.Timestamp
	if timeStr == "" {
		return ts, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Missing %s :(", fieldName))
	}

	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return ts, echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid %s format :(", fieldName))
	}

	ts.Time = parsedTime
	ts.Valid = true
	return ts, nil
}

func validateTimeRange(startTime, endTime pgtype.Timestamp) error {
	if !startTime.Valid || !endTime.Valid {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid time values :(")
	}

	if endTime.Time.Before(startTime.Time) {
		return echo.NewHTTPError(http.StatusBadRequest, "End time cannot be before start time :(")
	}

	return nil
}
