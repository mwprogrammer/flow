package flow

import (
	"github.com/mwprogrammer/flow/internal/client"
	payloads "github.com/mwprogrammer/flow/internal/payloads"
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

func (f *Flow) MarkAsRead(phone string, message_id string) error {

	payload := payloads.MarkAsReadPayload{}

	payload.Product = "whatsapp"
	payload.ReceipientType = "individual"
	payload.To = phone
	payload.Type = "text"
	payload.MessageId = message_id
	payload.Status = "read"

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

func (f *Flow) DisplayTypingIndicator(phone string, message_id string) error {

	payload := payloads.DisplayTypingPayload{}

	payload.Product = "whatsapp"
	payload.ReceipientType = "individual"
	payload.To = phone
	payload.Type = "text"
	payload.MessageId = message_id
	payload.Status = "read"
	payload.Indicator = payloads.TypingIndicator{}
	payload.Indicator.Type = "text"

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

func (f *Flow) ReplyWithText(phone string, message string, previewUrl bool) error {

	payload := payloads.ReplyWithTextPayload{}

	payload.Product = "whatsapp"
	payload.ReceipientType = "individual"
	payload.To = phone
	payload.Type = "text"
	payload.TextMessage = &payloads.TextMessage{}
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
