package main

import (
	"fmt"
	"os"
)

// Storage hash map structure for the generated urls
type Storage map[string]string

func newStorage() Storage {
	return make(Storage)
}

func (s Storage) add(shortURL string, fullURL string) string {
	s[shortURL] = fullURL

	return fmt.Sprintf("%s: %s", "OK", shortURL)
}

func (s Storage) read(shortURL string) string {
	i, ok := s[shortURL]

	if ok == false {
		fmt.Println(shortURL, ": not found")
		os.Exit(1)
	}

	return i
}
