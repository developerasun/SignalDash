package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestApiHealth(t *testing.T) {
	r := gin.Default()
	r.GET("/api/health", Health)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/health", nil)
	r.ServeHTTP(w, req)

	require.True(t, w.Code == http.StatusOK)
}
