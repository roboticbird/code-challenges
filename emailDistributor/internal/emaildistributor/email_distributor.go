package emaildistributor

import (
	"fmt"
	"time"
)

func DistributeEmails(emailList []string) {
	const workers = 5

	// set up input and output channels
	sucess := make(chan string)
	emails := make(chan string)

	// establish workers
	for w := 0; w < workers; w++ {
		go sendEmails(w, emails, sucess)
	}

	// feed emails into input channel
	for _, email := range emailList {
		emails <- email
	}
	close(emails)

	// wait for all results to be returned
	for range emailList {
		fmt.Println(<-sucess)
	}

}

// emulating the time complexity of emails being send with a sleep
func sendEmails(id int, emails <-chan string, sucess chan<- string) {
	// read emails from input channel
	for email := range emails {
		time.Sleep(time.Second)
		sucess <- email
	}
}
