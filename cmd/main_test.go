package main

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestParser(t *testing.T) {
	bytes, err := os.ReadFile("./example.json")
	if err != nil {
		panic(err)
	}
	var excel = &Excel{}
	json.Unmarshal(bytes, excel)
	fmt.Printf("-->", excel.Header.Title)
}
