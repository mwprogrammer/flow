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

// Settings represents configuration like your WhatsApp Business ID, Cloud API Version,
// Access Token and Sender Phone Number to allow you to send messages.
type Settings struct {
	ID      string
	Version string
	Token   string
	Sender  string
}

// Event represents the key information contained within a JSON payload sent by Meta as an event.
type Event struct {
	Type    string
	Payload string
	Message *Message
}

// Message contains the details of a message sent by a receipient to your app.
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

// ParseEvent reads the JSON payload Meta sends to your app based on events that
// occur as users interact with your app and parses it to return an Event.
func (f *Flow) ParseEvent(payload string) (*Event, error) {

	var event Event
	var message Message

	data, err := client.ReadMessage(payload)

	debug, _ := json.Marshal(data)

	fmt.Println(string(debug))

	if err != nil {
		return nil, err
	}

	product, ok := data["messaging_product"]

	if !ok {
		return nil, errors.New("incorrect format: product property not defined in payload")
	}

	if product != "whatsapp" {
		return nil, errors.New("incorrect format: product is not whatsapp")
	}

	event = Event{}
	event.Payload = string(debug)

	contacts, ok := data["contacts"]

	if !ok {
		//TODO: Add other data if not message event

		event.Type = "Notification Event"
		event.Message = nil

		return nil, nil
	}

	contactsData, err := json.Marshal(contacts)

	if err != nil {
		return nil, errors.New("incorrect format : could not retrieve contacts property data")
	}

	var receipients []client.Recipient

	err = json.Unmarshal(contactsData, &receipients)

	if err != nil {
		return nil, errors.New("incorrect format : could not retrieve contacts property data")
	}

	message.Name = receipients[0].Profile.Name

	messages, ok := data["messages"]

	if !ok {
		return nil, errors.New("incorrect format : messages property not defined")
	}

	messagesData, err := json.Marshal(messages)

	if err != nil {
		return nil, errors.New("incorrect format : could not retrieve messages property data")
	}

	var receipientsMessages []client.RecipientMessage

	err = json.Unmarshal(messagesData, &receipientsMessages)

	if err != nil {
		return nil, errors.New("incorrect format : could not retrieve messages property data")
	}

	if len(receipientsMessages) == 0 {
		return nil, errors.New("incorrect format : messages array is empty")
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
			return nil, errors.New("incorrect format : could not retrieve messages property text data")
		}

		message.Content = textMessages[0].Text.Body

	}

	event.Type = "Message Event"
	event.Message = &message

	return &event, nil

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
