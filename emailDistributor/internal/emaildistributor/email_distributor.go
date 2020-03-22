package emaildistributor

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// sends emails out in parallel, returns list of emails that were sents and were failed to be sent
func DistributeEmails(emailList []string, workers int) ([]string, []string) {
	// set up input and output channels
	sucess := make(chan string)
	fail := make(chan string)
	emails := make(chan string, len(emailList))

	// we can't have less than 1 worker
	if workers <= 0 {
		workers = 1
	}

	// establish workers
	for w := 0; w < workers; w++ {
		go sendEmails(w, emails, sucess, fail)
	}

	// feed emails into input channel
	for _, email := range emailList {
		emails <- email
	}
	close(emails)

	// wait for all results to be returned
	succeeded := make([]string, 0, len(emailList))
	failed := make([]string, 0, len(emailList))
	for i := 0; i < len(emailList); i++ {
		select {
		case msg := <-sucess:
			succeeded = append(succeeded, msg)
		case msg := <-fail:
			log.Print(fmt.Sprintf("Failed to send email to: %s\n", msg))
			failed = append(failed, msg)
		}
	}
	return succeeded, failed
}

// emulating the time complexity of emails being send with a sleep
func sendEmails(id int, emails <-chan string, sucess chan<- string, fail chan<- string) {
	// read emails from input channel
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	for email := range emails {
		// sleep to emulate work done
		time.Sleep(500 * time.Millisecond)
		// emulate potential failures
		if r1.Intn(100000)%999999 == 0 {
			fail <- email
		} else {
			sucess <- email
		}
	}
}
