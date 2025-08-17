package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	apperrors "github.com/ShaimaaSabry/recipes/internal/domain/errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"

	//"gorm.io/gorm"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// If a handler already wrote a response, don't overwrite it.
		if c.Writer.Written() {
			return
		}

		if len(c.Errors) == 0 {
			return
		}

		// Take the last error (most recent). You could iterate to pick the "highest priority".
		err := c.Errors.Last().Err

		if handleJSONSyntaxError(c, err) {
			return
		}
		if handleJSONTypeError(c, err) {
			return
		}
		if handleValidationErrors(c, err) {
			return
		}
		if handleAppError(c, err) {
			return
		}

		handleUnknownError(c, err)
	}
}

/* ---------- Individual handlers ---------- */

// JSON syntax errors → {"errors":{"body":"invalid JSON at byte offset N"}}, 400
func handleJSONSyntaxError(c *gin.Context, err error) bool {
	var syn *json.SyntaxError
	if !errors.As(err, &syn) {
		return false
	}
	writeJSONFieldErrors(c, http.StatusBadRequest, map[string]string{
		"body": fmt.Sprintf("invalid JSON at byte offset %d", syn.Offset),
	})
	return true
}

// JSON type errors → {"errors":{"<field>":"invalid type: got \"X\", expected Y"}}, 400
func handleJSONTypeError(c *gin.Context, err error) bool {
	var typ *json.UnmarshalTypeError
	if !errors.As(err, &typ) {
		return false
	}
	field := typ.Field
	if field == "" {
		field = "body"
	}
	writeJSONFieldErrors(c, http.StatusBadRequest, map[string]string{
		field: fmt.Sprintf("invalid type: got %q, expected %s", typ.Value, typ.Type.String()),
	})
	return true
}

// {"errors":{"field":"message"}}, 400 (e.g., missing required field "name")
func handleValidationErrors(c *gin.Context, err error) bool {
	var verrs validator.ValidationErrors
	if !errors.As(err, &verrs) {
		return false
	}

	fields := make(map[string]string, len(verrs))
	for _, fe := range verrs {
		fields[fieldName(fe)] = humanizeValidationMsg(fe)
	}
	writeJSONFieldErrors(c, http.StatusBadRequest, fields)
	return true
}

func fieldName(fe validator.FieldError) string {
	// If the validator is configured to use json tags, this is already the json field.
	// If not, fall back to the struct field name but lower-case it.
	n := fe.Field()
	if tag := fe.StructField(); tag != "" && tag != n {
		n = tag
	}
	return n
}

func humanizeValidationMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "is required"
	case "min":
		return fmt.Sprintf("must be at least %s", fe.Param())
	case "max":
		return fmt.Sprintf("must be at most %s", fe.Param())
	case "oneof":
		return fmt.Sprintf("must be one of [%s]", fe.Param())
	default:
		return "is invalid"
	}
}

// App/domain errors → {"error":"<message>"}, 4xx
func handleAppError(c *gin.Context, err error) bool {
	switch {
	case errors.Is(err, apperrors.ErrInvalidInput):
		writeJSONError(c, http.StatusBadRequest, err.Error())
		return true
	case errors.Is(err, apperrors.ErrNotFound):
		writeJSONError(c, http.StatusNotFound, err.Error())
		return true
	case errors.Is(err, apperrors.ErrConflict):
		writeJSONError(c, http.StatusConflict, err.Error())
		return true
	default:
		return false
	}
}

// Everything else → {"error":"internal server error"}, 500
func handleUnknownError(c *gin.Context, _ error) {
	writeJSONError(c, http.StatusInternalServerError, "internal server error")
}

/* ---------- Small response helpers ---------- */

func writeJSONError(c *gin.Context, status int, message string) {
	c.AbortWithStatusJSON(status, gin.H{"error": message})
}

func writeJSONFieldErrors(c *gin.Context, status int, fields map[string]string) {
	c.AbortWithStatusJSON(status, gin.H{"errors": fields})
}
