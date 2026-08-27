package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/guptakartike/qubit/internal/auth"
)

type mockRegistrationService struct {
	registerFunc func(
		ctx context.Context,
		req auth.RegisterRequest,
	) (auth.User, error)
}

func (m *mockRegistrationService) Register(
	ctx context.Context,
	req auth.RegisterRequest,
) (auth.User, error) {
	return m.registerFunc(ctx, req)
}

func setupTestRouter(handler *RegistrationHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)

	router := gin.New()
	handler.RegisterRoutes(router)

	return router
}

func TestRegistrationHandler_Register_Success(t *testing.T) {
	userID := uuid.New()

	mockService := &mockRegistrationService{
		registerFunc: func(
			ctx context.Context,
			req auth.RegisterRequest,
		) (auth.User, error) {
			return auth.User{
				ID:     userID,
				Name:   "Kartike Gupta",
				Email:  "kartike@example.com",
				Status: "active",
			}, nil
		},
	}

	handler := NewRegistrationHandler(mockService)
	router := setupTestRouter(handler)

	body := `{
		"name": "Kartike Gupta",
		"email": "KARTIKE@EXAMPLE.COM",
		"password": "testpassword123"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/auth/register",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusCreated {
		t.Fatalf(
			"status code = %d, want %d",
			recorder.Code,
			http.StatusCreated,
		)
	}

	// Verify full response body structure.
	var resp map[string]any
	if err := json.Unmarshal(recorder.Body.Bytes(), &resp); err != nil {
		t.Fatalf("response body is not valid JSON: %v", err)
	}

	if id, ok := resp["id"].(string); !ok || id == "" {
		t.Errorf("response id = %q, want non-empty UUID string", resp["id"])
	}

	if name, ok := resp["name"].(string); !ok || name != "Kartike Gupta" {
		t.Errorf("response name = %q, want %q", resp["name"], "Kartike Gupta")
	}

	if email, ok := resp["email"].(string); !ok || email != "kartike@example.com" {
		t.Errorf("response email = %q, want normalized %q", resp["email"], "kartike@example.com")
	}

	if status, ok := resp["status"].(string); !ok || status != "active" {
		t.Errorf("response status = %q, want %q", resp["status"], "active")
	}

	if _, hasPassword := resp["password"]; hasPassword {
		t.Error("response must not contain a password field")
	}
}

func TestRegistrationHandler_Register_InvalidJSON(t *testing.T) {
	mockService := &mockRegistrationService{
		registerFunc: func(
			ctx context.Context,
			req auth.RegisterRequest,
		) (auth.User, error) {
			t.Fatal("service should not be called")
			return auth.User{}, nil
		},
	}

	handler := NewRegistrationHandler(mockService)
	router := setupTestRouter(handler)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/auth/register",
		strings.NewReader(`invalid-json`),
	)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf(
			"status code = %d, want %d",
			recorder.Code,
			http.StatusBadRequest,
		)
	}
}

func TestRegistrationHandler_Register_InvalidInput(t *testing.T) {
	mockService := &mockRegistrationService{
		registerFunc: func(
			ctx context.Context,
			req auth.RegisterRequest,
		) (auth.User, error) {
			return auth.User{}, &auth.ValidationError{
				Field:   "name",
				Message: "is required",
			}
		},
	}

	handler := NewRegistrationHandler(mockService)
	router := setupTestRouter(handler)

	body := `{
		"name": "",
		"email": "test@example.com",
		"password": "testpassword123"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/auth/register",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf(
			"status code = %d, want %d",
			recorder.Code,
			http.StatusBadRequest,
		)
	}

	var resp map[string]any
	if err := json.Unmarshal(recorder.Body.Bytes(), &resp); err != nil {
		t.Fatalf("response body is not valid JSON: %v", err)
	}

	if resp["error"] != "validation_error" {
		t.Errorf("response error = %q, want %q", resp["error"], "validation_error")
	}

	if resp["field"] != "name" {
		t.Errorf("response field = %q, want %q", resp["field"], "name")
	}
}

func TestRegistrationHandler_Register_DuplicateEmail(t *testing.T) {
	mockService := &mockRegistrationService{
		registerFunc: func(
			ctx context.Context,
			req auth.RegisterRequest,
		) (auth.User, error) {
			return auth.User{}, auth.ErrEmailAlreadyExists
		},
	}

	handler := NewRegistrationHandler(mockService)
	router := setupTestRouter(handler)

	body := `{
		"name": "Second User",
		"email": "test@example.com",
		"password": "testpassword123"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/auth/register",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusConflict {
		t.Fatalf(
			"status code = %d, want %d",
			recorder.Code,
			http.StatusConflict,
		)
	}

	expected := `"error":"email already exists"`

	if !strings.Contains(recorder.Body.String(), expected) {
		t.Fatalf(
			"response body = %s, want duplicate email error",
			recorder.Body.String(),
		)
	}
}

func TestRegistrationHandler_Register_InternalError(t *testing.T) {
	mockService := &mockRegistrationService{
		registerFunc: func(
			ctx context.Context,
			req auth.RegisterRequest,
		) (auth.User, error) {
			return auth.User{}, errors.New("database unavailable")
		},
	}

	handler := NewRegistrationHandler(mockService)
	router := setupTestRouter(handler)

	body := `{
		"name": "Test User",
		"email": "test@example.com",
		"password": "testpassword123"
	}`

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/v1/auth/register",
		strings.NewReader(body),
	)
	req.Header.Set("Content-Type", "application/json")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusInternalServerError {
		t.Fatalf(
			"status code = %d, want %d",
			recorder.Code,
			http.StatusInternalServerError,
		)
	}
}
