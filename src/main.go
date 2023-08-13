package main

import (
	"flag"
	"os"
	"path/filepath"
	"xks/cachecord/util"
)

func main() {
	cacheDirFlag := flag.String("d", "", "Cache directory path")
	formatsFlag := flag.String("f", "", "Comma-separated list of formats to convert")
	outputFlag := flag.String("o", "./output", "Output folder")
	interactiveFlag := flag.Bool("i", false, "Enable interactive mode")
	verboseFlag := flag.Bool("v", false, "Enable verbose logging")

	flag.Parse()

	if *verboseFlag {
		util.EnableVerboseLogging()
	}

	util.LogVerbose("Starting CacheCord...")

	var selectedFormats map[string]bool
	if *interactiveFlag {
		selectedFormats = util.GetSelectedFormatsInteractively()
	} else if *formatsFlag != "" {
		formats := util.ParseFormats(*formatsFlag)
		selectedFormats = make(map[string]bool)
		for _, format := range formats {
			selectedFormats[format] = true
		}
	} else {
		util.LogError("No formats selected")
		return
	}

	outputFolder := *outputFlag
	if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
		err := os.Mkdir(outputFolder, 0755)
		if err != nil {
			util.LogError("Failed to create dir: %s", err)
			os.Exit(0)
		}
	}

	formatCount := make(map[string]int)

	var cacheDirectories []string
	if *cacheDirFlag != "" {
		cacheDirectories = []string{*cacheDirFlag}
	} else {
		util.LogVerbose("Scanning Cache Directories...")
		cacheDirectories = util.FindCacheDirectories()
	}

	if len(cacheDirectories) == 0 {
		util.LogError("No cache directories found!")
		return
	}

	for _, cacheDir := range cacheDirectories {
		contents, err := util.ListDirectoryContents(cacheDir)
		if err != nil {
			util.LogError("Error listing contents of %s: %s\n", cacheDir, err)
			continue
		}

		for _, content := range contents {
			cacheFilePath := filepath.Join(cacheDir, content)
			err := util.ConvertCacheToImage(cacheFilePath, outputFolder, selectedFormats, formatCount)
			if err != nil {
				util.LogError("Error converting %s: %s\n", content, err)
			}
		}
	}
}
