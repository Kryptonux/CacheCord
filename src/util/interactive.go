package util

import (
	"fmt"
	"strings"
)

func SelectDirectoryFromList(directories []string) int {
	fmt.Println("Select a cache directory:")
	for i, dir := range directories {
		fmt.Printf("[%d] %s\n", i+1, dir)
	}

	var choice int
	fmt.Print("Enter the number of the directory: ")
	fmt.Scanln(&choice)

	return choice
}

func GetSelectedFormatsInteractively() map[string]bool {
	selectedFormats := map[string]bool{
		"png":   false,
		"jpg":   false,
		"jpeg":  false,
		"gif":   false,
		"webp":  false,
		"mp4":   false,
		"mp3":   false,
		"webm":  false,
		"gifv":  false,
		"woff":  false,
		"other": false,
	}

	fmt.Println("Select which formats you want to convert (y/n):")
	for format := range selectedFormats {
		fmt.Printf("[%s] Convert %s files? (y/n): ", format, format)
		var choice string
		fmt.Scanln(&choice)
		if strings.ToLower(choice) == "y" {
			selectedFormats[format] = true
		}
	}

	return selectedFormats
}
