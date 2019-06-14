package generator

import (
	"server/internal/config"

	"github.com/matcornic/hermes"
)

func Export(e hermes.Email) (string, string) {

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

	// Generate an HTML email with the provided contents (for modern clients)
	emailBody, err := h.GenerateHTML(e)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}

	// Generate the plaintext version of the e-mail (for clients that do not support xHTML)
	emailText, err := h.GeneratePlainText(e)
	if err != nil {
		panic(err) // Tip: Handle error with something else than a panic ;)
	}
	return emailBody, emailText
}
