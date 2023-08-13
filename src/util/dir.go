package util

import (
	"log"
	"os"
	"path/filepath"
)

func ListDirectoryContents(directoryPath string) ([]string, error) {
	var contents []string

	files, err := os.ReadDir(directoryPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		contents = append(contents, file.Name())
	}

	return contents, nil
}

func FindCacheDirectories() []string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	possibleDirs := []string{
		filepath.Join(homeDir, ".config/discord/Cache/Cache_Data"),
		filepath.Join(homeDir, ".config/discordptb/Cache/Cache_Data"),
		filepath.Join(homeDir, ".config/discorddevelopment/Cache/Cache_Data"),
		filepath.Join(homeDir, ".config/discordcanary/Cache/Cache_Data"),
		// Add more potential cache directories here
	}

	var foundDirs []string

	for _, dir := range possibleDirs {
		_, err := os.Stat(dir)
		if err == nil {
			foundDirs = append(foundDirs, dir)
			LogVerbose("Found cache dirrectory: %s", dir)
		}
	}

	return foundDirs
}
