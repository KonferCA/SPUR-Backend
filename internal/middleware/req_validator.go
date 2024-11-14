package middleware

import (
	"net/http"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

// Struct solely exists to comply with Echo's interface to add a custom validator...
type RequestBodyValidator struct {
	validator *validator.Validate
}

func (rv *RequestBodyValidator) Validate(i interface{}) error {
	log.Info().Msgf("Validating struct: %+v\n", i)
	if err := rv.validator.Struct(i); err != nil {
		log.Error().Err(err).Msg("Validation error")
		return err
	}

	return nil
}

// Creates a new request validator that can be set to an Echo instance
// and used for validating request bodies with c.Validate()
func NewRequestBodyValidator() *RequestBodyValidator {
	return &RequestBodyValidator{validator: validator.New()}
}

// Middleware that validates the incoming request body with the given structType.
// Optionally pass any argument(s) like the fmt.Sprintf() method to customize
// a log message before ending the connection.
func ValidateRequestBody(structType reflect.Type, args ...interface{}) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			reqStruct := reflect.New(structType)

			if err := c.Bind(reqStruct.Interface()); err != nil {
				return echo.NewHTTPError(http.StatusBadRequest, "Invalid request body.")
			}

			if err := c.Validate(reqStruct.Interface()); err != nil {
				// this will let the global error handler handle
				// the ValidationError and get error string for
				// the each invalid field.
				return err
			}

			return next(c)
		}
	}
}
