package smtp

import (
	"errors"
	"fmt"
	"net/mail"
	"server/internal/config"
	"strconv"

	"server/internal/helpers/generator"

	"github.com/matcornic/hermes"
	"golang.org/x/crypto/ssh/terminal"
	"gopkg.in/gomail.v2"
)

type smtpAuthentication struct {
	Server         string
	Port           int
	SenderEmail    string
	SenderIdentity string
	SMTPUser       string
	SMTPPassword   string
}

// sendOptions are options for sending an email
type sendOptions struct {
	To      string
	Subject string
}

// send sends the email
func send(smtpConfig smtpAuthentication, options sendOptions, htmlBody string, txtBody string) error {

	if smtpConfig.Server == "" {
		return errors.New("SMTP server config is empty")
	}
	if smtpConfig.Port == 0 {
		return errors.New("SMTP port config is empty")
	}

	if smtpConfig.SMTPUser == "" {
		return errors.New("SMTP user is empty")
	}

	if smtpConfig.SenderIdentity == "" {
		return errors.New("SMTP sender identity is empty")
	}

	if smtpConfig.SenderEmail == "" {
		return errors.New("SMTP sender email is empty")
	}

	if options.To == "" {
		return errors.New("no receiver emails configured")
	}

	from := mail.Address{
		Name:    smtpConfig.SenderIdentity,
		Address: smtpConfig.SenderEmail,
	}

	m := gomail.NewMessage()
	m.SetHeader("From", from.String())
	m.SetHeader("To", options.To)
	m.SetHeader("Subject", options.Subject)

	m.SetBody("text/plain", txtBody)
	m.AddAlternative("text/html", htmlBody)

	d := gomail.NewDialer(smtpConfig.Server, smtpConfig.Port, smtpConfig.SMTPUser, smtpConfig.SMTPPassword)

	return d.DialAndSend(m)
}

func SendEmail(email, subject string, template hermes.Email) {

	config := config.GetConfig()

	port, _ := strconv.Atoi(config.GetString("smtp.port"))
	password := config.GetString("smtp.password")
	SMTPUser := config.GetString("smtp.username")
	if password == "" {
		fmt.Printf("Enter SMTP password of '%s' account: ", SMTPUser)
		bytePassword, _ := terminal.ReadPassword(0)
		password = string(bytePassword)
	}
	smtpConfig := smtpAuthentication{
		Server:         config.GetString("smtp.host"),
		Port:           port,
		SenderEmail:    SMTPUser,
		SenderIdentity: config.GetString("product.name"),
		SMTPPassword:   password,
		SMTPUser:       SMTPUser,
	}
	options := sendOptions{
		To: email,
	}

	options.Subject = subject
	htmlBytes, txtBytes := generator.Export(template)

	err := send(smtpConfig, options, string(htmlBytes), string(txtBytes))
	if err != nil {
		panic(err)
	}

}
