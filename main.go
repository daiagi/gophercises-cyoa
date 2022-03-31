package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string   `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option `json:"options"`
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

func parseJson(jsonData []byte) (story Story, err error) {
	err = json.Unmarshal(jsonData, &story)
	return
}

func readFile(path string) (content []byte, err error) {
	content, err = os.ReadFile(path)
	return

}

func main() {
	fileContent, err := readFile("./gopher.json")
	if err != nil {
		panic(err)
	}

	story, err := parseJson(fileContent)
	if err != nil {
		panic(err)
	}
	for _, v := range story {
		fmt.Println(v.Title)
		
	}

	

}
