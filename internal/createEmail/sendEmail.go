package createemail

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/spf13/viper"
)

func SendEmail() error {
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	msg := "Subject:" + "wsgm" + "\n" + headers + "\n\n" + "<div>Hi, This is a test email</div>"
	auth := smtp.PlainAuth("", viper.GetString("Email"), viper.GetString("Password"), viper.GetString("SMTP_SERVER"))
	err := smtp.SendMail(viper.GetString("SMTP_SERVER")+":587", auth, "shashwatsatna@gmail.com", []string{"shashwatsatna@gmail.com"}, []byte(msg))
	if err != nil {
		log.Fatal(err)
		return err
	}
	// Send the email
	fmt.Println(os.Getenv("SMTP_SERVER"))
	return nil
}
