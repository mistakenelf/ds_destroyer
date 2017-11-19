package main

import (
	"os"
	"path/filepath"
)

func main() {
	searchDir := "."
	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if !f.IsDir() {
			if f.Name() == ".DS_Store" {
				os.Remove(path)
			}
		}
		return nil
	})
}
