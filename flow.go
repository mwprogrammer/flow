package flow

import (
	"github.com/mwprogrammer/flow/client"
	"github.com/mwprogrammer/flow/models"
)

var settings models.Settings

func Setup(

	business_account_id string,
	api_version string,
	access_token string) {

	settings = models.Settings{}

	settings.Id = business_account_id
	settings.Version = api_version
	settings.AccessToken = access_token

}

func BlueTickMessage(messageId int) error {

	payload := models.Payload{}

	payload.Product = "whatsapp"
	payload.Status = "read"
	payload.MessageId = messageId

	err := client.PostMessage(settings, payload, "messages")

	if err != nil {
		return err
	}

	return nil

}

func ReplyWithText(phone string, message string) error {

	payload := models.Payload{}

	payload.Product = "whatsapp"
	payload.PreviewUrl = false
	payload.To = phone
	payload.Type = "text"
	payload.TextMessage.Body = message

	err := client.PostMessage(settings, payload, "messages")

	if err != nil {
		return err
	}

	return nil

}

func ReplyWithButtons(
	phone string,
	message string,
	buttons []models.ButtonMessage) error {

	payload := models.Payload{}

	payload.Product = "whatsapp"
	payload.ReceipientType = "individual"
	payload.To = phone
	payload.Type = "interactive"

	interactive_message := models.InteractiveMessage{}

	interactive_message.Type = "button"
	interactive_message.Body.Text = message
	interactive_message.Action.Buttons = buttons

	payload.InteractiveMessage = &interactive_message

	err := client.PostMessage(settings, payload, "messages")

	if err != nil {
		return err
	}

	return nil

}

func ReplyWithList(
	phone string,
	header string,
	body string,
	footer string,
	action string,
	sections []models.SectionMessage) error {

	payload := models.Payload{}

	payload.Product = "whatsapp"
	payload.ReceipientType = "individual"
	payload.To = phone
	payload.Type = "interactive"

	interactive_message := models.InteractiveMessage{}

	interactive_message.Type = "list"
	interactive_message.Header.Text = header
	interactive_message.Body.Text = body
	interactive_message.Footer.Text = footer
	interactive_message.Action.Button = action
	interactive_message.Action.Sections = sections

	payload.InteractiveMessage = &interactive_message

	err := client.PostMessage(settings, payload, "messages")

	if err != nil {
		return err
	}

	return nil

}
