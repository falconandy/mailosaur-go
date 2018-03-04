package examples

import (
	"github.com/mailosaur/mailosaur-go/mailosaur"
)

func filesGetAttachment() {
	client := mailosaur.NewMailosaurClient("YOUR_API_KEY")

	client.Files().GetAttachment("ID")
}

func filesGetEmail() {
	client := mailosaur.NewMailosaurClient("YOUR_API_KEY")

	client.Files().GetEmail("MESSAGE_ID")
}
