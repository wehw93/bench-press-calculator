package http

import (
	"bench_press_calculator/internal/model"
	"bytes"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockCalculatorService struct {
	mock.Mock
}

func (m *MockCalculatorService) Calculate(user *model.User, weight float32, quantity float32) (*model.Stat, error) {
	args := m.Called(user, weight, quantity)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Stat), args.Error(1)
}

func Test_HandlerCreate(t *testing.T) {
	mockSvc := new(MockCalculatorService)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError}))

	server := &Server{
		Logger: logger,
		Svc:    mockSvc,
	}
	reqBody := CreateRequest{
		Email:    "test@example.ru",
		Password: "pwd123",
		Weight:   100,
		Quantity: 10,
	}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/create", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	expectedStat := &model.Stat{
		MaxPress:      150.0,
		PersentBetter: 25.0,
	}
	mockSvc.On("Calculate", mock.MatchedBy(func(u *model.User) bool {
		return u.Email == reqBody.Email && u.Password != ""
	}), reqBody.Weight, reqBody.Quantity).Return(expectedStat, nil)

	w := httptest.NewRecorder()
	handler := server.Create()
	handler(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	var response CreateResponce
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, expectedStat.MaxPress, response.MaxWeight)
	assert.Equal(t, expectedStat.PersentBetter, response.PercentBetter)
}
