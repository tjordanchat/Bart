import (
	"fmt"
	"io"
	"os"
)

func readFromPipe(path string) ([]byte, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("failed to open named pipe for reading: %w", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read from named pipe: %w", err)
	}
	return data, nil
}
