package main

import (
	"fmt"
	"github.com/clbanning/x2j"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	xmlFiles []string = findXmlFiles()
)

// Generic error checking function
func checkError(reason string, err error) {
	if err != nil {
		fmt.Printf("%s: %s\n", reason, err)
		os.Exit(1)
	}
}

// Create file directories
func prepDir() {
	os.Mkdir("./json", 0777)
}

func main() {
	log.Printf("Working digging up files...")

	t0 := time.Now()

	prepDir()

	var wg sync.WaitGroup

	// Start goroutine for each files segment of xmlFiles slice
	fileSegments := getFileSegments(xmlFiles)
	for _, fileSegment := range fileSegments {
		wg.Add(1)
		go func(fileSegment []string) {
			defer wg.Done()
			processXmlFiles(fileSegment)
		}(fileSegment)
	}

	// Wait until all goroutines finish
	wg.Wait()

	countOutputFiles()

	t1 := time.Now()
	fmt.Printf("\nThe program took %v seconds to run.\n", t1.Sub(t0).Seconds())
}

func findXmlFiles() []string {
	pathS, err := os.Getwd()
	checkError("get path failed", err)
	var files []string
	filepath.Walk(pathS, func(path string, f os.FileInfo, _ error) error {
		if !f.IsDir() {
			r, err := regexp.MatchString(".xml", f.Name())
			if err == nil && r {
				files = append(files, f.Name())
			}
		}
		return nil
	})
	return files
}

// Create fileSegments slice of slice for concurrent processing
func getFileSegments(xmlFiles []string) [][]string {
	var increaseRate int
	// Create a slice of xmlFiles
	fileSegments := make([][]string, 0)
	// Determine the length of the subslices based on amount of files and how many files can be open at the same time in PuTTY
	if len(xmlFiles) == 0 {
		fmt.Println("No XML files found in current working directory, program exiting.")
		os.Exit(1)
	}
	if len(xmlFiles) <= 250 {
		increaseRate = 1
	} else {
		increaseRate = len(xmlFiles) / 5
	}

	// Add subslices to fileSegments slice
	for i := 0; i < len(xmlFiles)-increaseRate; i += increaseRate {
		fileSegments = append(fileSegments, xmlFiles[i:i+increaseRate])
	}
	fileSegments = append(fileSegments, xmlFiles[len(xmlFiles)-increaseRate:])
	return fileSegments
}

func processXmlFiles(xmlFiles []string) {
	for _, xmlFile := range xmlFiles {
		xml2json(xmlFile)
	}
}

func xml2json(xmlFile string) {
	xmlInput, err := os.Open(xmlFile)
	checkError("open XML file failed", err)
	// Convert XML to JSON and pretty print
	jsonOutput, err := x2j.ToJsonIndent(xmlInput, false)
	checkError("x2j.ToJsonIndent failed", err)
	// Write JSON to file
	err = ioutil.WriteFile("./json/"+strings.TrimSuffix(filepath.Base(xmlFile), ".xml")+".json", []byte(jsonOutput), 0644)
	checkError("write JSON to file failed", err)
	if err == nil {
		fmt.Printf("%v.json written to JSON directory\n", strings.TrimSuffix(filepath.Base(xmlFile), ".xml"))
	}
}

func countOutputFiles() {
	files, err := ioutil.ReadDir("./json")
	checkError("counting output files failed, program exiting", err)
	log.Printf("\n%d files written to ./json directory\n", len(files))
}
