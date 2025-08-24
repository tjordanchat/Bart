import (
	"fmt"
	"os"
)

func writeToPipe(path string, data []byte) error {
	file, err := os.OpenFile(path, os.O_WRONLY, 0)
	if err != nil {
		return fmt.Errorf("failed to open named pipe for writing: %w", err)
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return fmt.Errorf("failed to write to named pipe: %w", err)
	}
	return nil
}
