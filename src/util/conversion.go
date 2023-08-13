package util

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/h2non/filetype"
)

func ParseFormats(formatString string) []string {
	formats := strings.Split(formatString, ",")
	var cleanedFormats []string
	for _, f := range formats {
		cleaned := strings.TrimSpace(f)
		if cleaned != "" {
			cleanedFormats = append(cleanedFormats, cleaned)
		}
	}
	return cleanedFormats
}

func ConvertCacheToImage(cacheFilePath string, outputFolder string, selectedFormats map[string]bool, formatCount map[string]int) error {
	cacheData, err := os.ReadFile(cacheFilePath)
	if err != nil {
		return err
	}

	cacheStr := string(cacheData)
	match := regexp.MustCompile(`(https?:\/\/\w*\.\w*\.?\w+\/[A-Za-z\/0-9_\-\(\)]*(\.png|\.gif|\.jpg|\.jpeg|\.mp4|\.mp3|\.webm|\.webp|\.gifv|\.woff))`).FindStringSubmatch(cacheStr)
	if len(match) == 0 {
		return nil
	}

	url := match[0]
	filename, format := ParseURL(url)

	if selectedFormats[format] {
		outputFileName := filepath.Join(outputFolder, filename)
		outputFile, err := os.Create(outputFileName)
		if err != nil {
			return err
		}
		defer outputFile.Close()

		dataStart := strings.Index(cacheStr, url)
		data := cacheData[dataStart+len(url):]

		buffer := make([]byte, 512)
		copy(buffer, data)
		kind, _ := filetype.Match(buffer)
		if kind == filetype.Unknown {
			LogVerbose("Found an unknown file type %s\n", kind.Extension)
			return nil
		}

		_, err = outputFile.Write(data)
		if err != nil {
			return err
		}

		formatCount[format]++

		err = outputFile.Sync()
		if err != nil {
			return err
		}

		LogSuccess("Converted and saved: %s\n", outputFileName)
	}

	return nil
}
