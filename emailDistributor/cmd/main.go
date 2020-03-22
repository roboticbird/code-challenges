package main

import (
	"code-challenges/emailDistributor/internal/emaildistributor"
)

func main() {
	emails := []string{"one@gmail.com", "two@gmail.com", "three@gmail.com"}
	emaildistributor.DistributeEmails(emails)
}
