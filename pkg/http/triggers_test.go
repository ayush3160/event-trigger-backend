package http

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    "testing"
    "go.uber.org/zap"
)


// Test generated using Keploy
func TestHandleError_LogsAndWritesError(t *testing.T) {
    logger, _ := zap.NewDevelopment()
    triggerService := &TriggerService{
        logger: logger,
    }

    w := httptest.NewRecorder()
    err := fmt.Errorf("test error")
    message := "Test error message"
    statusCode := http.StatusInternalServerError

    triggerService.handleError(w, err, message, statusCode)

    if w.Code != statusCode {
        t.Errorf("Expected status code %d, got %d", statusCode, w.Code)
    }

    if w.Body.String() != "test error\n" {
        t.Errorf("Expected error message 'test error', got '%s'", w.Body.String())
    }
}

// Test generated using Keploy
func TestWriteJSONResponse_ValidResponse(t *testing.T) {
    logger, _ := zap.NewDevelopment()
    triggerService := &TriggerService{
        logger: logger,
    }

    w := httptest.NewRecorder()
    data := map[string]string{"key": "value"}

    triggerService.writeJSONResponse(w, data)

    if w.Code != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
    }

    expectedBody := `{"key":"value"}`
    if w.Body.String() != expectedBody+"\n" {
        t.Errorf("Expected response body '%s', got '%s'", expectedBody, w.Body.String())
    }
}

