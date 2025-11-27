package tests

import (
	"testing"

	"github.com/mwprogrammer/flow"
)

func TestReplyingWithText(t *testing.T) {

	account_id := "335597306303950" // Your Whatsapp Business Account Id
	version := "24.0"               // The API version
	access_token := "EAASykfivjusBQFcVrtZA5MuRAIuA0HGZCG21DvAsKMYmZA9dNnVpyZARYxWwySPwrZBkixQ8ZCVfiHGAMUYZBpxugKcx3RVCXMFnYMKVFA5aGBoDh55Rtq80sTi5Xbgl4fxDt1ZCFlljP4zTQ6al1beLSNmO1nhEnZASpOQmyZApQBpGhMwOZATvnzC8jFkGAd6iE3HuDw94hcJLUG9eN23cPASjfGtKo8d2cZAZBhG06ttM7siCXKD9TsjSSZAjSf6mXWHGWZADYe4632559qx8MRvhSuY"
	phone_number_Id := "340311965831782" // Your Phone number Id

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
	err := new_flow.ReplyWithText("265884286800", "This is a test message using the flow library", false)

	if err != nil {
		t.Errorf(`Error replying with text: %v`, err)
	}
}
