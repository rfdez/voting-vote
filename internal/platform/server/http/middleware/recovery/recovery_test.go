package recovery_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rfdez/voting-vote/internal/platform/server/http/middleware/recovery"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRecoveryMiddleware(t *testing.T) {
	// Setting up the Gin server
	gin.SetMode(gin.TestMode)
	engine := gin.New()
	engine.Use(recovery.Middleware())
	engine.GET("/test-middleware", func(context *gin.Context) {
		panic("something unexpected")
	})

	// Setting up the HTTP recorder and the request
	httpRecorder := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/test-middleware", nil)
	require.NoError(t, err)

	// Asserting the request does not produce a panic
	assert.NotPanics(t, func() {
		engine.ServeHTTP(httpRecorder, req)
	})
}
