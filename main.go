package main

import (
	"flag"
	"log"
	"time"
)

func main() {
	var dumpFilePath, query string
	flag.StringVar(&dumpFilePath, "p", "./data/wiki.xml", "wikipedia abstract dump path")
	flag.StringVar(&query, "q", "Members of parliament", "search query")
	flag.Parse()

	log.Println("Text search in progress")

	// Load dump files (documents)
	start := time.Now()
	docs, err := utils.LoadDocuments(dumpFilePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

}
