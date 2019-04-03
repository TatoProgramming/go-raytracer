package main

import (
	"fmt"
	"os"
	"os/exec"
)

func ConvertToPNG(path string) {
	cmd := "magick"
	args := []string{path + ".ppm", path + ".png"}
	if err := exec.Command(cmd, args...).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("Successfully converted ppm to pdf")
}
