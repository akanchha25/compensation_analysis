package mocks

import (
    "net/http"
    "github.com/stretchr/testify/mock"
)

type MockCompensationHandler struct {
    mock.Mock
}

func (m *MockCompensationHandler) GetCompensations(w http.ResponseWriter, r *http.Request) {
    m.Called(w, r)
    // Optionally write mock responses here if needed
}
