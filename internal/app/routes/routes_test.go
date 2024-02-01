package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorldRoute(t *testing.T) {
	router := gin.New()

	RegisterRoute(router)

	req, err := http.NewRequest("GET", "/hello-world", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}

	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	expectedBody := "Hello World Program"
	if res.Body.String() != expectedBody {
		t.Errorf("Expected body to be %s but got %s", expectedBody, res.Body.String())
	}
}
