package createemail

import (
	"fmt"
	"log"
	"net/smtp"
	"os"

	"github.com/spf13/viper"
)

func SendEmail() error {
	auth := smtp.PlainAuth("", viper.GetString("Email"), viper.GetString("Password"), viper.GetString("SMTP_SERVER"))
	err := smtp.SendMail(viper.GetString("SMTP_SERVER")+":587", auth, "shashwatsatna@gmail.com", []string{"shashwatsatna@gmail.com"}, []byte("Hello, World!"))
	if err != nil {
		log.Fatal(err)
		return err
	}
	// Send the email
	fmt.Println(os.Getenv("SMTP_SERVER"))
	return nil
}
