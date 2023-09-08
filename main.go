package main

import (
	"log"
	"os"

	"github.com/jacksonopp/send-email/mail"
	"github.com/joho/godotenv"
)

var (
	from       = "jackson@jopp.dev"
	recipients = []string{"nijolem507@horsgit.com"}
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("could not load env", err)
	}

	pw := os.Getenv("GMAIL_PW")

	sender := mail.NewGmailSender("Jackson Oppenheim", "jackson@jopp.dev", pw)

	if err := sender.SendEmail("test email", "this is the test email content", []string{"nijolem507@horsgit.com"}); err != nil {
		log.Fatal(err)
	}

}
