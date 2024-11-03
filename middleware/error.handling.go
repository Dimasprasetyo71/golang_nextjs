package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	StatusCode   int
	ErrorMessage string
}

func (e *ErrorResponse) Error() string {
	return e.ErrorMessage
}

func NewErrorResponse(err error) *ErrorResponse {
	return &ErrorResponse{
		StatusCode:   500,
		ErrorMessage: err.Error(),
	}
}
func ErrorHandlerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Menjalankan handler yang akan dipanggil
		c.Next()

		// Cek apakah ada error
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			response := NewErrorResponse(err.Err)
			response.Render(c.Writer)
		}
	}
}

func NewCustomErrorResponse(statusCode int, errorMessage string) *ErrorResponse {
	return &ErrorResponse{
		StatusCode:   statusCode,
		ErrorMessage: errorMessage,
	}
}

func (e *ErrorResponse) Render(w http.ResponseWriter) {
	w.WriteHeader(e.StatusCode)

	if err := json.NewEncoder(w).Encode(e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
