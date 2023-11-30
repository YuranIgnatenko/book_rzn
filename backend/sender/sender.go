package sender

import (
	"fmt"
	"log"
	"net/smtp"
)

// used mail Yandex.ru account
func Send_mail(Subj, text string) {
	user := "yuran.ignatenko@yandex.ru"
	password := "Quizzaciously1"

	to := []string{
		"asokolova365@gmail.com",
	}

	from := "yuran.ignatenko@yandex.ru"

	addr := "smtp.yandex.ru:25"
	host := "smtp.yandex.ru"

	msg := []byte("From: yuran.ignatenko@yandex.ru\r\n" +
		"asokolova365@gmail.com\r\n" +
		"Subject: " + Subj + "\r\n\r\n" +
		text + "\r\n")

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, to, msg)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Email sent successfully")

}
