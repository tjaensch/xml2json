package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func init() {
	os.Mkdir("./json", 0777)
}

func TestFindXmlFiles(t *testing.T) {
	expected := findXmlFiles()
	if len(expected) <= 0 {
		t.Error("Got ", len(expected))
	}
}

func TestProcessXmlFiles(t *testing.T) {
	xmlFiles := []string{
		"./testfiles/woa13_95A4_s03_04.xml",
		"./testfiles/KAQP_20091229v10001.xml",
		"./testfiles/NDBC_41009_201101_D991_v00.xml",
		"./testfiles/NOS_1770000_201504_D1_v00.xml",
	}
	processXmlFiles(xmlFiles)
	files,_ := ioutil.ReadDir("./json")
	if len(files) < len(xmlFiles) {
		t.Error("Expected %v, got %v ", len(xmlFiles), len(files))
	}
}

func TestXml2json(t *testing.T) {
	xml2json("./testfiles/NDBC_41008_201105_D2_adcp_v00.xml")
	if _, err := os.Stat("./json/woa13_95A4_s03_04.json"); os.IsNotExist(err) {
		t.Error("woa13_95A4_s03_04.xml should have been converted to ./json/woa13_95A4_s03_04.json")
	}
}