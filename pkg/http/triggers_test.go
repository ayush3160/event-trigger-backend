package http

import (
    "fmt"
    "io"
    "net/http"
    "net/http/httptest"
    "testing"
    "go.uber.org/zap"
    "go.mongodb.org/mongo-driver/mongo"
    "eventtrigger-backend/pkg/services/cronjobs"
    "eventtrigger-backend/pkg/models"
)


// Test generated using Keploy
func TestHandleError_LogsAndWritesErrorResponse(t *testing.T) {
    logger := zap.NewExample()
    defer logger.Sync()
    triggersCollection := &mongo.Collection{}
    cronJobSchedulerSvc := &cronjobs.CronScheduler{}
    eventsSvc := &EventsService{}
    service := NewTriggerService(logger, triggersCollection, cronJobSchedulerSvc, eventsSvc)

    w := httptest.NewRecorder()
    err := fmt.Errorf("test error")
    message := "Test error message"
    statusCode := http.StatusInternalServerError

    service.handleError(w, err, message, statusCode)

    response := w.Result()
    body, _ := io.ReadAll(response.Body)

    if response.StatusCode != statusCode {
        t.Errorf("Expected status code %d, got %d", statusCode, response.StatusCode)
    }

    if string(body) != "test error\n" {
        t.Errorf("Expected response body 'test error', got '%s'", string(body))
    }
}

// Test generated using Keploy
func TestCreateHttpJob_InvalidURL_ReturnsError(t *testing.T) {
    logger := zap.NewExample()
    defer logger.Sync()
    triggersCollection := &mongo.Collection{}
    cronJobSchedulerSvc := &cronjobs.CronScheduler{}
    eventsSvc := &EventsService{}
    service := NewTriggerService(logger, triggersCollection, cronJobSchedulerSvc, eventsSvc)

    trigger := models.Trigger{
        Name:       "invalid-trigger",
        Endpoint:   "://invalid-url",
        MethodType: "GET",
    }

    _, err := service.createHttpJob(trigger)
    if err == nil {
        t.Error("Expected error for invalid URL, got nil")
    }
}


// Test generated using Keploy
func TestWriteJSONResponse_ValidData_WritesResponse(t *testing.T) {
    logger := zap.NewExample()
    defer logger.Sync()
    triggersCollection := &mongo.Collection{}
    cronJobSchedulerSvc := &cronjobs.CronScheduler{}
    eventsSvc := &EventsService{}
    service := NewTriggerService(logger, triggersCollection, cronJobSchedulerSvc, eventsSvc)

    w := httptest.NewRecorder()
    data := map[string]string{"key": "value"}

    service.writeJSONResponse(w, data)

    response := w.Result()
    if response.StatusCode != http.StatusOK {
        t.Errorf("Expected status code %d, got %d", http.StatusOK, response.StatusCode)
    }

    body, _ := io.ReadAll(response.Body)
    expected := `{"key":"value"}`
    if string(body) != expected+"\n" {
        t.Errorf("Expected response body '%s', got '%s'", expected, string(body))
    }
}

