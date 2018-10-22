// https://tour.go-zh.org/concurrency/10

package main

import (
	"fmt"
	"sync"
)

type urlStatus struct {
	v   map[string]bool
	mux sync.Mutex
}

var wg sync.WaitGroup

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher, url_status *urlStatus) {
	// 并发抓取 URL
	// 不重复抓取页面

	defer wg.Done() // paired to wg.Add(1)

	if depth <= 0 {
		return
	}
	url_status.mux.Lock() // lock Read
	_, ok := url_status.v[url]
	url_status.mux.Unlock()

	if !ok { // de-duplicate
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)
		for _, u := range urls {
			wg.Add(1)
			go Crawl(u, depth-1, fetcher, url_status) //
		}
		url_status.mux.Lock() // lock Write
		url_status.v[url] = true
		url_status.mux.Unlock()
	}
	return
}

func main() {
	wg.Add(1)
	url_status := urlStatus{v: make(map[string]bool)}
	Crawl("https://golang.org/", 4, fetcher, &url_status)
	wg.Wait() // till all done

	fmt.Println("\nCrawled urls:")
	for k, _ := range url_status.v {
		fmt.Println(k)
	}
}

// fakeFetcher 是返回若干结果的 Fetcher。
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

// fetcher 是填充后的 fakeFetcher。
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
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
