package store

import (
	"fmt"
	"io"
	"os"
)

func SaveImage(readCloser io.ReadCloser, filepath string) error {
	defer func(readCloser io.ReadCloser) {
		err := readCloser.Close()
		if err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}(readCloser)
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer func(out *os.File) {
		err := out.Close()
		if err != nil {
			fmt.Printf("Error closing file: %s\n", err)
		}
	}(out)

	_, err = io.Copy(out, readCloser)
	if err != nil {
		return fmt.Errorf("failed to save image: %w", err)
	}
	return nil
}
