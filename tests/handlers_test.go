package tests

// import (
//     "net/http"
//     "net/http/httptest"
//     "testing"
//     "compensation-api/handlers"
//     "compensation-api/models"
//     "github.com/stretchr/testify/assert"
//     "github.com/stretchr/testify/mock"
//     "encoding/json"
// 	"compensation-api/tests/mocks"
// )

// func TestGetCompensations(t *testing.T) {
//     mockService := new(mocks.MockCompensationService)
//     mockService.On("GetCompensations", mock.Anything).Return([]models.Compensation{
//         {ID: "1", Location: "New York, NY", BaseSalary: 120000},
//     }, nil)

//     handler := handlers.NewCompensationHandler(mockService)

//     req, err := http.NewRequest("GET", "/compensation_data", nil)
//     if err != nil {
//         t.Fatal(err)
//     }

//     rr := httptest.NewRecorder()
//     handler.GetCompensations(rr, req)

//     assert.Equal(t, http.StatusOK, rr.Code)

//     var response []models.Compensation
//     if err := json.Unmarshal(rr.Body.Bytes(), &response); err != nil {
//         t.Fatal(err)
//     }

//     assert.Len(t, response, 1)
//     assert.Equal(t, "1", response[0].ID)
//     assert.Equal(t, "New York, NY", response[0].Location)
//     assert.Equal(t, 120000, response[0].BaseSalary)
// }
