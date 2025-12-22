package flow

import (
	"encoding/json"
	"errors"
	"fmt"

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

func (f *Flow) ParseMessage(response string) (*Message, error) {

	var message Message

	data, err := client.ReadMessage(response)

	debug, _ := json.Marshal(data)

	fmt.Println(string(debug))

	if err != nil {
		return nil, err
	}

	product, ok := data["messaging_product"]

	if !ok {
		return nil, errors.New("product property not defined")
	}

	if product != "whatsapp" {
		return nil, errors.New("product is not whatsapp")
	}

	contacts, ok := data["contacts"]

	if !ok {
		return nil, errors.New("contacts property not defined")
	}

	contacts_data, err := json.Marshal(contacts)

	if err != nil {
		return nil, errors.New("could not retrieve contacts property data")
	}

	var receipients []payloads.Receipient

	err = json.Unmarshal(contacts_data, &receipients)

	if err != nil {
		return nil, errors.New("could not retrieve contacts property data")
	}

	message.Name = receipients[0].Profile.Name

	messages, ok := data["messages"]

	if !ok {
		return nil, errors.New("messages property not defined")
	}

	messages_data, err := json.Marshal(messages)

	if err != nil {
		return nil, errors.New("could not retrieve messages property data")
	}

	var receipient_messages []payloads.ReceipientMessage

	err = json.Unmarshal(messages_data, &receipient_messages)

	if err != nil {
		return nil, errors.New("could not retrieve messages property data")
	}

	if len(receipient_messages) == 0 {
		return nil, errors.New("messages array is empty")
	}

	message.Id = receipient_messages[0].Id
	message.PhoneNumber = receipient_messages[0].From
	message.Type = receipient_messages[0].Type

	if receipient_messages[0].Type == "text" {

		var text_message payloads.TextReceipientMessage

		err = json.Unmarshal(messages_data, &text_message)

		if err != nil {
			return nil, errors.New("could not retrieve messages property text data")
		}

		message.Content = text_message.Text.Body
	}

	return &message, nil

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
