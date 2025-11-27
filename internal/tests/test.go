package tests

import (
	"testing"

	"github.com/mwprogrammer/flow"
)

func TestReplyingWithText(t *testing.T) {

	business_account_id := "340311965831782"
	version := "24.0"
	access_token := "EAASykfivjusBP7N4KvRI51qmarzTcTXlNjzTiqZBHOLDDI2jOmayYUl6udArCPKImb2OIYhjFd1Vd8YbvZBWOpZBmD4uZCX9pvLZBTFeFcKoL43fF64ZBSZCuHKsxL1R0E9snypOytbiLRAVr9zvgLhEs3FZCyqIPk0ZA12hCTpDWI58ppNqYZCRwbyHdVm9dzQfhzSum01n2iLL8L4RsPScZCjifAeqlRQbViyke1LhgMRgXQTE5F3nNxoFuF1l4NcJGZBQf5ozQYII5aCALJyB51ZBsnBnXbAZDZD"
	sender_phone_number := "01245678788"

	flow_settings := flow.FlowSettings{
		Id:      business_account_id,
		Version: version,
		Token:   access_token,
		Sender:  sender_phone_number,
	}

	new_flow := flow.New(flow_settings)

	new_flow.ReplyWithText("10830576685", "Hello world!")
}
