/*
Package flow provides a library for building interactive applications on the WhatsApp Business platform.

It offers a high-level wrapper over the WhatsApp Cloud API, abstracting tasks such as
marking messages as read, displaying typing indicators, sending text messages and more.

Author: Chisomo Chiweza (mwprogrammer)
*/
package flow

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mwprogrammer/flow/internal/client"
)

// Flow provides access to methods which allow you to
// build interactive apps over the WhatsApp Business platform.
type Flow struct {
	settings Settings
}

// Settings Represents configuration like your WhatsApp Business ID, Cloud API Version,
// Access Token and Sender Phone Number to allow you to send messages.
type Settings struct {
	ID      string
	Version string
	Token   string
	Sender  string
}

// Message represents the key information of a receipient message sent to your app over WhatsApp.
type Message struct {
	ID          string
	PhoneNumber string
	Name        string
	Type        string
	Content     string
}

// New constructs a new Flow object.
func New(settings Settings) *Flow {

	//TODO: Validate the configured settings somewhere here.
	return &Flow{
		settings: settings,
	}

}

// ParseMessage reads the JSON payload Meta sends as an event once someone sends a message to your
// app on WhatsApp and parses it to return a FlowMessage.
func (f *Flow) ParseMessage(payload string) (*Message, error) {

	var message Message

	data, err := client.ReadMessage(payload)

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

	contactsData, err := json.Marshal(contacts)

	if err != nil {
		return nil, errors.New("could not retrieve contacts property data")
	}

	var receipients []client.Recipient

	err = json.Unmarshal(contactsData, &receipients)

	if err != nil {
		return nil, errors.New("could not retrieve contacts property data")
	}

	message.Name = receipients[0].Profile.Name

	messages, ok := data["messages"]

	if !ok {
		return nil, errors.New("messages property not defined")
	}

	messagesData, err := json.Marshal(messages)

	if err != nil {
		return nil, errors.New("could not retrieve messages property data")
	}

	var receipientsMessages []client.RecipientMessage

	err = json.Unmarshal(messagesData, &receipientsMessages)

	if err != nil {
		return nil, errors.New("could not retrieve messages property data")
	}

	if len(receipientsMessages) == 0 {
		return nil, errors.New("messages array is empty")
	}

	message.ID = receipientsMessages[0].ID
	message.PhoneNumber = receipientsMessages[0].From
	message.Type = receipientsMessages[0].Type

	if receipientsMessages[0].Type == "text" {

		var textMessages []client.TextRecipientMessage

		err = json.Unmarshal(messagesData, &textMessages)

		fmt.Println(string(messagesData))

		if err != nil {
			fmt.Println(err.Error())
			return nil, errors.New("could not retrieve messages property text data")
		}

		message.Content = textMessages[0].Text.Body

	}

	return &message, nil

}

// MarkAsRead marks a message a receipient has sent to your app as read (ie the blueticks).
func (f *Flow) MarkAsRead(phone string, messageID string) error {

	payload := client.MarkAsReadPayload{}

	payload.Product = "whatsapp"
	payload.RecipientType = "individual"
	payload.To = phone
	payload.Type = "text"
	payload.MessageID = messageID
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

// DisplayTypingIndicator renders the "..." icon to the receipient to simulate your app processing a response.
func (f *Flow) DisplayTypingIndicator(phone string, messageID string) error {

	payload := client.DisplayTypingPayload{}

	payload.Product = "whatsapp"
	payload.RecipientType = "individual"
	payload.To = phone
	payload.Type = "text"
	payload.MessageID = messageID
	payload.Status = "read"
	payload.Indicator = client.TypingIndicator{}
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

// ReplyWithText Allows your app to reply with simple text to a receipient.
// If text is a link, setting previewUrl to true allows the final message sent
// to display a preview of the link.
func (f *Flow) ReplyWithText(phone string, message string, previewURL bool) error {

	payload := client.ReplyWithTextPayload{}

	payload.Product = "whatsapp"
	payload.RecipientType = "individual"
	payload.To = phone
	payload.Type = "text"
	payload.TextMessage = &client.TextMessage{}
	payload.TextMessage.PreviewURL = previewURL
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
