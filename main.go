package main

import (
	"flag"
	"log"
	"time"

	utils "github.com/sionkim00/inverted-text-indexer/utils"
)

func main() {
	var dumpFilePath, query string
	flag.StringVar(&dumpFilePath, "p", "./data/wiki.xml.gz", "wikipedia abstract dump path")
	flag.StringVar(&query, "q", "Members of parliament", "search query")
	flag.Parse()

	log.Println("Text search in progress")

	// Take the path and load files (documents) from the path
	start := time.Now()
	docs, err := utils.LoadDocuments(dumpFilePath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))

	// Create a new instance of index and send documents to Add function
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))

	// Search the query from the indexes
	matchingIDs := idx.Search(query)
	log.Printf("Found %d documents in %v", len(matchingIDs), time.Since(start))
	for _, id := range matchingIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)
	}
}
