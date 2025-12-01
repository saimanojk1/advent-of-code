package pkg

import (
	"bufio"
	"os"
	"strconv"
)

func Day1(inputFilePath string) (int, error) {
	zeroCounter := 0
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return 0, err
	}

	currentIndex := 50

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := scanner.Text()

		initialIndex := currentIndex
		moveIndex, err := strconv.Atoi(line[1:])
		numberOfRotations := moveIndex / 100
		zeroCounter += numberOfRotations
		moveIndex = moveIndex % 100
		if err != nil {
			return 0, err
		}
		if string(line[0]) == "L" {
			currentIndex = currentIndex - moveIndex
		} else {
			currentIndex = currentIndex + moveIndex
		}

		if currentIndex == 0 && moveIndex != 0 {
			zeroCounter++
		} else if currentIndex < 0 {
			if initialIndex != 0 {
				zeroCounter++
			}
			currentIndex = 100 + currentIndex
		} else if currentIndex > 99 {
			currentIndex = currentIndex - 100
			if initialIndex != 0 {
				zeroCounter++
			}
		}
	}
	return zeroCounter, nil
}
