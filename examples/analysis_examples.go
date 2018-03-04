package examples

import (
	"github.com/mailosaur/mailosaur-go/mailosaur"
)

func analysisSpam() {
	client := mailosaur.NewMailosaurClient("YOUR_API_KEY")

	client.Analysis().Spam("MESSAGE_ID")
}
