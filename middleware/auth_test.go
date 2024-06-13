package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		username       string
		password       string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid credentials",
			username:       "user",
			password:       "pass",
			expectedStatus: http.StatusOK,
			expectedBody:   "OK",
		},
		{
			name:           "Invalid credentials",
			username:       "user",
			password:       "wrongpass",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"Invalid authorization token"}`,
		},
		{
			name:           "No credentials",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"error":"kepo"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			router := gin.New()
			router.Use(AuthMiddleware)
			router.GET("/test", func(ctx *gin.Context) {
				ctx.String(http.StatusOK, "OK")
			})

			req, _ := http.NewRequest(http.MethodGet, "/test", nil)
			if tt.username != "" || tt.password != "" {
				req.SetBasicAuth(tt.username, tt.password)
			}

			res := httptest.NewRecorder()

			router.ServeHTTP(res, req)

			require.Equal(t, tt.expectedStatus, res.Code)

			if tt.expectedStatus == http.StatusOK {
				require.Equal(t, tt.expectedBody, res.Body.String())
			} else {
				require.JSONEq(t, tt.expectedBody, res.Body.String())
			}
		})
	}
}
