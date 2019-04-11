package main

import (
	"fmt"
	"os"
	"os/exec"
	"math"
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

func schlick(cosine, refractionIndex float64) float64{
	r0 := (1.0 - refractionIndex) / (1.0 + refractionIndex)
	r0 = r0*r0
	return r0 + (1.0 - r0) * math.Pow((1.0-cosine), 5)
}
