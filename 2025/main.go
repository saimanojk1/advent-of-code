package main

import (
	"advent-of-code-2025/pkg"
	"log/slog"
	"path/filepath"
)

func main() {
	password, err := pkg.Day1(filepath.Join(pkg.INPUTFILES_DIR, "day1-input.txt"))
	if err != nil {
		slog.Info("", "error", err.Error())
		panic("Failed to get password")
	}
	slog.Info("", slog.Int("password", password))
	invalidIdsSum, err := pkg.Day2(filepath.Join(pkg.INPUTFILES_DIR, "day2-input.txt"))
	if err != nil {
		slog.Info("", "error", err.Error())
		panic("Failed to get password")
	}
	slog.Info("", slog.Int("invalidIdsSums", invalidIdsSum))
}
