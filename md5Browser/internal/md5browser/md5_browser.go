package md5browser

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

type Response struct {
	Url  string
	Hash []byte
}

type ReportError struct {
	Url    string
	Reason error
}

func BrowseList(urlList []string, workers int) ([]Response, []ReportError) {
	// set up input and output channels
	sucess := make(chan Response)
	fail := make(chan ReportError)
	urls := make(chan string, len(urlList))

	// we can't have less than one worker
	if workers <= 0 {
		workers = 1
	}

	// establish workers
	for w := 0; w < workers; w++ {
		go fetchUrlAndHash(urls, sucess, fail)
	}

	// feed urls int the input channel
	for _, url := range urlList {
		urls <- url
	}
	close(urls)

	// wait for all results to be returned
	succeeded := make([]Response, 0, len(urlList))
	failed := make([]ReportError, 0, len(urlList))
	for i := 0; i < len(urlList); i++ {
		select {
		case msg := <-sucess:
			succeeded = append(succeeded, msg)
		case msg := <-fail:
			log.Print(fmt.Sprintf("Failed to fetch url: %s\n Reason: %s\n", msg.Url, msg.Reason))
			failed = append(failed, msg)
		}
	}
	return succeeded, failed
}

func fetchUrlAndHash(urls <-chan string, sucess chan<- Response, fail chan<- ReportError) {
	for url := range urls {
		url, err := ensureProtocolScheme(url)
		if err != nil {
			fail <- ReportError{url, err}
			continue
		}

		// fetch reposne from url
		resp, err := http.Get(url)
		if err != nil {
			fail <- ReportError{url, err}
			continue
		}
		// read body
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fail <- ReportError{url, err}
			resp.Body.Close()
			continue
		}
		// return md5 hash of response
		hash := md5.New()
		hash.Write([]byte(body))
		sucess <- Response{url, hash.Sum(nil)}
		resp.Body.Close()
	}
}

func ensureProtocolScheme(url string) (string, error) {
	matched, err := regexp.MatchString(`.+://.+`, url)
	if err != nil {
		return url, err
	}
	if !matched {
		return "http://" + url, nil
	}
	return url, nil
}
