# Instructions
Before you begin you must have Go installed and configured properly for your computer. Please see https://golang.org/doc/install

We would like you to demonstrate basic programming competency and clear communication skills.

We would like to see:

- A working tool, written in the Go programming language

- Some unit testing, you don't need to test every edge-case

- A short README describing the tool Please don't add any extra features to the tool.

Please don't include dependencies beyond Go's standard library.

Please submit the results as a link to a github repository.

# Task
You must build a tool which makes http requests and prints the address of the request along with the MD5 hash of the response.

- The tool must be able to perform the requests in parallel so that the tool can complete sooner. The order in which addresses are printed is not important.

- The tool must be able to limit the number of parallel requests, to prevent exhausting local resources.

The tool must accept a flag to indicate this limit, and it should default to 10 if the flag is not provided.

- The tool must have unit tests

- A README.md must be included describing the usage of this tool.


# Setup

- Clone git repo into `$GOPATH/src`

- Compile into a binary `$ go build -o myhttp $GOPATH/src/code-challenges/md5Browser/cmd/main.go`

- Run binary `$ ./myhttp -parallel <number of parallel workers> <list of urls>`

The parallel flag is optional. It allows you to set the number of parallel workers used to fetch
the urls. If it is not included then the number of parallel processes will match the number of urls
that were listed.

The list of urls will be fetched and returned with an md5 hash of their response body.

# Output
For each url given as input a line of output will be displayed in this format:
`<URL> <MD5 hash of response body>`

If there is an error fetching the url it will not be reported. This is because the problem
definition did not mention error reporting and stated that it did not want me to add additional
features. 

# Example

Example output of the program.

```
$ ./myhttp -parallel 3 google.com http://adjust.com https://facebook.com

 http://google.com 495cb13d1eb3c0245849b855f7b0d1de

 http://adjust.com 0535f7f447e79337585bb0b7685cbe06

 https://facebook.com 6cefb0296430b3c586bea81ab5ba9146
```

# Testing

- Run tests `$ cd $GOPATH/src/code-challenges/md5Browser/internal/md5browser; go test`

