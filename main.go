package main

import (
	"fmt"
	"log"
	"strconv"

	"nihongo-search/internal/dictionary"
	"nihongo-search/internal/search"
	"nihongo-search/internal/server"
)

func main() {
	dict := dictionary.New()

	// Load Kanji Data
	for i := 1; i <= 2; i++ {
		filename := "./data/ja/kanji_bank_" + strconv.Itoa(i) + ".json"
		if err := dict.LoadKanji(filename); err != nil {
			log.Fatalf("Error loading kanji data from %s: %v", filename, err)
		}
	}

	// Load JMDict Data
	for i := 1; i <= 75; i++ {
		filename := "./data/ja/term_bank_" + strconv.Itoa(i) + ".json"
		if err := dict.LoadJMDict(filename); err != nil {
			log.Fatalf("Error loading JMDict data from %s: %v", filename, err)
		}
	}
	fmt.Println("JMDict words loaded")

	searchService := search.NewService(dict)
	srv := server.NewServer(searchService)

	if err := srv.Start("8080"); err != nil {
		log.Fatal(err)
	}
}
