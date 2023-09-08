package mail

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

const (
	smtpAuthAddress   = "smtp.gmail.com"
	smtpServerAddress = "smtp.gmail.com:587"
)

type EmailSender interface {
	SendEmail(subject, content string, to []string) error
}

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(name, fromEmailAddress, fromEmailPassword string) EmailSender {
	return &GmailSender{
		name, fromEmailAddress, fromEmailPassword,
	}
}

func (s *GmailSender) SendEmail(subject, content string, to []string) error {
	e := email.NewEmail()
	e.From = fmt.Sprintf("%s <%s>", s.name, s.fromEmailAddress)
	e.Subject = subject
	c, err := s.BuildTemplate("steve")
	if err != nil {
		return err
	}
	e.HTML = c
	e.To = to

	auth := smtp.PlainAuth("", s.fromEmailAddress, s.fromEmailPassword, smtpAuthAddress)
	return e.Send(smtpServerAddress, auth)
}

func (s *GmailSender) BuildTemplate(name string) ([]byte, error) {
	t := template.New("template.html")

	t, err := t.ParseFiles("mail/template.html")
	if err != nil {
		log.Panic(err)
		return nil, err
	}

	verifyCode := [6]string{}
	for i, _ := range verifyCode {
		verifyCode[i] = fmt.Sprint(rand.Intn(9))
	}

	data := struct {
		VerifyUrl string
		Name      string
	}{
		VerifyUrl: fmt.Sprintf("http://localhost:3000/verify/%s", strings.Join(verifyCode[:], "")),
		Name:      name,
	}

	var tpl bytes.Buffer

	if err := t.Execute(&tpl, data); err != nil {
		return nil, err
	}

	return tpl.Bytes(), nil
}

func ServeEmailTemplate(name string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := template.New("template.html")

		t, err := t.ParseFiles("mail/template.html")
		if err != nil {
			log.Panic(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		verifyCode := [6]string{}
		for i, _ := range verifyCode {
			verifyCode[i] = fmt.Sprint(rand.Intn(9))
		}

		data := struct {
			VerifyUrl string
			Name      string
		}{
			VerifyUrl: fmt.Sprintf("http://localhost:3000/verify/%s", strings.Join(verifyCode[:], "")),
			Name:      name,
		}

		if err := t.Execute(w, data); err != nil {
			log.Panic(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}

func Serve(w http.ResponseWriter, r *http.Request) {
	t := template.New("template.html")

	t, err := t.ParseFiles("mail/template.html")
	if err != nil {
		log.Panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// verifyCode := rand.Intn(9)
	verifyCode := [6]string{}

	for i, _ := range verifyCode {
		verifyCode[i] = fmt.Sprint(rand.Intn(9))
	}

	log.Println(verifyCode)
	log.Println(strings.Join(verifyCode[:], ""))

	data := struct {
		VerifyUrl string
	}{
		VerifyUrl: fmt.Sprintf("http://localhost:3000/verify/%s", strings.Join(verifyCode[:], "")),
	}

	if err := t.Execute(w, data); err != nil {
		log.Panic(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
