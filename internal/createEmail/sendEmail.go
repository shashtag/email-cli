package createemail

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/spf13/viper"
)

func SendEmail(table string) error {
	headers := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
	msg := "Subject:" + "wsgm" + "\n" + headers + "\n\n" + table
	fmt.Println(viper.GetString("CLIENT_EMAIL"), "ssmsmms")
	auth := smtp.PlainAuth("", viper.GetString("CLIENT_EMAIL"), viper.GetString("CLIENT_PASSWORD"), viper.GetString("SMTP_SERVER"))
	err := smtp.SendMail(viper.GetString("SMTP_SERVER")+":587", auth, "shashwatsatna@gmail.com", []string{"shashwatsatna@gmail.com"}, []byte(msg))
	if err != nil {
		log.Fatal(err)
		return err
	}
	// Send the email
	fmt.Println(os.Getenv("SMTP_SERVER"))
	return nil
}
