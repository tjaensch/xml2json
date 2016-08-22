+++ XML2JSON COMMAND LINE TOOL +++

WHAT DOES IT DO?
Convert all XML files in the current working directory to JSON. Runs on 250 files concurrently which is configurable and scalable as needed

PREREQUISITES
Working Go installation

HOW TO RUN IT?
- Navigate to the folder with the XML files to be converted, then run "go run xml2json.go"
OR
- "go build xml2json.go" and then run the binary "./xml2json" in the folder with the XML files to be converted
OR
- "go install xml2json.go" and then navigate to XML folder and run "$GOPATH/bin/xml2json"

CAVEATS
Currently only works on XML folder without additional XML subfolders