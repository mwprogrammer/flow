package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/mwprogrammer/flow"
)

func TestParsingMessage(t *testing.T) {

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

	sample_response := `{"object":"whatsapp_business_account","entry":[{"id":"0","changes":[{"field":"messages","value":{"messaging_product":"whatsapp","metadata":{"display_phone_number":"16505551111","phone_number_id":"123456123"},"contacts":[{"profile":{"name":"test user name"},"wa_id":"16315551181"}],"messages":[{"from":"16315551181","id":"ABGGFlA5Fpa","timestamp":"1504902988","type":"text","text":{"body":"this is a text message"}}]}}]}]}`

	message, err := new_flow.ParseMessage(sample_response)

	debug, err := json.Marshal(message)

	fmt.Println(string(debug))

	if err != nil {
		t.Errorf(`parsing message: %v`, err)
	}

	if err != nil {
		t.Errorf(`parsing message: %v`, err)
	}

}
