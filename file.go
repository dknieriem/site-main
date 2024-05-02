package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func fileCreatePage(page *Page) error {
	filename := fmt.Sprintf("%s.json", page.Title)
	// Check if {page.title}.json already exists
	_, err := os.Open(filename)
	// If not, create it
	if err == nil {
		return errors.New("File already exists")
	}

	val, err := json.Marshal(page)

	err = os.WriteFile(filename, val, 0664)

	if err != nil {
		return err
	}

	return nil
}

func fileGetHomepage() (*Page, error) {
	filename := "content/homepage.json"
	file, err := os.Open(filename)
	file.Close()
	if err != nil {
		return nil, err
	}
	var data []byte

	data, err = os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	homepage := new(Page)
	err = json.Unmarshal(data, &homepage)
	if err != nil {
		return nil, err
	}

	return homepage, nil
}

func fileGetPage(pageID string) (*Page, error) {
	filename := fmt.Sprintf("content/pages/%s.json", pageID)
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	var data []byte

	_, err = file.Read(data)
	if err != nil {
		return nil, err
	}
	homepage := new(Page)
	err = json.Unmarshal(data, &homepage)
	if err != nil {
		return nil, err
	}

	return homepage, nil
}
