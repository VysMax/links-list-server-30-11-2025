package entities

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	SavedLinksDir string
	LastNum       int
)

func ExtractLastNum(dir string) int {
	maxNum := 0

	files, err := os.ReadDir(SavedLinksDir)
	if err != nil {
		fmt.Printf("Error reading directory: %v\n", err)
	}

	re := regexp.MustCompile(`\d+`)

	for _, file := range files {
		match := re.FindStringSubmatch(file.Name())
		extractedNum := match[0]
		num, err := strconv.Atoi(extractedNum)
		if err != nil {
			fmt.Printf("cannot not convert '%s' to number in file name '%s'\n", extractedNum, file.Name())
			continue
		}

		if num > maxNum {
			maxNum = num
		}
	}
	return maxNum
}
