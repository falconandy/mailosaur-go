package examples

import (
	"github.com/mailosaur/mailosaur-go/mailosaur"
)

func messagesGet() {
	client := mailosaur.NewMailosaurClient("YOUR_API_KEY")

	client.Messages().Get("MESSAGE_ID")
}

func messagesDelete() {
	client := mailosaur.NewMailosaurClient("YOUR_API_KEY")

	client.Messages().Delete("MESSAGE_ID")
}

func messagesList() {
	client := mailosaur.NewMailosaurClient("YOUR_API_KEY")

	client.Messages().List("SERVER_ID")
}

func messagesListPage() {
	client := mailosaur.NewMailosaurClient("YOUR_API_KEY")

	client.Messages().ListPage("SERVER_ID", 3, 100)
}

func messagesDeleteAll() {
	client := mailosaur.NewMailosaurClient("YOUR_API_KEY")

	client.Messages().DeleteAll("SERVER_ID")
}

func messagesSearch() {
	client := mailosaur.NewMailosaurClient("YOUR_API_KEY")

	criteria := mailosaur.SearchCriteria{
		SentTo: "someone.abc@mailosaur.io",
	}
	client.Messages().Search("SERVER_ID", criteria)
}

func messagesSearchPage() {
	client := mailosaur.NewMailosaurClient("YOUR_API_KEY")

	criteria := mailosaur.SearchCriteria{
		SentTo: "someone.abc@mailosaur.io",
	}
	client.Messages().SearchPage("SERVER_ID", criteria, 3, 100)
}

func messagesWaitFor() {
	client := mailosaur.NewMailosaurClient("YOUR_API_KEY")

	criteria := mailosaur.SearchCriteria{
		SentTo: "someone.abc@mailosaur.io",
	}
	client.Messages().WaitFor("SERVER_ID", criteria)
}
