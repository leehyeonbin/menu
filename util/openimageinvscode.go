package util

import (
	"fmt"
	"os/exec"
)

func OpenImageInVSCode(filename string) error {
	// VSCode에서 이미지 열기 명령 실행
	cmd := exec.Command("code", filename)
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to open image in VSCode: %w", err)
	}
	return nil
}
