package reader

import (
	"os"
)

type FileReader struct {
}

func (f FileReader) ReadAll(path string) (string, error) {
	data, err := os.ReadFile(path)
	return string(data), err
}
