package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"
)

var source = "1#AVoir"
var dest = "Series"

var seriesTest = []struct {
	filename string // input
	expected string // path created to put the series in
}{
	{"fleche.4x02.IMMERSIF.mp4", "fleche/season04/fleche.4x02.immersif.mp4"},
	{"elementaire.302.IMMERSION.mp4", "elementaire/season03/elementaire.302.immersion.mp4"},
	{"Foire.du.Trone.S01E01.IMMERSE.720p.mkv", "foire.du.trone/season01/foire.du.trone.s01e01.immerse.720p.mkv"},
	{"Foire.du.Trone.S06E03.IMMERSIFASF.1080p/Foire.du.Trone.S06E03.IMMERSIFASF.1080p.avi", "foire.du.trone/season06/foire.du.trone.s06e03.immersifasf.1080p.avi"},
	{"marcel.agents.of.p.a.i.n.s03e19.720p.hdtv.x264-killers.[vtv].mkv", "marcel.agents.of.p.a.i.n/season03/marcel.agents.of.p.a.i.n.s03e19.720p.hdtv.x264-killers.[vtv].mkv"},
	{"normalnatural.1119.hdtv-lol[ettv].mkv", "normalnatural/season11/normalnatural.1119.hdtv-lol[ettv].mkv"},
	{"normalnatural.119.hdtv-lol.720p.[ettv].mkv", "normalnatural/season01/normalnatural.119.hdtv-lol.720p.[ettv].mkv"},
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
		os.MkdirAll(filepath.Dir(path), 0777) // needed for episode in sub folder
		_, err := os.Create(path)
		if err != nil {
			t.Fatalf("Unable to create file : %s", path)
		}
	}

	// actual testing

	destPath := fmt.Sprintf("%s/%s", tmpDir, dest)
	Sort(sourcePath, destPath) // the magic happen here
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

	//teardown
	err := os.RemoveAll(sourcePath)
	if err != nil {
		t.Errorf("could not remove : %s", sourcePath)
	}
	err = os.RemoveAll(destPath)
	if err != nil {
		t.Errorf("could not remove : %s", sourcePath)
	}
}

func TestSampleFile(t *testing.T) {
	tmpDir := os.TempDir()
	sourcePath := fmt.Sprintf("%s/%s", tmpDir, source)
	sampleFile := "2.broke.girls.417.hdtv-lol.sample.mp4"
	os.MkdirAll(sourcePath, 0777)
	_, err := os.Create(filepath.Join(sourcePath, sampleFile))
	if err != nil {
		t.Fatalf("Unable to create file : %s", sampleFile)
	}
	destPath := fmt.Sprintf("%s/%s", tmpDir, dest)
	Sort(sourcePath, destPath) // the magic happen here
	shouldExist := fmt.Sprintf("%s/%s", sourcePath, sampleFile)
	shouldNotExistAnymore := fmt.Sprintf("%s/%s", destPath, createNewSeriePath(sampleFile))
	if _, err = os.Stat(shouldExist); os.IsNotExist(err) {
		t.Errorf("%s should exist", shouldExist)
	}
	if _, err = os.Stat(shouldNotExistAnymore); err == nil {
		t.Errorf("%s shoud not exist anymore", shouldNotExistAnymore)
	}

	//teardown
	err = os.RemoveAll(sourcePath)
	if err != nil {
		t.Errorf("could not remove : %s", sourcePath)
	}
	err = os.RemoveAll(destPath)
	if err != nil {
		t.Errorf("could not remove : %s", sourcePath)
	}

}

func isDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// read in ONLY one file
	_, err = f.Readdir(1)

	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return true, nil
	}
	return false, err
}

func TestEmptyAfterCleaning(t *testing.T) {
	tmpDir := os.TempDir()
	sourcePath := fmt.Sprintf("%s/%s", tmpDir, source)
	file := "Splinter Cell Blacklist Shower.mp4"
	os.MkdirAll(sourcePath, 0777)
	_, err := os.Create(filepath.Join(sourcePath, file))
	if err != nil {
		t.Fatalf("Unable to create file : %s", file)
	}
	Clean(sourcePath)
	ok, err := isDirEmpty(sourcePath)
	if err != nil {
		t.Fatalf("Unable to determine if dir is empty : %s", err)
	}
	if !ok {
		t.Errorf("After cleaning dir should be empty : %s", sourcePath)
	}
}
