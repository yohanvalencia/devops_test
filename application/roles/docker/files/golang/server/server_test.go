package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	t.Run("it return status of the server", func(t *testing.T) {
		wanted := map[string]string{
			"status": "OK",
		}

		s := NewServer()
		request, _ := http.NewRequest(http.MethodGet, "/health", nil)
		response := httptest.NewRecorder()

		s.ServeHTTP(response, request)

		var got map[string]string

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response from server, '%v'", err)
		}
		assertStatus(t, response.Code, http.StatusOK)
		if !reflect.DeepEqual(got, wanted) {
			t.Errorf("got %v want %v", got, wanted)
		}
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
