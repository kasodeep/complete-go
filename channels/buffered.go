package channels

/*
To create a buffered channel we can use the `make` function.
*/

/*
This function can run in the same go routine as the function that reads the buffer.
*/
func AddEmailsToQueue(emails []string) chan string {
	emailsToSend := make(chan string, len(emails))
	for _, email := range emails {
		emailsToSend <- email
	}
	return emailsToSend
}
