package search

import (
	"testing"
	"value-index/internal/search"

	"github.com/stretchr/testify/assert"
	"value-index/logger"
)

type MockProvider struct{}

func (m *MockProvider) Load() ([]int, error) {
	return []int{0, 10, 20, 30, 40, 50, 60, 62, 70}, nil
}

func TestFindIndex_ExactMatch(t *testing.T) {
	mockProvider := &MockProvider{}
	mockLogger := logger.NewMockLogger()

	service, err := search.NewSearchService(mockProvider, mockLogger)
	assert.NoError(t, err)

	index, value := service.FindIndex(62)
	assert.Equal(t, 7, index)
	assert.Equal(t, 62, value)
}

func TestFindIndex_ClosestMatch_Lower(t *testing.T) {
	mockProvider := &MockProvider{}
	mockLogger := logger.NewMockLogger()

	service, err := search.NewSearchService(mockProvider, mockLogger)
	assert.NoError(t, err)

	index, value := service.FindIndex(63)
	assert.Equal(t, 7, index)
	assert.Equal(t, 62, value)
}

func TestFindIndex_ClosestMatch_Higher(t *testing.T) {
	mockProvider := &MockProvider{}
	mockLogger := logger.NewMockLogger()

	service, err := search.NewSearchService(mockProvider, mockLogger)
	assert.NoError(t, err)

	index, value := service.FindIndex(28)
	assert.Equal(t, 3, index)
	assert.Equal(t, 30, value)
}

func TestFindIndex_NoMatch(t *testing.T) {
	mockProvider := &MockProvider{}
	mockLogger := logger.NewMockLogger()

	service, err := search.NewSearchService(mockProvider, mockLogger)
	assert.NoError(t, err)

	index, value := service.FindIndex(100)
	assert.Equal(t, -1, index)
	assert.Equal(t, -1, value)
}

type NegativeMockProvider struct{}

func (m *NegativeMockProvider) Load() ([]int, error) {
	return []int{-100, -50, -10, 0, 10, 20, 30}, nil
}

func TestFindIndex_NegativeNumbers(t *testing.T) {
	mockProvider := &NegativeMockProvider{}
	mockLogger := logger.NewMockLogger()

	service, err := search.NewSearchService(mockProvider, mockLogger)
	assert.NoError(t, err)

	index, value := service.FindIndex(-10)
	assert.Equal(t, 2, index)
	assert.Equal(t, -10, value)

	index, value = service.FindIndex(-51)
	assert.Equal(t, 1, index)
	assert.Equal(t, -50, value)
}
