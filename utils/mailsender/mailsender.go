package mailsender

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"gopkg.in/gomail.v2"
)


func Sendmail(sujet string, body string) {

    msg := gomail.NewMessage()
    msg.SetHeader("From", config.AppConfig.Email.From)
    msg.SetHeader("To", config.AppConfig.Email.To)
    msg.SetHeader("Subject", sujet)
    msg.SetBody("text/html", body)

	n := gomail.NewDialer(config.AppConfig.Email.HostName, config.AppConfig.Email.SmtpPort, config.AppConfig.Email.From, config.AppConfig.Email.Pwd)

    if err := n.DialAndSend(msg); err != nil {
        panic(err)
    }
	
	fmt.Println("Email sent successfully")

}
