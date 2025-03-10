package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProtect(t *testing.T) {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	t.Run("ok", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://example.com/foo", nil)
		req.Header.Set("Authorization", "good_key")
		rec := httptest.NewRecorder()
		Protect(handler, "good_key").ServeHTTP(rec, req)
		if rec.Result().StatusCode != http.StatusOK {
			t.Error("unexpected status code")
		}
	})
	t.Run("fail", func(t *testing.T) {
		req := httptest.NewRequest("GET", "http://example.com/foo", nil)
		rec := httptest.NewRecorder()
		Protect(handler, "good_key").ServeHTTP(rec, req)
		if rec.Result().StatusCode != http.StatusUnauthorized {
			t.Error("unexpected status code")
		}
	})
}
