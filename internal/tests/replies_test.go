package tests

import (
	"testing"

	"github.com/mwprogrammer/flow"
)

func TestReplyingWithText(t *testing.T) {

	account_id := "XXXXXX"
	version := "24.0"
	access_token := "XXXXXXXXX"
	phone_number_Id := "XXXXXX"

	flow_settings := flow.FlowSettings{
		Id:      account_id,
		Version: version,
		Token:   access_token,
		Sender:  phone_number_Id,
	}

	new_flow := flow.New(flow_settings)
	if new_flow == nil {
		t.Errorf(`Could not create new flow`)
	}

	err := new_flow.ReplyWithText("XXXXX", "Hello, world!", false)

	if err != nil {
		t.Errorf(`Error replying with text: %v`, err)
	}
}
