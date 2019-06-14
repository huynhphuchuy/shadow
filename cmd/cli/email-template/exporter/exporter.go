package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"server/cmd/cli/email-template/templates"
	"server/internal/config"

	"github.com/matcornic/hermes"
)

type Template struct {
	Email hermes.Email
	Type  string
}

func main() {
	environment := flag.String("e", "dev", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	config.Init(*environment)

	config := config.GetConfig()
	// Configure hermes by setting a theme and your product info
	h := hermes.Hermes{
		// Optional Theme
		Theme: new(hermes.Flat),
		Product: hermes.Product{
			// Appears in header & footer of e-mails
			Name: config.GetString("product.name"),
			Link: config.GetString("product.url"),
			// Optional product logo
			Logo: config.GetString("product.logo"),
		},
	}

	allTemplates := [1]Template{
		Template{
			templates.Confirmation{
				"Anonymous",
				[]string{
					"Welcome to Golang! We're very excited to have you on board.",
				},
				"To get started with Golang, please click here:",
				templates.Button{
					"#22BC66", // Optional action button color
					"Confirm your account",
					"#",
				},
				[]string{
					"Need help, or have questions? Just reply to this email, we'd love to help.",
				},
			}.Init(),
			"confirmation",
		},
	}

	for _, t := range allTemplates {
		// Generate an HTML email with the provided contents (for modern clients)
		emailBody, err := h.GenerateHTML(t.Email)
		if err != nil {
			panic(err) // Tip: Handle error with something else than a panic ;)
		}

		// Generate the plaintext version of the e-mail (for clients that do not support xHTML)
		emailText, err := h.GeneratePlainText(t.Email)
		if err != nil {
			panic(err) // Tip: Handle error with something else than a panic ;)
		}

		// Optionally, preview the generated HTML e-mail by writing it to a local file
		err = ioutil.WriteFile("./cmd/cli/email-template/exported/html/"+t.Type+".html", []byte(emailBody), 0644)
		if err != nil {
			panic(err) // Tip: Handle error with something else than a panic ;)
		}

		// Optionally, preview the generated HTML e-mail by writing it to a local file
		err = ioutil.WriteFile("./cmd/cli/email-template/exported/text/"+t.Type+".txt", []byte(emailText), 0644)
		if err != nil {
			panic(err) // Tip: Handle error with something else than a panic ;)
		}
	}

}
