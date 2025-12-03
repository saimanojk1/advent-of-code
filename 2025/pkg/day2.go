package pkg

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
	"slices"
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

		mid := 0
		if len(endValueString)%2 == 0 {
			mid = len(endValueString) / 2
		} else {
			mid = (len(endValueString) + 1) / 2
		}
		firstHalf := endValueString[:mid]
		var repeatedValues []int
		AddRepeatedValuesInRange(firstHalf, startValue, endValue, &repeatedValues)

		slog.Info("", "repeatedValues", repeatedValues)
		for _, val := range repeatedValues {
			invalidIdSum += val
		}
	}
	return invalidIdSum, nil
}

func AddRepeatedValuesInRange(repeatStringPrefix string, minValue, maxValue int, repeatedValues *[]int) *[]int {
	//slog.Info("", "repeatStringPrefix", repeatStringPrefix)
	repeatStringPrefixLen := len(repeatStringPrefix)
	if repeatStringPrefixLen == 0 {
		return repeatedValues
	}

	repeatPrefixVal, _ := strconv.Atoi(repeatStringPrefix)
	for repeatPrefixVal > 0 {
		//slog.Info("", "repeatStringPrefix", repeatStringPrefix)
		minValPrefix, _ := strconv.Atoi(strconv.Itoa(minValue)[:len(repeatStringPrefix)-1])
		maxValPrefix, _ := strconv.Atoi(strconv.Itoa(maxValue)[:len(repeatStringPrefix)-1])
		if (repeatPrefixVal < minValPrefix) && (repeatPrefixVal > maxValPrefix) {
			return AddRepeatedValuesInRange(repeatStringPrefix[:repeatStringPrefixLen-1], minValue, maxValPrefix, repeatedValues)
		}
		for lenOfRepeatVals := len(strconv.Itoa(minValue)); lenOfRepeatVals <= len(strconv.Itoa(maxValue)); lenOfRepeatVals++ {
			repeatedValue, err := GetRepeatedValue(repeatStringPrefix, lenOfRepeatVals)
			if err != nil {
				//slog.Debug("", "err", err)
			} else {
				if minValue <= repeatedValue && repeatedValue <= maxValue {
					//slog.Info("", "repeatedValue", repeatedValue)
					if !slices.Contains(*repeatedValues, repeatedValue) {
						*repeatedValues = append(*repeatedValues, repeatedValue)
					}
				}
			}
		}
		repeatStringPrefix = decrementString(repeatStringPrefix)
		repeatPrefixVal, _ = strconv.Atoi(repeatStringPrefix)
	}
	return repeatedValues
}

func GetRepeatedValue(repeatString string, maxSize int) (int, error) {
	repeatStringSize := len(repeatString)
	if maxSize%repeatStringSize == 0 && maxSize/repeatStringSize > 1 {
		repeatedString := strings.Repeat(repeatString, maxSize/repeatStringSize)
		repeatedVal, err := strconv.Atoi(repeatedString)
		if err != nil {
			return 0, fmt.Errorf("not possible to repeat %v", err)
		}
		return repeatedVal, nil
	}
	return 0, fmt.Errorf("not possible to repeat")
}

func decrementString(valString string) string {
	val, _ := strconv.Atoi(valString)
	val--
	return strconv.Itoa(val)
}
