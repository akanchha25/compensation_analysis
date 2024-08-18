package tests

// import (
//     "testing"
//     "compensation-api/services"
//     "compensation-api/models"
//     "github.com/stretchr/testify/assert"
//     "github.com/stretchr/testify/mock"
// 	"compensation-api/tests/mocks"
// )

// func TestGetCompensations(t *testing.T) {
//     mockRepo := new(mocks.MockCompensationRepository)
//     mockRepo.On("Search", mock.Anything).Return([]models.Compensation{
//         {ID: "1", Location: "San Francisco, CA", BaseSalary: 150000},
//     }, nil)

//     service := services.NewCompensationService(mockRepo)
//     compensations, err := service.GetCompensations(nil)

//     assert.NoError(t, err)
//     assert.Len(t, compensations, 1)
//     assert.Equal(t, "San Francisco, CA", compensations[0].Location)
//     assert.Equal(t, 150000, compensations[0].BaseSalary)
// }
