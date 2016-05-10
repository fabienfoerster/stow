package main

import (
	"fmt"
	"os"
	"testing"
)

var source = "1#AVoir"
var dest = "Series"

var seriesTest = []struct {
	filename string // input
	expected string // path created to put the series in
}{
	{"fleche.4x02.IMMERSIF.mp4", "fleche/Season04/fleche.4x02.IMMERSIF.mp4"},
	{"elementaire.302.IMMERSION.mp4", "elementaire/Season03/elementaire.302.IMMERSION.mp4"},
	{"Foire.du.Trone.S01E01.IMMERSE.720p.mkv", "Foire.du.Trone/Season01/Foire.du.Trone.S01E01.IMMERSE.720p.mkv"},
}

func TestCreatingPath(t *testing.T) {
	for _, serie := range seriesTest {
		actual := createNewSeriePath(serie.filename)
		if actual != serie.expected {
			t.Errorf("CreateNewSeriePath(%s): expected %s, actual %s", serie.filename, serie.expected, actual)
		}
	}
}

func TestMovingSeries(t *testing.T) {
	// create fake episode for the tests
	tmpDir := os.TempDir()
	sourcePath := fmt.Sprintf("%s/%s", tmpDir, source)
	os.MkdirAll(sourcePath, 0777)
	for _, serie := range seriesTest {
		path := fmt.Sprintf("%s/%s", sourcePath, serie.filename)
		_, err := os.Create(path)
		if err != nil {
			t.Fatalf("Unable to create file : %s", path)
		}
	}

	// actual testing

	destPath := fmt.Sprintf("%s/%s", tmpDir, dest)
	Sort(sourcePath, destPath)
	for _, serie := range seriesTest {
		shouldExist := fmt.Sprintf("%s/%s", destPath, serie.expected)
		shouldNotExistAnymore := fmt.Sprintf("%s/%s", sourcePath, serie.filename)
		if _, err := os.Stat(shouldExist); os.IsNotExist(err) {
			t.Errorf("%s should exist", shouldExist)
		}
		if _, err := os.Stat(shouldNotExistAnymore); err == nil {
			t.Errorf("%s shoud not exist anymore", shouldNotExistAnymore)
		}
	}
}
