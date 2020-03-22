# Problem
Please write a program that sends emails to recipients from a large list (1 Mio entries) in a
performant way.
You do not need to send real emails but just fake the email sending by waiting for half a second.

# Setup

- Clone git repo into `$GOPATH/src`

- Compile into a binary `$ go build $GOPATH/src/code-challenges/emailDistributor/cmd/main.go`

- Run binary `$ ./main <emails.txt> <number of workers>`

The `emails.txt` is a file containing one email address on each line

The `number of workers` is an integer spesifying how many parallel workers are required to send out
the emails. This option allows you to scale the program to your needs.

# Run tests

There are a couple test files with different number of emails inside. 

example output:
```
$ ./main test/email1mil.txt 10000
2020/03/22 20:10:51 Failed to send email to: email177843@gmail.com
2020/03/22 20:10:55 Failed to send email to: email246790@gmail.com
2020/03/22 20:11:05 Failed to send email to: email446150@gmail.com
2020/03/22 20:11:06 Failed to send email to: email478248@gmail.com
2020/03/22 20:11:10 Failed to send email to: email552404@gmail.com
2020/03/22 20:11:16 Failed to send email to: email672720@gmail.com
-----Finished sending emails.-----
Successfully sent: 999994
Failed to send: 6
Number of workers: 10000
Execution time: 50.113478474s
```

The program emulates emails with a half second wait. Additionally it emulates errors based on a
random number. These failurs will be logged in the console. When the program has completed its task
of sending emails it will report how many where send sucessfully, how many failed, number of workers
used, and execution time.

