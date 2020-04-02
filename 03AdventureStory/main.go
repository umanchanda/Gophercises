package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// Chapter is a chapter
type Chapter struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []Option `json:"options"`
}

// Option is an option
type Option struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

func main() {
	chapters := parseJSON("gopher.json")
	fmt.Println(chapters)
}

func parseJSON(filename string) []Chapter {
	jsonFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	gopher, _ := ioutil.ReadAll(jsonFile)

	var chapters []Chapter
	err = json.Unmarshal(gopher, &chapters)
	if err != nil {
		fmt.Println(err)
	}

	return chapters
}
