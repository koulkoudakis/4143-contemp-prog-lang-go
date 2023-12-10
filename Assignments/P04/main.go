// Sharome Burton

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
    // TODO
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
    // TODO
}

func main() {
    urls := []string{
        "https://unsplash.com/photos/Stki0y7Pepk/download?ixid=M3wxMjA3fDB8MXxhbGx8MTgzfHx8fHx8Mnx8MTcwMTEwOTgyOXw&force=true&w=640",
        "https://pixabay.com/get/gd82a8a399c8d636a2f57470aa88783f31f0f412886622cde2e1d43e38d7faa9bd2ab0841c07c67da98f4e85a231027458b90db89cce1dfc9b62ccb6c575d8bd36d9e2f98c4f153b007be80d55097cd9d_640.jpg",
        "https://unsplash.com/photos/a-man-in-a-hoodie-smoking-a-cigarette-etcN-HM_JRU",
        "https://pixabay.com/photos/cat-kitten-pet-kitty-young-cat-551554/".
        "https://cdn.stocksnap.io/img-thumbs/960w/food-food%20photography_GNAFKFNM2R.jpg",
        "https://unsplash.com/photos/a-green-and-red-bird-sitting-on-top-of-a-white-fence-vB6HQIerLMI",
        "https://unsplash.com/photos/M7iMdnG4R_g/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
        "https://images.unsplash.com/photo-1599474401061-465ef4b5bc05?auto=format&fit=crop&q=80&w=2070&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
        "https://images.pexels.com/photos/164729/pexels-photo-164729.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1",
        "https://cdn.stocksnap.io/img-thumbs/960w/deer-animal_NERJJRVKBO.jpg",
        "https://unsplash.com/photos/a-cup-of-hot-chocolate-with-marshmallows-in-it-s0Cnecr8W4U",
        "https://www.pexels.com/photo/aigle-19361893/",
        "https://www.pexels.com/photo/man-people-woman-vehicle-6763716/",
        "https://images.unsplash.com/photo-1610035974356-3e9f2c818347?q=80&w=1935&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
        "https://unsplash.com/photos/a-view-of-the-top-of-a-mountain-in-the-clouds-wdQ7DUGJmk8",
        "https://cdn.pixabay.com/photo/2017/05/11/12/35/girl-2304038_1280.jpg",
        "https://images.unsplash.com/photo-1692003996152-59e60ea85fe5?auto=format&fit=crop&q=80&w=1935&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D",
        "https://unsplash.com/photos/HQaZKCDaax0/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
        "https://cdn.pixabay.com/photo/2023/11/08/20/11/mountains-8375693_1280.jpg",
        "https://unsplash.com/photos/a-view-of-the-top-of-a-mountain-in-the-clouds-wdQ7DUGJmk8".
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
