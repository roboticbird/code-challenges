package md5browser

import (
	"crypto/md5"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Result struct {
	Url  string
	Hash []byte
}

type ReportError struct {
	Url    string
	Reason error
}

type getRequester interface {
	Req(url string) (*http.Response, error)
}

func BrowseList(urlList []string, workers int) ([]Result, []ReportError) {
	// set up input and output channels
	sucess := make(chan Result)
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
	succeeded := make([]Result, 0, len(urlList))
	failed := make([]ReportError, 0, len(urlList))
	for i := 0; i < len(urlList); i++ {
		select {
		case msg := <-sucess:
			succeeded = append(succeeded, msg)
		case msg := <-fail:
			failed = append(failed, msg)
		}
	}
	return succeeded, failed
}

var fetchUrlAndHash = func(urls <-chan string, sucess chan<- Result, fail chan<- ReportError) {
	for url := range urls {
		url, err := ensureProtocolScheme(url)
		if err != nil {
			fail <- ReportError{url, err}
			continue
		}

		// fetch reposne from url
		//resp, err := http.Get(url)
		resp, err := get(url)
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
		sucess <- Result{url, hash.Sum(nil)}
		resp.Body.Close()
	}
}

var get = func(url string) (*http.Response, error) {
	return http.Get(url)
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
