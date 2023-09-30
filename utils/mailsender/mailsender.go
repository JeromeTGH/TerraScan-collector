package mailsender

import (
	"fmt"

	"github.com/JeromeTGH/TerraScan-collector/config"
	"github.com/JeromeTGH/TerraScan-collector/utils/logger"
	"gopkg.in/gomail.v2"
)


func Sendmail(sujet string, body string) {

    message := gomail.NewMessage()
    message.SetHeader("From", config.AppConfig.Email.From)
    message.SetHeader("To", config.AppConfig.Email.To)
    message.SetHeader("Subject", sujet)
    message.SetBody("text/html", body)

	dialer := gomail.NewDialer(config.AppConfig.Email.HostName, config.AppConfig.Email.SmtpPort, config.AppConfig.Email.From, config.AppConfig.Email.Pwd)

    if err := dialer.DialAndSend(message); err != nil {
		logger.WriteLogWithoutPrinting("mailsender", "failed to send mail")
        panic(err)
    }
	
	fmt.Println("Email sent successfully")

}
