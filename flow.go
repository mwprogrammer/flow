package flow

import (
	"github.com/mwprogrammer/flow/internal/client"
	models "github.com/mwprogrammer/flow/internal/payloads"
)

type FlowSettings struct {
	Id      string
	Version string
	Token   string
	Sender  string
}

func New(settings FlowSettings) *Flow {

	return &Flow{
		settings: settings,
	}

}

type Flow struct {
	settings FlowSettings
}

func (f *Flow) ReplyWithText(phone string, message string, previewUrl bool) error {

	payload := models.ReplyWithTextPayload{}

	payload.Product = "whatsapp"
	payload.ReceipientType = "individual"
	payload.To = phone
	payload.Type = "text"
	payload.TextMessage = &models.TextMessage{}
	payload.TextMessage.PreviewUrl = previewUrl
	payload.TextMessage.Body = message

	err := client.PostMessage(
		f.settings.Version,
		f.settings.Token,
		f.settings.Sender, payload,
		"messages")

	if err != nil {
		return err
	}

	return nil

}
