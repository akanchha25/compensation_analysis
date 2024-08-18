package tests

// import (
// 	"compensation-api/server"
// 	"compensation-api/tests/mocks"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestStartServer(t *testing.T) {
//     mockHandler := new(mocks.MockCompensationHandler)
//     mockHandler.On("GetCompensations", mock.Anything, mock.Anything).Return()

//     r := server.NewRouter(mockHandler)

//     req, err := http.NewRequest("GET", "/health", nil)
//     if err != nil {
//         t.Fatal(err)
//     }

//     rr := httptest.NewRecorder()
//     r.ServeHTTP(rr, req)

//     assert.Equal(t, http.StatusOK, rr.Code, "Expected status code 200")
//     assert.Equal(t, "Server is healthy", rr.Body.String(), "Expected response body 'Server is healthy'")
// }
