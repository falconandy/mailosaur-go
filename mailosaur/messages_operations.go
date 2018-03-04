package mailosaur

import "fmt"

type MessagesOperations struct {
	client *MailosaurClient
}

func newMessagesOperations(client *MailosaurClient) *MessagesOperations {
	return &MessagesOperations{
		client: client,
	}
}

// Retrieves the detail for a single email message.
// Simply supply the unique identifier for the required message.
func (op *MessagesOperations) Get(id string) (*Message, error) {
	var message *Message
	err := op.client.get("messages/"+id, &message)
	return message, err
}

// Permanently deletes a message. This operation cannot be undone.
// Also deletes any attachments related to the message.
func (op *MessagesOperations) Delete(id string) error {
	return op.client.delete("messages/" + id)
}

// Returns a list of your messages in summary form.
// The summaries are returned sorted by received date, with the most recently-received messages appearing first.
func (op *MessagesOperations) List(server string) (MessageListResult, error) {
	var messageList MessageListResult
	err := op.client.get("messages?server="+server, &messageList)
	return messageList, err
}

// Returns the page with a list of your messages in summary form.
// The summaries are returned sorted by received date, with the most recently-received messages appearing first.
func (op *MessagesOperations) ListPage(server string, page int, itemsPerPage int) (MessageListResult, error) {
	apiPath := op.getPagingApiPath("messages?server="+server, page, itemsPerPage)
	var messageList MessageListResult
	err := op.client.get(apiPath, &messageList)
	return messageList, err
}

// Permanently deletes all messages held by the specified server. This operation cannot be undone.
// Also deletes any attachments related to each message.
func (op *MessagesOperations) DeleteAll(server string) error {
	return op.client.delete("messages?server=" + server)
}

// Returns a list of messages matching the specified search criteria, in summary form.
// The messages are returned sorted by received date, with the most recently-received messages appearing first.
func (op *MessagesOperations) Search(server string, criteria SearchCriteria) (MessageListResult, error) {
	var messageList MessageListResult
	err := op.client.post("messages/search?server="+server, criteria, &messageList)
	return messageList, err
}

// Returns the page with a list of messages matching the specified search criteria, in summary form.
// The messages are returned sorted by received date, with the most recently-received messages appearing first.
func (op *MessagesOperations) SearchPage(server string, criteria SearchCriteria, page int, itemsPerPage int) (MessageListResult, error) {
	apiPath := op.getPagingApiPath("messages/search?server="+server, page, itemsPerPage)
	var messageList MessageListResult
	err := op.client.post(apiPath, criteria, &messageList)
	return messageList, err
}

// Returns as soon as a message matching the specified search criteria is found.
// This is the most efficient method of looking up a message.
func (op *MessagesOperations) WaitFor(server string, criteria SearchCriteria) (*Message, error) {
	var message *Message
	err := op.client.post("messages/await?server="+server, criteria, &message)
	return message, err
}

func (op *MessagesOperations) getPagingApiPath(baseApiPath string, page int, itemsPerPage int) string {
	result := baseApiPath
	if page != 0 {
		result += fmt.Sprintf("&page=%d", page)
	}
	if itemsPerPage != 0 {
		result += fmt.Sprintf("&itemsPerPage=%d", itemsPerPage)
	}
	return result
}
