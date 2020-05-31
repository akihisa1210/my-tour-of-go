package main

import (
	"fmt"
	"sync"
)

// Fetcher is my fetcher type.
type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, c *SafeCounter) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer c.wg.Done()

	if depth <= 0 {
		return
	}

	if c.ExistKey(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)

	c.Inc(url)

	for _, u := range urls {
		c.wg.Add(1)
		go Crawl(u, depth-1, fetcher, c)
	}
	return
}

// SafeCounter can be accessed parallelly.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
	wg  sync.WaitGroup
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	c.v[key]++
	c.mux.Unlock()
}

// ExistKey returns how the key exist in the map.
func (c *SafeCounter) ExistKey(url string) bool {
	c.mux.Lock()
	defer c.mux.Unlock()
	if _, ok := c.v[url]; ok {
		return true
	}
	return false
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	c.wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher, &c)
	c.wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Packages fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Packages os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
