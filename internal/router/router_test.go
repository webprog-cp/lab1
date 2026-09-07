package router

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestRequireAPIKey(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.POST("/protected", requireAPIKey, func(c *gin.Context) { c.Status(http.StatusNoContent) })

	w := httptest.NewRecorder()
	r.ServeHTTP(w, httptest.NewRequest(http.MethodPost, "/protected", nil))
	if w.Code != http.StatusUnauthorized {
		t.Fatalf("status without key = %d, want %d", w.Code, http.StatusUnauthorized)
	}

	w = httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/protected", nil)
	req.Header.Set("X-API-Key", apiKey)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusNoContent {
		t.Fatalf("status with key = %d, want %d", w.Code, http.StatusNoContent)
	}
}
