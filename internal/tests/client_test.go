package tests

import (
	"os"
	"testing"

	"github.com/mwprogrammer/flow/internal/client"
)

var id string
var version string
var token string
var senderPhoneNumber string

func configure() {

	id = os.Getenv("WHATSAPP_BUSINESS_ID")
	version = os.Getenv("WHATSAPP_CLOUD_API_VERSION")
	token = os.Getenv("WHATSAPP_ACCESS_TOKEN")
	senderPhoneNumber = os.Getenv("WHATSAPP_SENDER_NUMBER")
}

func TestReadingMessage(t *testing.T) {

	samplePayload := `{"object":"whatsapp_business_account","entry":[{"id":"0","changes":[{"field":"messages","value":{"messaging_product":"whatsapp","metadata":{"display_phone_number":"16505551111","phone_number_id":"123456123"},"contacts":[{"profile":{"name":"test user name"},"wa_id":"16315551181"}],"messages":[{"from":"16315551181","id":"ABGGFlA5Fpa","timestamp":"1504902988","type":"text","text":{"body":"this is a text message"}}]}}]}]}`

	_, err := client.ReadMessage(samplePayload)

	if err != nil {
		t.Errorf("Error reading message: %v", err)
	}

}

func TestPostingMessage(t *testing.T) {
	
	configure()

	err := client.PostMessage(version, token, senderPhoneNumber, "", "messages")

	if err != nil {
		t.Errorf("Error posting message: %v", err)
	}

}