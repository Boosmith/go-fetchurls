package main

import (
  "fmt"
	"io/ioutil"
	"log"
	"strings"
	"net/http"
	"time"
  "sort"
)

func main() {
    content := ReadFile("SampleListOfUrls.txt")
    urls := SliceString(content)
    FetchUrls(urls)
}

func ReadFile(filePath string) string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}

func SliceString(content string) []string {
    urls := strings.Split(content, "\n")
    for i := 1; i < len(urls) - 1; i++ {
        url := urls[i]
        urls[i] = url[:(len(url) - 1)]
    }
    return urls
  }

func FetchUrlContentLength(url string, integers chan int) {
    fmt.Printf("Fetching %s\n", url)
    timeout := time.Duration(5 * time.Second)
    transport := &http.Transport{
      DisableCompression: true,
    }
    client := http.Client{
        Timeout: timeout,
        Transport: transport,
    }
    resp, err := client.Get(url)
    if err != nil {

      integers <- 0
      return
    }
    contentLength := resp.ContentLength
    integers <- int(contentLength)
}

func FetchUrls (urls []string) (int, int, int) {
    integers := make(chan int)

    for i := 0; i < len(urls) - 1; i++ {
            go FetchUrlContentLength(urls[i], integers)
    }
    var totalContentLengths int
    var contentLengths []int
    for i := 0; i < len(urls) - 1; i++ {
         contentLength := <-integers
         if contentLength != 0 {
           contentLengths = append(contentLengths, contentLength)
         }
         totalContentLengths += contentLength
    }
    sort.SliceStable(contentLengths, func(i,j int) bool {return contentLengths[i] < contentLengths[j]})
    smallest := contentLengths[0]
    longest := contentLengths[len(contentLengths) - 1]

    fmt.Printf("Smallest content length is %d, longest is %d\n", smallest, longest)
    fmt.Printf("Total content length = %d\n", totalContentLengths)
    return smallest, longest, totalContentLengths
}
