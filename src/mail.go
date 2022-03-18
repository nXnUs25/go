package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

var (
	pf  = fmt.Printf
	pln = fmt.Println
)

func SendEmail(smtpPort int, smtpServer, subject, msg, to string) {
	pf("INFO SendEmail to %v\n", to)
	host, _ := os.Hostname()

	from := "golangrun@" + host

	pf("INFO From: %v\n", from)

	smtpHost := fmt.Sprintf("%v:%v", smtpServer, smtpPort)
	pf("INFO Opening connection to SMTP Server %v\n", smtpHost)
	conn, err := smtp.Dial(smtpHost)
	pln("Got 24: ")
	checkError(err)

	defer conn.Close()
	defer pf("INFO Closing connection to SMTP Server %v\n", smtpHost)

	err = conn.Mail(from)
	pln("Got 31: ")
	checkError(err)
	err = conn.Rcpt(to)
	pln("Got 34: ")
	checkError(err)

	wc, err := conn.Data()
	checkError(err)
	pln("Got 39: ")
	defer wc.Close()

	pln(from)
	pln(to)
	pln(subject)
	header_from := "From: " + from
	header_to := "To: " + to
	header_sub := "Subject: " + subject

	body := header_from + "\n" + header_to + "\n" + header_sub + "\n" + msg

	buffer := bytes.NewBufferString(body)
	if _, err := buffer.WriteTo(wc); err != nil {
		pln("Got 50: ")
		pln(err)
	}
	pln("INFO SendEmail done")
}

func checkError(err error) {
	if err != nil {
		pln(err)
	}
}

func SendMail(addr, from, subject, body string, to []string) error {
	r := strings.NewReplacer("\r\n", "", "\r", "", "\n", "", "%0a", "", "%0d", "")

	c, err := smtp.Dial(addr)
	if err != nil {
		return err
	}
	defer c.Close()
	if err = c.Mail(r.Replace(from)); err != nil {
		return err
	}
	for i := range to {
		to[i] = r.Replace(to[i])
		if err = c.Rcpt(to[i]); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	msg := "To: " + strings.Join(to, ",") + "\r\n" +
		"From: " + from + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n" +
		"Content-Transfer-Encoding: base64\r\n" +
		"\r\n" + base64.StdEncoding.EncodeToString([]byte(body))

	_, err = w.Write([]byte(msg))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	return c.Quit()
}

type Address struct {
	val1 string
	val2 string
}

func main() {
	//ex: SendMail("127.0.0.1:25", (&mail.Address{"from name", "from@example.com"}).String(), "Email Subject", "message body", []string{(&mail.Address{"to name", "to@example.com"}).String()})
	//to := []string{"chmielua@gmail.com", "augustyn.chmiel@acoustic.com"}
	//SendMail("mailproxy.atlis1:25", "proxy@bg53-ca-ue1.prod.awspr", "Subject 1", "body text 1", to)
	SendEmail(25, "mailproxy.atlis1", "testing sink", "test golang 1", "chmielua@gmail.com")
}
