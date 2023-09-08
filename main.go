package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jacksonopp/send-email/mail"
	"github.com/joho/godotenv"
)

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/send", sendEmail)
	r.HandleFunc("/verify/{code}", verifyEmail)
	r.HandleFunc("/example", mail.ServeEmailTemplate("Steve"))

	log.Println("Runnning on :3000")
	http.ListenAndServe(":3000", r)
}

func sendEmail(w http.ResponseWriter, r *http.Request) {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("could not load env", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	pw := os.Getenv("GMAIL_PW")

	sender := mail.NewGmailSender("Jackson Oppenheim", "jackson@jopp.dev", pw)

	if err := sender.SendEmail("test email", "this is the test email content", []string{"nijolem507@horsgit.com"}); err != nil {
		log.Fatal(err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusNoContent)
}

func verifyEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	qp := r.URL.Query()
	code := vars["code"]
	fmt.Fprintf(w, "code: %s, %s", code, qp.Get("email"))
}
