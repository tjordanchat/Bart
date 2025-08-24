import (
	"fmt"
	"os"
	"syscall"
)

func createNamedPipe(path string) error {
	err := syscall.Mkfifo(path, 0666) // 0666 grants read/write permissions to all
	if err != nil && !os.IsExist(err) { // Handle case where pipe already exists
		return fmt.Errorf("failed to create named pipe: %w", err)
	}
	return nil
}
