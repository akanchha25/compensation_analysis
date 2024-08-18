package tests

// import (
//     "bytes"
//     "testing"
//     "compensation-api/repositories"
//     "github.com/stretchr/testify/assert"
//     "github.com/stretchr/testify/mock"
// 	"compensation-api/tests/mocks"
	
// )

// func TestSearch(t *testing.T) {
//     mockClient := new(mocks.MockElasticsearchClient)
//     mockClient.On("Search", mock.Anything).Return(`{
//         "hits": {
//             "hits": [
//                 {
//                     "_source": {
//                         "id": "1",
//                         "location": "Los Angeles, CA",
//                         "base_salary": 130000
//                     }
//                 }
//             ]
//         }
//     }`, nil)

//     repo := repositories.NewCompensationRepository(mockClient, "compensation_data")
//     query := bytes.NewBufferString(`{"query": {"match_all": {}}}`)
//     compensations, err := repo.Search(query)

//     assert.NoError(t, err)
//     assert.Len(t, compensations, 1)
//     assert.Equal(t, "Los Angeles, CA", compensations[0].Location)
//     assert.Equal(t, 130000, compensations[0].BaseSalary)
// }
