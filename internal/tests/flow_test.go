// Package tests includes tests for the flow library.
package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/mwprogrammer/flow"
)

var settings flow.Settings
var app *flow.Flow

func setup() {

	settings = flow.Settings{

		ID:      os.Getenv("WHATSAPP_BUSINESS_ID"),
		Version: os.Getenv("WHATSAPP_CLOUD_API_VERSION"),
		Token:   os.Getenv("WHATSAPP_ACCESS_TOKEN"),
		Sender:  os.Getenv("WHATSAPP_SENDER_NUMBER"),
	}

	app = flow.New(settings)

}

// TestParsingMessage tests that the JSON payloads sent by Meta as events once users send messages are
// being parsed properly as Flow Messages.
func TestParsingMessage(t *testing.T) {

	setup()

	testPayload := `{"object":"whatsapp_business_account","entry":[{"id":"0","changes":[{"field":"messages","value":{"messaging_product":"whatsapp","metadata":{"display_phone_number":"16505551111","phone_number_id":"123456123"},"contacts":[{"profile":{"name":"test user name"},"wa_id":"16315551181"}],"messages":[{"from":"16315551181","id":"ABGGFlA5Fpa","timestamp":"1504902988","type":"text","text":{"body":"this is a text message"}}]}}]}]}`
	message, err := app.ParseEvent(testPayload)

	if err != nil {
		t.Errorf(`parsing message: %v`, err)
	}

	debug, err := json.Marshal(message)

	fmt.Println(string(debug))

	if err != nil {
		t.Errorf(`parsing message: %v`, err)
	}

}

// TestReplyingWithText that the flow app is able to send text messages to receipients
func TestReplyingWithText(t *testing.T) {

	setup()

	err := app.ReplyWithText("XXXXX", "Hello, world!", false)

	if err != nil {
		t.Errorf(`Error replying with text: %v`, err)
	}
}
