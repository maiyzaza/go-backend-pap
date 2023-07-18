package middleware

import (
	"PattayaAvenueProperty/models/apiErrors"
	"PattayaAvenueProperty/models/handler"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func ErrorHandle() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Content-Type", "application/json; charset=utf-8")
		c.Next()
		if len(c.Errors) > 0 {
			for _, err := range c.Errors {
				if errors.Is(err, io.ErrUnexpectedEOF) || errors.Is(err, io.EOF) || strings.Contains(err.Error(), io.EOF.Error()) {
					var errorMessages []string
					message := err.Error()
					errorMessages = append(errorMessages, message)
					c.JSON(http.StatusBadRequest,
						handler.Wrapper{
							StatusCode: http.StatusBadRequest,
							Message:    "EOF Errors",
							Error:      errorMessages,
						})
					return
				} else if uuid.IsInvalidLengthError(err.Err) {
					message := err.Error()
					c.JSON(http.StatusBadRequest,
						handler.Wrapper{
							StatusCode: http.StatusBadRequest,
							Message:    message,
							Error:      nil,
						})
					return
				}
				switch e := err.Err.(type) {
				case validator.ValidationErrors:
					var errorMessages []string
					for _, ve := range e {
						message := fmt.Sprintf("%s is %s", ve.Field(), ve.ActualTag())
						errorMessages = append(errorMessages, message)
					}
					c.JSON(http.StatusBadRequest,
						handler.Wrapper{
							StatusCode: http.StatusBadRequest,
							Message:    "Validation Errors",
							Error:      errorMessages,
						})
				case *json.SyntaxError:
					c.JSON(http.StatusBadRequest,
						handler.Wrapper{
							StatusCode: http.StatusBadRequest,
							Message:    "Json Errors",
							Error:      nil,
						})
				case *apiErrors.CustomError:
					c.JSON(e.StatusCode,
						handler.Wrapper{
							StatusCode: e.StatusCode,
							Message:    e.Message,
						})
				default:
					log.Println(e)
					c.JSON(http.StatusInternalServerError, handler.Wrapper{
						StatusCode: http.StatusInternalServerError,
						Message:    "Unexpected Errors",
						Error:      []string{},
					})
				}
			}
		}
	}
}
