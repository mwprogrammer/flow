package flow

import (
	"github.com/mwprogrammer/flow/internal/client"
	models "github.com/mwprogrammer/flow/internal/types"
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

func (f *Flow) BlueTickMessage(messageId int) error {

	payload := models.Payload{}

	payload.Product = "whatsapp"
	payload.Status = "read"
	payload.MessageId = messageId

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

func (f *Flow) ReplyWithText(phone string, message string) error {

	payload := models.Payload{}

	payload.Product = "whatsapp"
	payload.PreviewUrl = false
	payload.To = phone
	payload.Type = "text"
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
