package provider

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type FileProvider struct {
	input string
}

func NewFileProvider(fileName string) *FileProvider {
	return &FileProvider{input: fileName}
}

func (p *FileProvider) Load() ([]int, error) {
	file, err := os.Open(p.input)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data []int
	var prevValue int
	firstLine := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		value, err := strconv.ParseInt(line, 10, 32)
		if err != nil {
			return nil, err
		}

		if !firstLine && int(value) < prevValue {
			return nil, fmt.Errorf("file is not sorted %d < %d", value, prevValue)
		}

		data = append(data, int(value))
		prevValue = int(value)
		firstLine = false
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}
