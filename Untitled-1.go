// cmd/minutil/main.go
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: minutil chemin_du_dossier")
		os.Exit(1)
	}

	rootDir := os.Args[1]

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		fmt.Println(path)
		return nil
	})

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
