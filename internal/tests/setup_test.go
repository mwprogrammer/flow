package tests

import (
	"testing"

	"github.com/mwprogrammer/flow"
)

func TestReplyingWithText(t *testing.T) {

	account_id := "XXXXXX" // Your Whatsapp Business Account Id
	version := "24.0"      // The API version
	access_token := "XXXXXXXXX"
	phone_number_Id := "XXXXXX" // Your Phone number Id

	// Begin by defining your flow settings with your Whatsapp Business Account App details.
	// You can create one using the instructions [here](https://business.facebook.com/business/loginpage/?next=https%3A%2F%2Fdevelopers.facebook.com%2Fapps#)
	new_flow_settings := flow.FlowSettings{
		Id:      account_id,
		Version: version,
		Token:   access_token,
		Sender:  phone_number_Id,
	}

	// Create a new flow using the .New() method
	new_flow := flow.New(new_flow_settings)
	if new_flow == nil {
		t.Errorf(`Could not create new flow`)
	}

	// Send hello world to a specified number.
	err := new_flow.ReplyWithText("XXXXX", "Hello, world!", false)

	if err != nil {
		t.Errorf(`Error replying with text: %v`, err)
	}
}
