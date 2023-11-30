package sender

import (
	"log"
	"net/smtp"
)

// used mail Yandex.ru account
func Send_mail(Subj, text string) {
	// fmt.Println("START EMAIL BLOCK CODE - SENDER")
	user := "yuran.ignatenko@yandex.ru"
	password := "Quizzaciously1"

	to := []string{
		"asokolova365@gmail.com",
	}

	from := "yuran.ignatenko@yandex.ru"

	//yandex
	addr := "smtp.yandex.ru:25"
	host := "smtp.yandex.ru"

	//gmail
	//smtpHost := "smtp.gmail.com"
	//smtpPort := "587"

	msg := []byte("From: yuran.ignatenko@yandex.ru\r\n" +
		"asokolova365@gmail.com\r\n" +
		"Subject: " + Subj + "\r\n\r\n" +
		text + "\r\n")

	auth := smtp.PlainAuth("", user, password, host)

	err := smtp.SendMail(addr, auth, from, to, msg)

	if err != nil {
		log.Fatal(err)
	}

}

// func Send_html() {

// 	from := "yuran.ignatenko@yandex.ru"
// 	password := "Quizzaciously1"
// 	// Receiver email address.
// 	to := []string{
// 		"asokolova365@gmail.com",
// 	}
// 	// error if not '[]'
// 	smtpHost := "[smtp.yandex.ru:25/]"
// 	smtpPort := "smtp.yandex.ru"

// 	auth := smtp.PlainAuth("", from, password, smtpHost)

// 	t, _ := template.ParseFiles("sender/mail.html")

// 	var body bytes.Buffer

// 	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
// 	body.Write([]byte(fmt.Sprintf("Subject: Новый Заказ \n%s\n\n", mimeHeaders)))

// 	t.Execute(&body, struct {
// 		Name    string
// 		Message string
// 	}{
// 		Name:    "Puneet Singh",
// 		Message: "This is a test message in a HTML template",
// 	})

// 	// Sending email.
// 	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, body.Bytes())
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("Email Sent!")
// }
