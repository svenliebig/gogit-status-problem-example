package main

import (
	"fmt"
	"log"
	"os"

	"github.com/go-git/go-git/v5"
)

func main() {
	dir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	r, err := git.PlainOpen(dir)

	if err != nil {
		log.Fatal(err)
	}

	w, err := r.Worktree()

	if err != nil {
		log.Fatal(err)
	}

	status, err := w.Status()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File in Status:")
	for key, s := range status {
		fmt.Printf("Key: %s, Status: %v\n", key, s)
	}
}
