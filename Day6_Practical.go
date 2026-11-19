
package main

import (
    "fmt"
    "net/http"
    "sync"
    "time"
)

func checkURL(url string, wg *sync.WaitGroup) {
    defer wg.Done()

    client := http.Client{
        Timeout: 5 * time.Second,
    }

    resp, err := client.Get(url)

    if err != nil {
        fmt.Printf("%s : ERROR: %v\n", url, err)
        return
    }

    defer resp.Body.Close()

    fmt.Printf("%s : %s\n", url, resp.Status)
}

func main() {
    urls := []string{
        "https://google.com",
        "https://github.com",
        "https://golang.org",
        "https://example.com",
    }

    var wg sync.WaitGroup

    for _, url := range urls {
        wg.Add(1)
        go checkURL(url, &wg)
    }

    wg.Wait()

    fmt.Println("All checks completed!")
}
