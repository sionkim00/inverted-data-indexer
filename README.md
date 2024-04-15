# Inverted Data Indexer

<img  src="https://github.com/sionkim00/inverted-data-indexer/blob/main/images/1.jpg?raw=true"  width="400"  height="200"  alt="logo">

This project is an implementation of an inverted data indexer using Go.

It allows you to index a collection of documents and perform efficient text searches on the indexed data.

The indexer is designed to handle large datasets and supports features like tokenization, stop-word removal, and stemming.

**[Demo video on Youtube](https://youtu.be/389qZhBWmUA)**

## Why Inverted Index?

The inverted index is a data structure used to store mapping from content keywords to their locations in a database file, or in this case, documents.

It is significantly faster than straightforward string matching, especially as data scales.

## Transforming Documents Into HashMap

<img  src="https://github.com/sionkim00/inverted-data-indexer/blob/main/images/2.jpg?raw=true"  width="400"  height="200"  alt="logo">

Given a set of documents, such as `{1: "a monkey on a tree eating banana", 2: "He is eating banana"}`, the inverted index transforms this into a structure like `{"a":[1], "monkey": [1], "banana": [1, 2], etc...}`.

### Transformation Sequence:

The sequence of operations to create an inverted index is as follows:

1.  **Indexing**: Initialize an empty index.
2.  **Tokenization**: Split the text of each document into individual terms.
3.  **Filtering**: Apply filters such as converting to lowercase, removing stop-words, and stemming to standardize the tokens.
4.  **Adding to Index**: Each word is added to the index with the document IDs where it appears.

## How We Use Our Index to Search

<img  src="https://github.com/sionkim00/inverted-data-indexer/blob/main/images/3.jpg?raw=true"  width="400"  height="200"  alt="logo">

To find documents relevant to a search query, the index is queried for each term in the search string.

Document IDs are retrieved and then combined using an intersection function to find common documents containing all the terms.

## Features

- **Data Indexing**: The indexer can efficiently index a large collection of documents by creating an inverted index data structure.
- **Text Search**: Once the data is indexed, you can perform text searches to retrieve relevant document IDs that contain the specified query terms.
- **Tokenization**: The indexer tokenizes the text data by breaking it down into individual words or terms.
- **Stop-word Removal**: Common stop-words (e.g., "the", "and", "a") are removed from the tokenized text to improve search accuracy and reduce index size.
- **Stemming**: Words are stemmed to their root form (e.g., "running" and "ran" are stemmed to "run") to improve search recall.

## Time & Space Complexity

### Time Complexity

1.  **Building the Index:**

- The process of tokenizing and filtering text is O(n), where n is the number of characters in all documents. Indexing each token generally has a time complexity of O(m), where m is the number of tokens, assuming constant time operations for hash table insertions.

2.  **Searching the Index:**

- Term lookup is O(1) per term using a hash table. Intersecting document lists for multiple terms has a complexity of **O(k\*d)**, where k is the number of query terms and d is the average number of document references per term.

### Space Complexity

1. **Index Storage:**

- The space complexity is **O(T\*d)**, where T is the number of unique tokens and d is the average number of document references per token.

## Usage

1. Clone the repository:

   git clone https://github.com/sionkim00/inverted-text-indexer

2. Download [dump file](https://dumps.wikimedia.org/enwiki/latest/enwiki-latest-abstract.xml.gz) from wikipedia and save it as `wiki.xml.gz` under data folder
3. Run the indexer `go run main.go`

The indexer will load the documents, index them, and then search for the specified query. The matching document IDs and text will saved into `output.csv` file.

## Project Structure

- `main.go`: The entry point of the application, responsible for parsing command-line arguments and orchestrating the indexing and search process.
- `tokenizer.go`: Contains functions for tokenizing and analyzing text data.
- `index.go`: Implements the inverted index data structure and provides methods for adding documents and searching for queries.
- `filter.go`: Includes functions for filtering tokens, such as converting to lowercase, removing stop-words, and stemming.
- `document.go`: Defines the document structure and provides a function for loading documents from a compressed file.

## Dependencies

This project relies on the following external dependencies:

- [github.com/kljensen/snowball](https://github.com/kljensen/snowball) - For stemming words to their root form.
