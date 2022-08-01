package weather

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/status", nil)
	w := httptest.NewRecorder()
	Status(w, req)
}
