## Program 4 - Concurrent Image Downloader
### Sharome Burton
### Description:

This program downloads images from a list of specified URLS in two ways. In one way, the images are downloaded sequentially. In the other way, a slice of URL's is passed to a function which then downloads them together at once using goroutines. The sequential and concurrent methods can then be compared.


### Files

|   #   | File                       | Description                                                |
| :---: | -------------------------- | ---------------------------------------------------------- |
|   1   | [main.go](./main.go)     | driver file.                                             |
|   2   | [go.mod](./go.mod)           | File defines path of Go module                   |
|   3   | [images](./images)           | Images downloaded for testing purposes from links in Google Sheet        |



### Instructions
- Download all files from P04 directory
- Execute `go run main.go` command
