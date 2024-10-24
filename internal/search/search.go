package search

import (
	"math"
	"value-index/internal/provider"
	"value-index/logger"
)

type SearchService struct {
	data   []int
	logger logger.Logger
}

func NewSearchService(repo provider.NumberProvider, logger logger.Logger) (*SearchService, error) {
	data, err := repo.Load()
	if err != nil {
		return nil, err
	}
	return &SearchService{data: data, logger: logger}, nil
}

// FindIndex searches for the exact match or closest index within 10% tolerance.
func (s *SearchService) FindIndex(value int) (int, int) {
	left, right := 0, len(s.data)-1

	// use binary search to find the index
	// it checks the middle of the array
	// and if the value is not found, it narrows down the search
	// by dividing the array in half
	for left <= right {
		mid := left + (right-left)/2
		if s.data[mid] == value {
			s.logger.Debug("exact match found")
			return mid, s.data[mid]
		} else if s.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	valueAbs := math.Abs(float64(value))

	// check if 'left' (higher) value is within 10% range
	if left < len(s.data) && math.Abs(math.Abs(float64(s.data[left]))-valueAbs) <= valueAbs*0.1 {
		s.logger.Debug("closest match found within 10% tolerance (left)")
		return left, s.data[left]
	}

	// check if 'right' (lower) value is within 10% range
	if right >= 0 && math.Abs(math.Abs(float64(s.data[right]))-valueAbs) <= valueAbs*0.1 {
		s.logger.Debug("closest match found within 10% tolerance (right)")
		return right, s.data[right]
	}

	// if no valid match found, return -1 and log a warning
	s.logger.Warn("no valid match found")
	return -1, -1
}
