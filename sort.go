package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

// all pattern should have 4 catch in it ( should make a better solution later)
var patterns = []string{"(.*?)[. -]S([0-9]{1,2})E([0-9]{1,2})(.*)", "(.*?)[.]([0-9])([0-9]{2})(.*)", "(.*?)[.]?([0-9]{1,2})x([0-9]{1,2})(.*)"}
var extensions = []string{".mp4", ".avi", ".mkv", ".srt"}

type serie struct {
	filename string
	name     string
	season   string
	episode  string
	version  string
}

func (s serie) twoDigitSeason() string {
	season := s.season
	if len(season) < 2 {
		return "0" + season
	}
	return season
}

func (s serie) seriePath() string {
	return fmt.Sprintf("%s/Season%s/%s", s.name, s.twoDigitSeason(), s.filename)
}

func isExtOk(file string) bool {
	for _, extension := range extensions {
		if extension == filepath.Ext(file) {
			return true
		}
	}
	return false
}

func createNewSeriePath(filename string) string {
	filename = filepath.Base(filename)
	for _, pattern := range patterns {
		r, _ := regexp.Compile(pattern)
		matches := r.FindStringSubmatch(filename)
		if matches != nil {
			s := serie{matches[0], matches[1], matches[2], matches[3], matches[4]}
			return s.seriePath()
		}

	}
	return ""
}

// Sort is a function that the source folder
// and move them to the correct place in the dest folder
func Sort(source string, dest string) {
	src, err := os.Open(source)
	defer src.Close()
	if err != nil {
		log.Fatalf("Could not open source folder : %s", err)
	}
	if filemode, _ := src.Stat(); !filemode.IsDir() {
		log.Fatalf("%s is not a directory", source)
	}
	files, err := src.Readdir(-1)
	if err != nil {
		log.Fatalf("Readding %s failed : %s", source, err)
	}
	for _, file := range files {
		filename := file.Name()
		if isExtOk(filename) {
			oldPath := filepath.Join(source, filename)
			newPath := createNewSeriePath(filename)
			newPath = filepath.Join(dest, newPath)
			os.MkdirAll(filepath.Dir(newPath), 0777)
			os.Rename(oldPath, newPath)
		} else if file.IsDir() {
			Sort(filepath.Join(source, filename), dest)
		}
	}
}
