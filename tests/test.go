package tests

import (
	"testing"

	"github.com/mwprogrammer/flow"
)

func TestReplyingWithText(t *testing.T) {

	business_account_id := "340311965831782"
	version := "24.0"
	access_token := "EAASykfivjusBP7N4KvRI51qmarzTcTXlNjzTiqZBHOLDDI2jOmayYUl6udArCPKImb2OIYhjFd1Vd8YbvZBWOpZBmD4uZCX9pvLZBTFeFcKoL43fF64ZBSZCuHKsxL1R0E9snypOytbiLRAVr9zvgLhEs3FZCyqIPk0ZA12hCTpDWI58ppNqYZCRwbyHdVm9dzQfhzSum01n2iLL8L4RsPScZCjifAeqlRQbViyke1LhgMRgXQTE5F3nNxoFuF1l4NcJGZBQf5ozQYII5aCALJyB51ZBsnBnXbAZDZD"

	flow.Setup(business_account_id, version, access_token)

	err := flow.ReplyWithText("265881931635", "Bho bho senior!")

	if err != nil {
		t.Errorf(`Test failed: %v`, err)
	}

}
