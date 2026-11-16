package main

import (
    "fmt"
    "strings"
)

func main() {

    // Input paragraph
    paragraph := "Go is simple and Go is powerful"

    // Convert paragraph into words
    words := strings.Fields(paragraph)

    // Create a map to store word counts
    wordCount := make(map[string]int)

    // Iterate through words
    for _, word := range words {

        // Increase word count
        wordCount[word]++
    }

    // Display word frequencies
    fmt.Println("Word Frequency:")

    for word, count := range wordCount {
        fmt.Println(word, ":", count)
    }
}
