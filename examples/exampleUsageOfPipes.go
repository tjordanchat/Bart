func main() {
	pipePath := "/tmp/my_named_pipe"

	// Create the named pipe
	err := createNamedPipe(pipePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.Remove(pipePath) // Clean up the pipe when done

	// Start a goroutine to read from the pipe
	go func() {
		data, err := readFromPipe(pipePath)
		if err != nil {
			fmt.Println("Reader error:", err)
			return
		}
		fmt.Printf("Received: %s\n", string(data))
	}()

	// Write to the pipe
	err = writeToPipe(pipePath, []byte("Hello from Go!"))
	if err != nil {
		fmt.Println("Writer error:", err)
	}

	// Give time for the reader to process
	// time.Sleep(1 * time.Second)
}
