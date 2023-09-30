package mailsender

import (
	"fmt"
	"net/smtp"
	"strconv"

	"github.com/JeromeTGH/TerraScan-collector/config"
)


func Sendmail(sujet string, body string) {

	// Construction du message à envoyer
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", config.AppConfig.Email.From)
	msg += fmt.Sprintf("To: %s\r\n", config.AppConfig.Email.To)
	msg += fmt.Sprintf("Subject: %s\r\n", sujet)
	msg += fmt.Sprintf("\r\n%s\r\n", body)


	// Envoi du message
	auth := smtp.PlainAuth("", config.AppConfig.Email.From, config.AppConfig.Email.Pwd, config.AppConfig.Email.HostName)
	err := smtp.SendMail(
		config.AppConfig.Email.HostName + ":" + strconv.Itoa(config.AppConfig.Email.SmtpPort),
		auth,
		config.AppConfig.Email.From,
		[]string{config.AppConfig.Email.To},
		[]byte(msg))


	// Message en cas d'erreur
	if err != nil {
		fmt.Println("Failed to send this email")
		fmt.Println(err)
	}

}
