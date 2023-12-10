package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "path/filepath"
    "time"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
    // TODO: Implement sequential download logic
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
    // TODO: Implement concurrent download logic
}

func main() {
    urls := []string{
        "https://unsplash.com/photos/hvdnff_bieQ/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
        "https://unsplash.com/photos/HQaZKCDaax0/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
        "https://images.unsplash.com/photo-1698778573682-346d219402b5?ixlib=rb-4.0.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640",
        "https://unsplash.com/photos/Bs2jGUWu4f8/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640"
        // Add more image URLs
    }

    // Sequential download
    start := time.Now()
    downloadImagesSequential(urls)
    fmt.Printf("Sequential download took: %v\n", time.Since(start))

    // Concurrent download
    start = time.Now()
    downloadImagesConcurrent(urls)
    fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}

// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
    // TODO: Implement download logic
    return nil
}