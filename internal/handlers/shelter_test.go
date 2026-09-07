package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"cp_lab1/internal/repositories"

	"github.com/gin-gonic/gin"
)

func TestShelterHandlerReturnsErrorResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := NewShelterHandler(repositories.NewShelterRepoMemory(), repositories.NewCatRepoMemory())
	r.GET("/shelters/:id", h.GetByID)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/shelters/nope", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusBadRequest)
	}

	var body map[string]string
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("invalid JSON response: %v", err)
	}
	if body["error"] == "" {
		t.Fatalf("missing error response: %s", w.Body.String())
	}
}
