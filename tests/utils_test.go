package tests

// import (
//     "testing"
//     "compensation-api/utils"
//     "github.com/stretchr/testify/assert"
// )

// func TestBuildQuery(t *testing.T) {
//     filters := map[string]string{
//         "salary[gte]": "120000",
//         "location":    "New York",
//     }
//     sortOptions := map[string]string{} // No sorting options provided

//     query := utils.BuildQuery(filters, sortOptions)

//     expected := `{"query":{"bool":{"filter":[{"range":{"salary":{"gte":120000}}},{"term":{"location":"New York"}}]}}}`
//     assert.JSONEq(t, expected, query.String())
// }
