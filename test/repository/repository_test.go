package repository

import (
	"os"
	"strings"
	"testing"
	"value-index/internal/provider"

	"github.com/stretchr/testify/assert"
)

func TestFileRepository_Load(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testdata")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	testData := "10\n20\n30\n40\n50\n"
	_, err = tempFile.WriteString(testData)
	assert.NoError(t, err)
	tempFile.Close()

	repo := provider.NewFileProvider(tempFile.Name())
	data, err := repo.Load()

	assert.NoError(t, err)
	assert.Equal(t, []int{10, 20, 30, 40, 50}, data)
}

func TestFileRepository_Load_Unsorted(t *testing.T) {
	tempFile, err := os.CreateTemp("", "testdata")
	assert.NoError(t, err)
	defer os.Remove(tempFile.Name())

	testData := "10\n30\n20\n40\n50\n"
	_, err = tempFile.WriteString(testData)
	assert.NoError(t, err)
	tempFile.Close()

	repo := provider.NewFileProvider(tempFile.Name())
	data, err := repo.Load()

	assert.Error(t, err)
	assert.Nil(t, data)
	assert.True(t, strings.HasPrefix(err.Error(), "file is not sorted"))
}
