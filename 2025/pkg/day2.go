package pkg

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

func Day2(inputFilePath string) (int, error) {
	invalidIdSum := 0
	inputFile, err := os.Open(inputFilePath)
	if err != nil {
		return 0, err
	}
	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	line := scanner.Text()
	slog.Info("", "line", line)
	for idRange := range strings.SplitSeq(line, ",") {
		startValueString, endValueString, _ := strings.Cut(idRange, "-")
		slog.Info("", "start Value", startValueString, "end Value", endValueString, "idRange", idRange)

		startValue, err := strconv.Atoi(startValueString)
		if err != nil {
			return invalidIdSum, err
		}
		endValue, err := strconv.Atoi(endValueString)
		if err != nil {
			return invalidIdSum, err
		}
		if startValue > endValue {
			return invalidIdSum, fmt.Errorf("invalid Ids format start value %d > end value %d", startValue, endValue)
		}

		if len(endValueString)%2 == 0 {
			mid := len(endValueString) / 2
			firstHalf := endValueString[:mid]
			currentDoubleNumber, err := strconv.Atoi(firstHalf + firstHalf)
			if err != nil {
				return invalidIdSum, err
			}
			for currentDoubleNumber >= startValue {
				slog.Info("", "currentDouble", currentDoubleNumber)
				if currentDoubleNumber <= endValue {
					slog.Info("", "currentDoubleInvalid", currentDoubleNumber)
					invalidIdSum += currentDoubleNumber
				}
				firstHalf = decrementString(firstHalf)
				currentDoubleNumber, err = strconv.Atoi(firstHalf + firstHalf)
				if err != nil {
					return invalidIdSum, err
				}
			}
		} else if len(startValueString)%2 == 0 {
			mid := len(startValueString) / 2
			firstHalf := startValueString[:mid]
			currentDoubleNumber, err := strconv.Atoi(firstHalf + firstHalf)
			if err != nil {
				return invalidIdSum, err
			}
			for currentDoubleNumber <= endValue {
				slog.Info("", "currentDouble", currentDoubleNumber)
				if currentDoubleNumber >= startValue {
					slog.Info("", "currentDoubleInvalid", currentDoubleNumber)
					invalidIdSum += currentDoubleNumber
				}
				firstHalf = incrementString(firstHalf)
				currentDoubleNumber, err = strconv.Atoi(firstHalf + firstHalf)
				if err != nil {
					return invalidIdSum, err
				}
			}
		}
	}
	return invalidIdSum, nil
}

func incrementString(valString string) string {
	val, _ := strconv.Atoi(valString)
	val++
	return strconv.Itoa(val)
}
func decrementString(valString string) string {
	val, _ := strconv.Atoi(valString)
	val--
	return strconv.Itoa(val)
}
