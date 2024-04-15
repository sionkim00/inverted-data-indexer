package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	utils "github.com/sionkim00/inverted-text-indexer/utils"
)

func main() {
	var dumpFilePath, query string

	flag.StringVar(&dumpFilePath, "p", "./data/wiki.xml.gz", "wikipedia abstract dump path")
	flag.StringVar(&query, "q", "Mexican nationalism in art music", "search query")
	flag.Parse()

	log.Printf("Search query: %s\n", query)
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

	// Create a outputfile
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatalf("Failed to create file: %v", err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Define headers for the CSV file
	headers := []string{"ID", "Document Text"}
	if err := writer.Write(headers); err != nil {
		log.Fatalf("Failed to write headers to CSV file: %v", err)
	}

	// Search the query from the indexes
	matchingIDs := idx.Search(query)
	log.Printf("Found %d documents in %v", len(matchingIDs), time.Since(start))
	for _, id := range matchingIDs {
		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)

		record := []string{ // Create a slice of strings for the record
			fmt.Sprintf("%d", id),
			doc.Text,
		}

		// Write the record to the CSV
		if err := writer.Write(record); err != nil {
			log.Fatalf("Failed to write record to CSV file: %v", err)
		}
	}
}
