// handlers/filter_strategy.go
package handlers

import "errors"

type FilterStrategy interface {
    Apply(key string, values []string) (interface{}, error)
}

type GteFilterStrategy struct{}

func (s *GteFilterStrategy) Apply(key string, values []string) (interface{}, error) {
    if len(values) == 0 {
        return nil, errors.New("missing value for gte filter")
    }
    return map[string]interface{}{
        "gte": values[0],
    }, nil
}

type MatchFilterStrategy struct{}

func (s *MatchFilterStrategy) Apply(key string, values []string) (interface{}, error) {
    if len(values) == 0 {
        return nil, errors.New("missing value for match filter")
    }
    return values[0], nil
}
