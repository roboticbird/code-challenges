package main

import (
	"code-challenges/md5Browser/internal/md5browser"
	"flag"
	"fmt"
)

func main() {
	var workers int
	flag.IntVar(&workers, "parallel", 0, "number of parallel workers")
	flag.Parse()
	urls := flag.Args()

	// if the number of workers isn't set then it should match the number of requests
	if workers <= 0 {
		workers = len(urls)
	}

	// execute email distributor
	succeeded, _ := md5browser.BrowseList(urls, workers)

	// Print url and hash
	for _, sucess := range succeeded {
		fmt.Printf("\n%s %x\n", sucess.Url, sucess.Hash)
	}

	// errors are not reported. The problem deffinition asked not to add additional features. However
	// they are returned by the BrowseList function
}
