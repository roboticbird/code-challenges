package md5browser

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestBrowseListSucess(t *testing.T) {
	// mocking the url and hash
	restore := fetchUrlAndHash
	defer func() { fetchUrlAndHash = restore }()
	fetchUrlAndHash = func(urls <-chan string, sucess chan<- Result, fail chan<- ReportError) {
		for url := range urls {
			sucess <- Result{url, []byte("cool cats")}
		}
	}

	urlList := []string{"google.com", "adjust.com"}

	succeeded, failed := BrowseList(urlList, 1)

	if len(succeeded) != len(urlList) {
		t.Errorf("Expected %d sucessful results but got %d\n", len(urlList), len(succeeded))
	}
	if len(failed) != 0 {
		t.Errorf("Expected no failed results but got %d\n", len(failed))
	}
}

func TestBrowseListFail(t *testing.T) {
	// mocking the url and hash
	restore := fetchUrlAndHash
	defer func() { fetchUrlAndHash = restore }()
	fetchUrlAndHash = func(urls <-chan string, sucess chan<- Result, fail chan<- ReportError) {
		for url := range urls {
			fail <- ReportError{url, errors.New("problem")}
		}
	}

	urlList := []string{"google.com", "adjust.com"}

	succeeded, failed := BrowseList(urlList, 1)

	if len(succeeded) != 0 {
		t.Errorf("Expected no sucessful results but got %d\n", len(succeeded))
	}
	if len(failed) != len(urlList) {
		t.Errorf("Expected %d failed results but got %d\n", len(urlList), len(failed))
	}
}

func TestFetchUrlAndHash(t *testing.T) {
	restore := get
	defer func() { get = restore }()
	get = func(string) (*http.Response, error) {
		t := http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString("That's pretty cool")),
		}
		return &t, nil
	}

	urlList := []string{"google.com", "adjust.com"}

	sucess := make(chan Result, len(urlList))
	fail := make(chan ReportError, len(urlList))
	urls := make(chan string, len(urlList))

	// feed urls int the input channel
	for _, url := range urlList {
		urls <- url
	}
	close(urls)

	fetchUrlAndHash(urls, sucess, fail)

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
	if len(succeeded) != len(urlList) {
		t.Errorf("Expected %d sucessful results but got %d\n", len(urlList), len(succeeded))
	}
	if len(failed) != 0 {
		t.Errorf("Expected no failed results but got %d\n", len(failed))
	}
}

func TestEnsureProtocolScheme(t *testing.T) {
	tables := []struct {
		url      string
		expected string
	}{
		{"google.com", "http://google.com"},
		{"http://google.com", "http://google.com"},
		{"https://google.com", "https://google.com"},
		{"", "http://"},
	}
	for _, table := range tables {
		result, err := ensureProtocolScheme(table.url)
		if err != nil {
			t.Errorf("Encounted an error parsing url %s: %s\n", table.url, err)

		}
		if result != table.expected {
			t.Errorf("Protocol scheme was not correctly ensured, got: %s, expecting: %s\n", result, table.expected)
		}
	}
}
